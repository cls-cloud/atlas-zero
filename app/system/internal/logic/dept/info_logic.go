package dept

import (
	"context"
	"github.com/jinzhu/copier"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.IdReq) (resp *types.DeptBase, err error) {
	resp = new(types.DeptBase)
	q := l.svcCtx.Query
	sysPost, err := q.SysDept.WithContext(l.ctx).Where(q.SysDept.DeptID.Eq(req.Id)).First()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	err = copier.Copy(&resp, sysPost)
	if sysPost.Leader != "" {
		resp.Leader = &sysPost.Leader
	} else {
		resp.Leader = nil
	}
	return
}
