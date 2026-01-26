package menu

import (
	"context"
	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"
	"ovra/toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IdReq) error {
	dal := l.svcCtx.Dal
	if req.Id == "" {
		return nil
	}
	isExist, err := dal.SysMenuDal.ExistChildMenu(l.ctx, []string{req.Id})
	if err != nil {
		return err
	}
	if isExist {
		return errx.BizErr("存在子级菜单，请先删除子级菜单")
	}
	if err := dal.SysMenuDal.Delete(l.ctx, req.Id); err != nil {
		return err
	}
	return nil
}
