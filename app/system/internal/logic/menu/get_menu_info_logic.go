package menu

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuInfoLogic {
	return &GetMenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuInfoLogic) GetMenuInfo() (resp *types.MenuInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
