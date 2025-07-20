package config

import (
	"context"
	"fmt"
	"strings"
	"toolkit/constants"
	"toolkit/errx"

	"resource/internal/svc"
	"resource/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IdsReq) error {
	ids := strings.Split(req.Ids, ",")
	q := l.svcCtx.Query
	configs, err := q.SysOssConfig.WithContext(l.ctx).Find()
	if err != nil {
		return errx.GORMErr(err)
	}
	for _, config := range configs {
		_, err := l.svcCtx.Rds.DelCtx(l.ctx, fmt.Sprintf(constants.OssConfigCache, config.OssConfigID))
		if err != nil {
			return err
		}
	}
	_, err = l.svcCtx.Rds.DelCtx(l.ctx, constants.OssConfigDefaultCache)
	if err != nil {
		return err
	}
	if _, err := q.SysOssConfig.WithContext(l.ctx).Where(q.SysOssConfig.OssConfigID.In(ids...)).Delete(); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
