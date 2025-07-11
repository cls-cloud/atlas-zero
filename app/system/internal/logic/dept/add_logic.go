package dept

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

func (l *AddLogic) Add(req *types.ModifyDeptReq) error {
	q := l.svcCtx.Query
	var ancestors string
	parentId := req.ParentID
	//获取祖级列表
	if parentId != "0" {
		if err := q.SysDept.WithContext(l.ctx).Select(q.SysDept.Ancestors).Where(q.SysDept.DeptID.Eq(parentId)).Scan(&ancestors); err != nil {
			return errx.GORMErr(err)
		}
		ancestors = ancestors + "," + req.ParentID
	} else {
		ancestors = req.ParentID
	}
	dept := &model.SysDept{
		DeptID:    utils.GetID(),
		DeptName:  req.DeptName,
		Ancestors: ancestors,
		ParentID:  parentId,
		OrderNum:  req.OrderNum,
		Phone:     req.Phone,
		Email:     req.Email,
		Status:    req.Status,
	}
	if req.TenantID != "" {
		dept.TenantID = req.TenantID
	}
	if err := q.SysDept.WithContext(l.ctx).Create(dept); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
