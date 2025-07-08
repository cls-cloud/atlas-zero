package _type

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageDictTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageDictTypeListLogic {
	return &QueryPageDictTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageDictTypeListLogic) QueryPageDictTypeList(req *types.QueryPageDictTypeListReq) (resp *types.QueryPageDictTypeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
