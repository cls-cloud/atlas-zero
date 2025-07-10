package _post

import (
	"context"
	"system/internal/dao/model"
	"toolkit/errx"
	"toolkit/utils"

	"system/internal/svc"
	"system/internal/types"

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

func (l *AddLogic) Add(req *types.ModifyPostReq) error {
	id := utils.GetID()
	post := &model.SysPost{
		PostID:       id,
		PostCode:     req.PostCode,
		PostName:     req.PostName,
		PostSort:     req.PostSort,
		DeptID:       req.DeptID,
		Status:       req.Status,
		PostCategory: req.PostCategory,
		Remark:       req.Remark,
	}
	if err := l.svcCtx.Query.SysPost.WithContext(l.ctx).Create(post); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
