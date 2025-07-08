package data

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageDictDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageDictDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageDictDataListLogic {
	return &QueryPageDictDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageDictDataListLogic) QueryPageDictDataList(req *types.QueryPageDictDataListReq) (resp *types.QueryPageDictDataListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
