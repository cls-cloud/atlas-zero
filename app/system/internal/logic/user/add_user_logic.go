package user

import (
	"context"
	"system/internal/dao/model"
	"time"
	"toolkit/errx"
	"toolkit/helper"
	"toolkit/utils"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddOrUpdateUserReq) error {
	if req.UserID != "" {
		return errx.BizErr("用户ID不为空")
	}
	if req.Password == "" {
		return errx.BizErr("密码不能为空")
	}
	req.UserID = utils.GetID()
	user := &model.SysUser{
		UserID:      req.UserID,
		UserName:    req.UserName,
		Password:    helper.Encrypt(req.Password),
		NickName:    req.NickName,
		Phonenumber: req.PhoneNumber,
		Email:       req.Email,
		Status:      req.Status,
		DeptID:      req.DeptID,
		Avatar:      req.Avatar,
		LoginIP:     req.LoginIp,
		UserType:    "00",
		LoginDate:   time.Now(),
		Remark:      req.Remark,
		Sex:         req.Sex,
	}
	if req.TenantID != "" {
		user.TenantID = req.TenantID
	}
	q := l.svcCtx.Query
	if err := q.SysUser.WithContext(l.ctx).Create(user); err != nil {
		return errx.GORMErr(err)
	}
	if len(req.RoleIds) != 0 {
		userRole := make([]*model.SysUserRole, len(req.RoleIds))
		for i, roleId := range req.RoleIds {
			userRole[i] = &model.SysUserRole{
				UserID: req.UserID,
				RoleID: roleId,
			}
		}
		if err := q.SysUserRole.WithContext(l.ctx).CreateInBatches(userRole, len(userRole)); err != nil {
			return errx.GORMErr(err)
		}
	}
	if len(req.PostIds) != 0 {
		userPost := make([]*model.SysUserPost, len(req.PostIds))
		for i, postId := range req.PostIds {
			userPost[i] = &model.SysUserPost{
				UserID: req.UserID,
				PostID: postId,
			}
		}
		if err := q.SysUserPost.WithContext(l.ctx).CreateInBatches(userPost, len(userPost)); err != nil {
			return errx.GORMErr(err)
		}
	}

	return nil
}
