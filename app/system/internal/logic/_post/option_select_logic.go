package _post

import (
	"context"
	"github.com/jinzhu/copier"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionSelectLogic {
	return &OptionSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionSelectLogic) OptionSelect(req *types.PostQuery) (resp []types.DeptBase, err error) {
	resp = make([]types.DeptBase, 0)
	q := l.svcCtx.Query
	posts, err := q.SysPost.WithContext(l.ctx).Where(q.SysPost.DeptID.Eq(req.DeptId)).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	err = copier.Copy(&resp, posts)
	if err != nil {
		return nil, err
	}
	return
}
