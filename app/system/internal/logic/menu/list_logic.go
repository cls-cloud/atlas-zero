package menu

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"system/internal/svc"
	"system/internal/types"
	"time"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.QueryMenuListReq) (resp []*types.MenuBase, err error) {
	q := l.svcCtx.Query
	do := q.SysMenu.WithContext(l.ctx)
	if req.MenuName != "" {
		do = do.Where(q.SysMenu.MenuName.Like(fmt.Sprintf("%%%s%%", req.MenuName)))
	}
	if req.Status != "" {
		do = do.Where(q.SysMenu.Status.Eq(req.Status))
	}
	if req.Visible != "" {
		do = do.Where(q.SysMenu.Visible.Eq(req.Visible))
	}
	result, err := do.Order(q.SysMenu.OrderNum.Asc(), q.SysMenu.CreateTime.Desc()).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	resp = make([]*types.MenuBase, 0)
	list := make([]*types.MenuBase, len(result))
	for i, item := range result {
		list[i] = new(types.MenuBase)
		if err = copier.Copy(&list[i], item); err != nil {
			return nil, err
		}
		list[i].CreateTime = item.CreateTime.Format(time.DateTime)
	}
	resp = list
	return
}
