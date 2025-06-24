package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(req *types.IdReq) (resp *types.UserDetailResp, err error) {
	userId := l.ctx.Value("user_id")
	l.Logger.Info("当前查询数据userId:", userId)
	resp = new(types.UserDetailResp)
	sysUser := l.svcCtx.Query.SysUser
	user, err := sysUser.WithContext(l.ctx).Where(sysUser.UserID.Eq(req.Id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	err = copier.Copy(resp, user)
	if err != nil {
		return nil, err
	}
	return
}
