package oss

import (
	"context"
	"resource/internal/dao/model"
	"toolkit/utils"

	"resource/internal/svc"
	"resource/internal/types"

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

func (l *AddLogic) Add(req *types.ModifyOssReq) error {
	_ = &model.SysOss{
		OssID:        utils.GetID(),
		FileName:     req.FileName,
		OriginalName: req.OriginalName,
		FileSuffix:   req.FileSuffix,
		URL:          req.Url,
		Ext1:         req.Ext1,
		Service:      req.Service,
	}

	return nil
}
