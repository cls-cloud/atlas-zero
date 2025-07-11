// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"system/internal/dao/model"
)

func newSysUserPost(db *gorm.DB, opts ...gen.DOOption) sysUserPost {
	_sysUserPost := sysUserPost{}

	_sysUserPost.sysUserPostDo.UseDB(db, opts...)
	_sysUserPost.sysUserPostDo.UseModel(&model.SysUserPost{})

	tableName := _sysUserPost.sysUserPostDo.TableName()
	_sysUserPost.ALL = field.NewAsterisk(tableName)
	_sysUserPost.UserID = field.NewString(tableName, "user_id")
	_sysUserPost.PostID = field.NewString(tableName, "post_id")

	_sysUserPost.fillFieldMap()

	return _sysUserPost
}

// sysUserPost 用户与岗位关联表
type sysUserPost struct {
	sysUserPostDo sysUserPostDo

	ALL    field.Asterisk
	UserID field.String // 用户ID
	PostID field.String // 岗位ID

	fieldMap map[string]field.Expr
}

func (s sysUserPost) Table(newTableName string) *sysUserPost {
	s.sysUserPostDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUserPost) As(alias string) *sysUserPost {
	s.sysUserPostDo.DO = *(s.sysUserPostDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUserPost) updateTableName(table string) *sysUserPost {
	s.ALL = field.NewAsterisk(table)
	s.UserID = field.NewString(table, "user_id")
	s.PostID = field.NewString(table, "post_id")

	s.fillFieldMap()

	return s
}

func (s *sysUserPost) WithContext(ctx context.Context) *sysUserPostDo {
	return s.sysUserPostDo.WithContext(ctx)
}

func (s sysUserPost) TableName() string { return s.sysUserPostDo.TableName() }

func (s sysUserPost) Alias() string { return s.sysUserPostDo.Alias() }

func (s sysUserPost) Columns(cols ...field.Expr) gen.Columns { return s.sysUserPostDo.Columns(cols...) }

func (s *sysUserPost) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUserPost) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["post_id"] = s.PostID
}

func (s sysUserPost) clone(db *gorm.DB) sysUserPost {
	s.sysUserPostDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUserPost) replaceDB(db *gorm.DB) sysUserPost {
	s.sysUserPostDo.ReplaceDB(db)
	return s
}

type sysUserPostDo struct{ gen.DO }

func (s sysUserPostDo) Debug() *sysUserPostDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserPostDo) WithContext(ctx context.Context) *sysUserPostDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserPostDo) ReadDB() *sysUserPostDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserPostDo) WriteDB() *sysUserPostDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserPostDo) Session(config *gorm.Session) *sysUserPostDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserPostDo) Clauses(conds ...clause.Expression) *sysUserPostDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserPostDo) Returning(value interface{}, columns ...string) *sysUserPostDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserPostDo) Not(conds ...gen.Condition) *sysUserPostDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserPostDo) Or(conds ...gen.Condition) *sysUserPostDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserPostDo) Select(conds ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserPostDo) Where(conds ...gen.Condition) *sysUserPostDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserPostDo) Order(conds ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserPostDo) Distinct(cols ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserPostDo) Omit(cols ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserPostDo) Join(table schema.Tabler, on ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserPostDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserPostDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserPostDo) Group(cols ...field.Expr) *sysUserPostDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserPostDo) Having(conds ...gen.Condition) *sysUserPostDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserPostDo) Limit(limit int) *sysUserPostDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserPostDo) Offset(offset int) *sysUserPostDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserPostDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserPostDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserPostDo) Unscoped() *sysUserPostDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserPostDo) Create(values ...*model.SysUserPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserPostDo) CreateInBatches(values []*model.SysUserPost, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserPostDo) Save(values ...*model.SysUserPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserPostDo) First() (*model.SysUserPost, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserPost), nil
	}
}

func (s sysUserPostDo) Take() (*model.SysUserPost, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserPost), nil
	}
}

func (s sysUserPostDo) Last() (*model.SysUserPost, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserPost), nil
	}
}

func (s sysUserPostDo) Find() ([]*model.SysUserPost, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUserPost), err
}

func (s sysUserPostDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUserPost, err error) {
	buf := make([]*model.SysUserPost, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserPostDo) FindInBatches(result *[]*model.SysUserPost, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserPostDo) Attrs(attrs ...field.AssignExpr) *sysUserPostDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserPostDo) Assign(attrs ...field.AssignExpr) *sysUserPostDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserPostDo) Joins(fields ...field.RelationField) *sysUserPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserPostDo) Preload(fields ...field.RelationField) *sysUserPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserPostDo) FirstOrInit() (*model.SysUserPost, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserPost), nil
	}
}

func (s sysUserPostDo) FirstOrCreate() (*model.SysUserPost, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserPost), nil
	}
}

func (s sysUserPostDo) FindByPage(offset int, limit int) (result []*model.SysUserPost, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserPostDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserPostDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserPostDo) Delete(models ...*model.SysUserPost) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserPostDo) withDO(do gen.Dao) *sysUserPostDo {
	s.DO = *do.(*gen.DO)
	return s
}
