package tenant

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"system/internal/dao/model"
	"system/internal/logic/dept"
	"system/internal/logic/user"
	"time"
	"toolkit/errx"
	"toolkit/utils"

	"system/internal/logic/role"
	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.ModifyTenantReq) error {
	q := l.svcCtx.Query
	//获取当前最大的租户编号
	var maxTenantId string
	if err := q.SysTenant.WithContext(l.ctx).Select(q.SysTenant.TenantID).Order(q.SysTenant.TenantID.Desc()).Limit(1).Scan(&maxTenantId); err != nil {
		return errx.GORMErr(err)
	}
	tenantId := strconv.Itoa(int(utils.StrAtoi(maxTenantId) + 1))
	expireTime, err := time.Parse(time.DateTime, req.ExpireTime)
	if err != nil {
		return errx.BizErr("时间格式错误")
	}
	tenant := &model.SysTenant{
		ID:              utils.GetID(),
		TenantID:        tenantId,
		ContactUserName: req.ContactUserName,
		ContactPhone:    req.ContactPhone,
		CompanyName:     req.CompanyName,
		LicenseNumber:   req.LicenseNumber,
		Intro:           req.Intro,
		Domain:          req.Domain,
		Remark:          req.Remark,
		PackageID:       req.PackageId,
		ExpireTime:      expireTime,
		AccountCount:    req.AccountCount,
		Address:         req.Address,
	}
	if err := q.SysTenant.WithContext(l.ctx).Create(tenant); err != nil {
		return errx.GORMErr(err)
	}
	//TODO 根据租户套餐新增租户的管理员账号
	err = l.createAdminUser(tenant, req)
	if err != nil {
		return err
	}
	return nil
}

func (l *AddLogic) createAdminUser(tenant *model.SysTenant, req *types.ModifyTenantReq) error {
	q := l.svcCtx.Query
	tenantPackage, err := q.SysTenantPackage.WithContext(l.ctx).Where(q.SysTenantPackage.PackageID.Eq(tenant.PackageID)).First()
	if err != nil {
		return errx.GORMErr(err)
	}
	//新增租户角色
	menuIds := strings.Split(tenantPackage.MenuIds, ",")
	roleId := utils.GetID()
	if err = role.NewAddLogic(l.ctx, l.svcCtx).Add(&types.AddOrUpdateRoleReq{
		RoleBase: types.RoleBase{
			RoleID:    roleId,
			RoleName:  "系统管理员",
			RoleKey:   "admin",
			RoleSort:  1,
			Status:    "0",
			Remark:    tenant.CompanyName,
			DataScope: "1",
			TenantID:  tenant.TenantID,
		},
		MenuIds: menuIds,
	}); err != nil {
		return err
	}
	//新增部门
	deptId := utils.GetID()
	err = dept.NewAddLogic(l.ctx, l.svcCtx).Add(&types.ModifyDeptReq{
		DeptBase: types.DeptBase{
			DeptID:   deptId,
			DeptName: tenant.CompanyName,
			Status:   "0",
			TenantID: tenant.TenantID,
			Phone:    tenant.ContactPhone,
			ParentID: "0",
			OrderNum: 1,
		},
	})
	//新增用户
	if err := user.NewAddUserLogic(l.ctx, l.svcCtx).AddUser(&types.AddOrUpdateUserReq{
		UserBase: types.UserBase{
			UserName: req.UserName,
			NickName: fmt.Sprintf("%s管理员", tenant.CompanyName),
			Password: req.Password,
			Status:   "0",
			TenantID: tenant.TenantID,
			DeptID:   deptId,
		},
		RoleIds: []string{roleId},
	}); err != nil {
		return err
	}
	return nil
}
