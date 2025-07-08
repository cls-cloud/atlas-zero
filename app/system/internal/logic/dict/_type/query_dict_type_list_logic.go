package _type

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeListLogic {
	return &QueryDictTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictTypeListLogic) QueryDictTypeList(req *types.QueryDictTypeListReq) (resp []types.DictTypeDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
