package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"system/internal/dao/model"
	"system/internal/svc"
	"system/internal/types"
	"time"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageUserListLogic {
	return &QueryPageUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageUserListLogic) QueryPageUserList(req *types.QueryPageUserListReq) (resp *types.QueryPageUserListResp, err error) {
	resp = new(types.QueryPageUserListResp)

	sysUser := l.svcCtx.Query.SysUser
	sysDept := l.svcCtx.Query.SysDept

	do := sysUser.WithContext(l.ctx).
		LeftJoin(sysDept, sysDept.DeptID.EqCol(sysUser.DeptID)).
		Select(
			sysUser.ALL,
			sysDept.DeptName, // 添加 dept_name 字段
		)

	// 条件查询
	if req.DeptId != 0 {
		deptIds, err := GetDeptIds(req, l.svcCtx.Db, l.ctx)
		if err != nil {
			return nil, errx.GORMErr(err)
		}
		do = do.Where(sysUser.DeptID.In(deptIds...))
	}
	if req.UserName != "" {
		do = do.Where(sysUser.UserName.Like("%" + req.UserName + "%"))
	}
	if req.NickName != "" {
		do = do.Where(sysUser.NickName.Like("%" + req.NickName + "%"))
	}
	if req.PhoneNumber != "" {
		do = do.Where(sysUser.Phonenumber.Like("%" + req.PhoneNumber + "%"))
	}
	if req.Status != "" {
		do = do.Where(sysUser.Status.Eq(req.Status))
	}
	if req.BeginTime != "" {
		beginTime, err := time.Parse(time.DateTime, req.BeginTime)
		if err != nil {
			return nil, errors.New("invalid beginTime format, expected: YYYY-MM-DD HH:mm:ss")
		}
		do = do.Where(sysUser.CreateTime.Gte(beginTime))
	}
	if req.EndTime != "" {
		endTime, err := time.Parse(time.DateTime, req.EndTime)
		if err != nil {
			return nil, errors.New("invalid endTime format, expected: YYYY-MM-DD HH:mm:ss")
		}
		do = do.Where(sysUser.CreateTime.Lte(endTime))
	}

	// 分页
	total, err := do.Count()
	if err != nil {
		return nil, err
	}

	offset := (req.PageNum - 1) * req.PageSize
	var result []struct {
		model.SysUser
		DeptName string `gorm:"column:dept_name"`
	}
	err = do.Offset(int(offset)).Limit(int(req.PageSize)).Scan(&result)
	if err != nil {
		return nil, err
	}

	// 复制并添加 dept_name
	resp.Total = total
	resp.Rows = make([]types.UserAndDeptName, len(result))
	for i, row := range result {
		var userDetail types.UserAndDeptName
		_ = copier.Copy(&userDetail, row.SysUser)
		userDetail.CreateTime = row.CreateTime.Format(time.DateTime)
		userDetail.DeptName = row.DeptName
		resp.Rows[i] = userDetail
	}

	return
}

func GetDeptIds(req *types.QueryPageUserListReq, db *gorm.DB, ctx context.Context) ([]int64, error) {
	var deptIds []int64
	err := db.WithContext(ctx).
		Table("sys_dept").
		Select("dept_id").
		Where("FIND_IN_SET(?, ancestors)", req.DeptId).
		Find(&deptIds).Error
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	deptIds = append(deptIds, req.DeptId)
	return deptIds, nil
}
