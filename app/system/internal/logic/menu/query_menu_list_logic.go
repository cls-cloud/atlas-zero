package menu

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuListLogic {
	return &QueryMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *QueryMenuListLogic) QueryMenuList(req *types.QueryMenuListReq) (resp []types.MenuDetailResp, err error) {
	sysMenu := l.svcCtx.Query.SysMenu
	do := sysMenu.WithContext(l.ctx)

	if req.MenuId != 0 {
		do = do.Where(sysMenu.MenuID.Eq(req.MenuId))
	}
	menus, err := do.Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("菜单不存在")
		}
		return nil, err
	}
	err = copier.Copy(&resp, menus)
	if err != nil {
		return nil, err
	}
	return
}
