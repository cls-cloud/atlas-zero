package user

import (
	"context"
	"github.com/jinzhu/copier"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req *types.IdReq) (resp []*types.UserBase, err error) {
	q := l.svcCtx.Query
	sysUsers, err := q.SysUser.WithContext(l.ctx).
		LeftJoin(q.SysDept, q.SysDept.DeptID.EqCol(q.SysUser.DeptID)).
		Where(q.SysDept.DeptID.Eq(req.Id)).
		Order(q.SysUser.CreateTime.Asc()).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	resp = make([]*types.UserBase, len(sysUsers))
	for i, sysUser := range sysUsers {
		resp[i] = new(types.UserBase)
		if err = copier.Copy(&resp[i], sysUser); err != nil {
			return nil, err
		}
		resp[i].UserID = sysUser.UserID
		resp[i].DeptID = sysUser.DeptID
	}
	return
}
