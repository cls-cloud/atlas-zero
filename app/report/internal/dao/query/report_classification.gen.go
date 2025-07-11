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

	"report/internal/dao/model"
)

func newReportClassification(db *gorm.DB, opts ...gen.DOOption) reportClassification {
	_reportClassification := reportClassification{}

	_reportClassification.reportClassificationDo.UseDB(db, opts...)
	_reportClassification.reportClassificationDo.UseModel(&model.ReportClassification{})

	tableName := _reportClassification.reportClassificationDo.TableName()
	_reportClassification.ALL = field.NewAsterisk(tableName)
	_reportClassification.UUID = field.NewString(tableName, "uuid")
	_reportClassification.CreatorID = field.NewString(tableName, "creator_id")
	_reportClassification.Creator = field.NewString(tableName, "creator")
	_reportClassification.CreateTime = field.NewTime(tableName, "create_time")
	_reportClassification.ModifierID = field.NewString(tableName, "modifier_id")
	_reportClassification.Modifier = field.NewString(tableName, "modifier")
	_reportClassification.ModifyTime = field.NewTime(tableName, "modify_time")
	_reportClassification.TenantUUID = field.NewString(tableName, "tenant_uuid")
	_reportClassification.CreatorOrgID = field.NewInt32(tableName, "creator_org_id")
	_reportClassification.ClassificationName = field.NewString(tableName, "classification_name")
	_reportClassification.ClassificationCode = field.NewString(tableName, "classification_code")
	_reportClassification.ClassificationSort = field.NewInt32(tableName, "classification_sort")
	_reportClassification.ClassificationRemark = field.NewString(tableName, "classification_remark")

	_reportClassification.fillFieldMap()

	return _reportClassification
}

// reportClassification 数据报表分类
type reportClassification struct {
	reportClassificationDo reportClassificationDo

	ALL                  field.Asterisk
	UUID                 field.String // 主键
	CreatorID            field.String // 创建人Id
	Creator              field.String // 创建人
	CreateTime           field.Time   // 创建时间
	ModifierID           field.String // 修改人Id
	Modifier             field.String // 修改人
	ModifyTime           field.Time   // 修改时间
	TenantUUID           field.String // 租户UUID
	CreatorOrgID         field.Int32  // 创建人组织机构ID
	ClassificationName   field.String // 分类名称
	ClassificationCode   field.String // 分类编码
	ClassificationSort   field.Int32  // 排序
	ClassificationRemark field.String // 备注

	fieldMap map[string]field.Expr
}

func (r reportClassification) Table(newTableName string) *reportClassification {
	r.reportClassificationDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reportClassification) As(alias string) *reportClassification {
	r.reportClassificationDo.DO = *(r.reportClassificationDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reportClassification) updateTableName(table string) *reportClassification {
	r.ALL = field.NewAsterisk(table)
	r.UUID = field.NewString(table, "uuid")
	r.CreatorID = field.NewString(table, "creator_id")
	r.Creator = field.NewString(table, "creator")
	r.CreateTime = field.NewTime(table, "create_time")
	r.ModifierID = field.NewString(table, "modifier_id")
	r.Modifier = field.NewString(table, "modifier")
	r.ModifyTime = field.NewTime(table, "modify_time")
	r.TenantUUID = field.NewString(table, "tenant_uuid")
	r.CreatorOrgID = field.NewInt32(table, "creator_org_id")
	r.ClassificationName = field.NewString(table, "classification_name")
	r.ClassificationCode = field.NewString(table, "classification_code")
	r.ClassificationSort = field.NewInt32(table, "classification_sort")
	r.ClassificationRemark = field.NewString(table, "classification_remark")

	r.fillFieldMap()

	return r
}

func (r *reportClassification) WithContext(ctx context.Context) *reportClassificationDo {
	return r.reportClassificationDo.WithContext(ctx)
}

func (r reportClassification) TableName() string { return r.reportClassificationDo.TableName() }

func (r reportClassification) Alias() string { return r.reportClassificationDo.Alias() }

func (r reportClassification) Columns(cols ...field.Expr) gen.Columns {
	return r.reportClassificationDo.Columns(cols...)
}

func (r *reportClassification) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reportClassification) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 13)
	r.fieldMap["uuid"] = r.UUID
	r.fieldMap["creator_id"] = r.CreatorID
	r.fieldMap["creator"] = r.Creator
	r.fieldMap["create_time"] = r.CreateTime
	r.fieldMap["modifier_id"] = r.ModifierID
	r.fieldMap["modifier"] = r.Modifier
	r.fieldMap["modify_time"] = r.ModifyTime
	r.fieldMap["tenant_uuid"] = r.TenantUUID
	r.fieldMap["creator_org_id"] = r.CreatorOrgID
	r.fieldMap["classification_name"] = r.ClassificationName
	r.fieldMap["classification_code"] = r.ClassificationCode
	r.fieldMap["classification_sort"] = r.ClassificationSort
	r.fieldMap["classification_remark"] = r.ClassificationRemark
}

func (r reportClassification) clone(db *gorm.DB) reportClassification {
	r.reportClassificationDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reportClassification) replaceDB(db *gorm.DB) reportClassification {
	r.reportClassificationDo.ReplaceDB(db)
	return r
}

type reportClassificationDo struct{ gen.DO }

func (r reportClassificationDo) Debug() *reportClassificationDo {
	return r.withDO(r.DO.Debug())
}

func (r reportClassificationDo) WithContext(ctx context.Context) *reportClassificationDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reportClassificationDo) ReadDB() *reportClassificationDo {
	return r.Clauses(dbresolver.Read)
}

func (r reportClassificationDo) WriteDB() *reportClassificationDo {
	return r.Clauses(dbresolver.Write)
}

func (r reportClassificationDo) Session(config *gorm.Session) *reportClassificationDo {
	return r.withDO(r.DO.Session(config))
}

func (r reportClassificationDo) Clauses(conds ...clause.Expression) *reportClassificationDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reportClassificationDo) Returning(value interface{}, columns ...string) *reportClassificationDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reportClassificationDo) Not(conds ...gen.Condition) *reportClassificationDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reportClassificationDo) Or(conds ...gen.Condition) *reportClassificationDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reportClassificationDo) Select(conds ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reportClassificationDo) Where(conds ...gen.Condition) *reportClassificationDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reportClassificationDo) Order(conds ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reportClassificationDo) Distinct(cols ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reportClassificationDo) Omit(cols ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reportClassificationDo) Join(table schema.Tabler, on ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reportClassificationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reportClassificationDo) RightJoin(table schema.Tabler, on ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reportClassificationDo) Group(cols ...field.Expr) *reportClassificationDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reportClassificationDo) Having(conds ...gen.Condition) *reportClassificationDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reportClassificationDo) Limit(limit int) *reportClassificationDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reportClassificationDo) Offset(offset int) *reportClassificationDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reportClassificationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *reportClassificationDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reportClassificationDo) Unscoped() *reportClassificationDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reportClassificationDo) Create(values ...*model.ReportClassification) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reportClassificationDo) CreateInBatches(values []*model.ReportClassification, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reportClassificationDo) Save(values ...*model.ReportClassification) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reportClassificationDo) First() (*model.ReportClassification, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportClassification), nil
	}
}

func (r reportClassificationDo) Take() (*model.ReportClassification, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportClassification), nil
	}
}

func (r reportClassificationDo) Last() (*model.ReportClassification, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportClassification), nil
	}
}

func (r reportClassificationDo) Find() ([]*model.ReportClassification, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReportClassification), err
}

func (r reportClassificationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReportClassification, err error) {
	buf := make([]*model.ReportClassification, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reportClassificationDo) FindInBatches(result *[]*model.ReportClassification, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reportClassificationDo) Attrs(attrs ...field.AssignExpr) *reportClassificationDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reportClassificationDo) Assign(attrs ...field.AssignExpr) *reportClassificationDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reportClassificationDo) Joins(fields ...field.RelationField) *reportClassificationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reportClassificationDo) Preload(fields ...field.RelationField) *reportClassificationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reportClassificationDo) FirstOrInit() (*model.ReportClassification, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportClassification), nil
	}
}

func (r reportClassificationDo) FirstOrCreate() (*model.ReportClassification, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportClassification), nil
	}
}

func (r reportClassificationDo) FindByPage(offset int, limit int) (result []*model.ReportClassification, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reportClassificationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reportClassificationDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reportClassificationDo) Delete(models ...*model.ReportClassification) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reportClassificationDo) withDO(do gen.Dao) *reportClassificationDo {
	r.DO = *do.(*gen.DO)
	return r
}
