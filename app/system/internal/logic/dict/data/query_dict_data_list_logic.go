package data

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictDataListLogic {
	return &QueryDictDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictDataListLogic) QueryDictDataList(req *types.QueryDictDataListReq) (resp []types.DictDataDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
