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

	"github.com/saltbo/zpan/internal/app/entity"
)

func newRecycleBin(db *gorm.DB, opts ...gen.DOOption) recycleBin {
	_recycleBin := recycleBin{}

	_recycleBin.recycleBinDo.UseDB(db, opts...)
	_recycleBin.recycleBinDo.UseModel(&entity.RecycleBin{})

	tableName := _recycleBin.recycleBinDo.TableName()
	_recycleBin.ALL = field.NewAsterisk(tableName)
	_recycleBin.Id = field.NewInt64(tableName, "id")
	_recycleBin.Uid = field.NewInt64(tableName, "uid")
	_recycleBin.Sid = field.NewInt64(tableName, "sid")
	_recycleBin.Mid = field.NewInt64(tableName, "mid")
	_recycleBin.Alias_ = field.NewString(tableName, "alias")
	_recycleBin.Name = field.NewString(tableName, "name")
	_recycleBin.Type = field.NewString(tableName, "type")
	_recycleBin.Size = field.NewInt64(tableName, "size")
	_recycleBin.DirType = field.NewInt8(tableName, "dirtype")
	_recycleBin.Parent = field.NewString(tableName, "parent")
	_recycleBin.Object = field.NewString(tableName, "object")
	_recycleBin.CreatedAt = field.NewTime(tableName, "created_at")
	_recycleBin.DeletedAt = field.NewField(tableName, "deleted_at")

	_recycleBin.fillFieldMap()

	return _recycleBin
}

type recycleBin struct {
	recycleBinDo

	ALL       field.Asterisk
	Id        field.Int64
	Uid       field.Int64
	Sid       field.Int64
	Mid       field.Int64
	Alias_    field.String
	Name      field.String
	Type      field.String
	Size      field.Int64
	DirType   field.Int8
	Parent    field.String
	Object    field.String
	CreatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r recycleBin) Table(newTableName string) *recycleBin {
	r.recycleBinDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r recycleBin) As(alias string) *recycleBin {
	r.recycleBinDo.DO = *(r.recycleBinDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *recycleBin) updateTableName(table string) *recycleBin {
	r.ALL = field.NewAsterisk(table)
	r.Id = field.NewInt64(table, "id")
	r.Uid = field.NewInt64(table, "uid")
	r.Sid = field.NewInt64(table, "sid")
	r.Mid = field.NewInt64(table, "mid")
	r.Alias_ = field.NewString(table, "alias")
	r.Name = field.NewString(table, "name")
	r.Type = field.NewString(table, "type")
	r.Size = field.NewInt64(table, "size")
	r.DirType = field.NewInt8(table, "dirtype")
	r.Parent = field.NewString(table, "parent")
	r.Object = field.NewString(table, "object")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *recycleBin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *recycleBin) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 13)
	r.fieldMap["id"] = r.Id
	r.fieldMap["uid"] = r.Uid
	r.fieldMap["sid"] = r.Sid
	r.fieldMap["mid"] = r.Mid
	r.fieldMap["alias"] = r.Alias_
	r.fieldMap["name"] = r.Name
	r.fieldMap["type"] = r.Type
	r.fieldMap["size"] = r.Size
	r.fieldMap["dirtype"] = r.DirType
	r.fieldMap["parent"] = r.Parent
	r.fieldMap["object"] = r.Object
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r recycleBin) clone(db *gorm.DB) recycleBin {
	r.recycleBinDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r recycleBin) replaceDB(db *gorm.DB) recycleBin {
	r.recycleBinDo.ReplaceDB(db)
	return r
}

type recycleBinDo struct{ gen.DO }

type IRecycleBinDo interface {
	gen.SubQuery
	Debug() IRecycleBinDo
	WithContext(ctx context.Context) IRecycleBinDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecycleBinDo
	WriteDB() IRecycleBinDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecycleBinDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecycleBinDo
	Not(conds ...gen.Condition) IRecycleBinDo
	Or(conds ...gen.Condition) IRecycleBinDo
	Select(conds ...field.Expr) IRecycleBinDo
	Where(conds ...gen.Condition) IRecycleBinDo
	Order(conds ...field.Expr) IRecycleBinDo
	Distinct(cols ...field.Expr) IRecycleBinDo
	Omit(cols ...field.Expr) IRecycleBinDo
	Join(table schema.Tabler, on ...field.Expr) IRecycleBinDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecycleBinDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecycleBinDo
	Group(cols ...field.Expr) IRecycleBinDo
	Having(conds ...gen.Condition) IRecycleBinDo
	Limit(limit int) IRecycleBinDo
	Offset(offset int) IRecycleBinDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecycleBinDo
	Unscoped() IRecycleBinDo
	Create(values ...*entity.RecycleBin) error
	CreateInBatches(values []*entity.RecycleBin, batchSize int) error
	Save(values ...*entity.RecycleBin) error
	First() (*entity.RecycleBin, error)
	Take() (*entity.RecycleBin, error)
	Last() (*entity.RecycleBin, error)
	Find() ([]*entity.RecycleBin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RecycleBin, err error)
	FindInBatches(result *[]*entity.RecycleBin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.RecycleBin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecycleBinDo
	Assign(attrs ...field.AssignExpr) IRecycleBinDo
	Joins(fields ...field.RelationField) IRecycleBinDo
	Preload(fields ...field.RelationField) IRecycleBinDo
	FirstOrInit() (*entity.RecycleBin, error)
	FirstOrCreate() (*entity.RecycleBin, error)
	FindByPage(offset int, limit int) (result []*entity.RecycleBin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecycleBinDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recycleBinDo) Debug() IRecycleBinDo {
	return r.withDO(r.DO.Debug())
}

func (r recycleBinDo) WithContext(ctx context.Context) IRecycleBinDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recycleBinDo) ReadDB() IRecycleBinDo {
	return r.Clauses(dbresolver.Read)
}

func (r recycleBinDo) WriteDB() IRecycleBinDo {
	return r.Clauses(dbresolver.Write)
}

func (r recycleBinDo) Session(config *gorm.Session) IRecycleBinDo {
	return r.withDO(r.DO.Session(config))
}

func (r recycleBinDo) Clauses(conds ...clause.Expression) IRecycleBinDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recycleBinDo) Returning(value interface{}, columns ...string) IRecycleBinDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recycleBinDo) Not(conds ...gen.Condition) IRecycleBinDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recycleBinDo) Or(conds ...gen.Condition) IRecycleBinDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recycleBinDo) Select(conds ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recycleBinDo) Where(conds ...gen.Condition) IRecycleBinDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recycleBinDo) Order(conds ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recycleBinDo) Distinct(cols ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recycleBinDo) Omit(cols ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recycleBinDo) Join(table schema.Tabler, on ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recycleBinDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recycleBinDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recycleBinDo) Group(cols ...field.Expr) IRecycleBinDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recycleBinDo) Having(conds ...gen.Condition) IRecycleBinDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recycleBinDo) Limit(limit int) IRecycleBinDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recycleBinDo) Offset(offset int) IRecycleBinDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recycleBinDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecycleBinDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recycleBinDo) Unscoped() IRecycleBinDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recycleBinDo) Create(values ...*entity.RecycleBin) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recycleBinDo) CreateInBatches(values []*entity.RecycleBin, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recycleBinDo) Save(values ...*entity.RecycleBin) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recycleBinDo) First() (*entity.RecycleBin, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RecycleBin), nil
	}
}

func (r recycleBinDo) Take() (*entity.RecycleBin, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RecycleBin), nil
	}
}

func (r recycleBinDo) Last() (*entity.RecycleBin, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RecycleBin), nil
	}
}

func (r recycleBinDo) Find() ([]*entity.RecycleBin, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RecycleBin), err
}

func (r recycleBinDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RecycleBin, err error) {
	buf := make([]*entity.RecycleBin, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recycleBinDo) FindInBatches(result *[]*entity.RecycleBin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recycleBinDo) Attrs(attrs ...field.AssignExpr) IRecycleBinDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recycleBinDo) Assign(attrs ...field.AssignExpr) IRecycleBinDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recycleBinDo) Joins(fields ...field.RelationField) IRecycleBinDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recycleBinDo) Preload(fields ...field.RelationField) IRecycleBinDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recycleBinDo) FirstOrInit() (*entity.RecycleBin, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RecycleBin), nil
	}
}

func (r recycleBinDo) FirstOrCreate() (*entity.RecycleBin, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RecycleBin), nil
	}
}

func (r recycleBinDo) FindByPage(offset int, limit int) (result []*entity.RecycleBin, count int64, err error) {
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

func (r recycleBinDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recycleBinDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recycleBinDo) Delete(models ...*entity.RecycleBin) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recycleBinDo) withDO(do gen.Dao) *recycleBinDo {
	r.DO = *do.(*gen.DO)
	return r
}