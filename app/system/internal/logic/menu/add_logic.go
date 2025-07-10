package menu

import (
	"context"
	"strconv"
	"system/internal/dao/model"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"
	"toolkit/utils"

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

func (l *AddLogic) Add(req *types.ModifyMenuReq) error {
	isFrame, _ := strconv.ParseInt(req.IsFrame, 10, 64)
	isCache, _ := strconv.ParseInt(req.IsCache, 10, 64)
	menu := &model.SysMenu{
		MenuID:    utils.GetID(),
		MenuName:  req.MenuName,
		ParentID:  req.ParentID,
		OrderNum:  req.OrderNum,
		Path:      req.Path,
		Component: req.Component,
		IsFrame:   int32(isFrame),
		IsCache:   int32(isCache),
		MenuType:  req.MenuType,
		Visible:   req.Visible,
		Status:    req.Status,
		Perms:     req.Perms,
		Icon:      req.Icon,
	}
	q := l.svcCtx.Query
	if err := q.SysMenu.WithContext(l.ctx).Create(menu); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
