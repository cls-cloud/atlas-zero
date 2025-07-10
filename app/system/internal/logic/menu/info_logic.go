package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"strconv"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.IdReq) (resp *types.MenuBase, err error) {
	resp = new(types.MenuBase)
	q := l.svcCtx.Query
	sysMenu, err := q.SysMenu.WithContext(l.ctx).Where(q.SysMenu.MenuID.Eq(req.Id)).First()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	if err = copier.Copy(&resp, sysMenu); err != nil {
		return nil, err
	}
	resp.IsFrame = strconv.Itoa(int(sysMenu.IsFrame))
	resp.IsCache = strconv.Itoa(int(sysMenu.IsCache))
	return
}
