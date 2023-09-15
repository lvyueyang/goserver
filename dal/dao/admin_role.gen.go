// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"server/dal/model"
)

func newAdminRole(db *gorm.DB, opts ...gen.DOOption) adminRole {
	_adminRole := adminRole{}

	_adminRole.adminRoleDo.UseDB(db, opts...)
	_adminRole.adminRoleDo.UseModel(&model.AdminRole{})

	tableName := _adminRole.adminRoleDo.TableName()
	_adminRole.ALL = field.NewAsterisk(tableName)
	_adminRole.ID = field.NewUint(tableName, "id")
	_adminRole.CreatedAt = field.NewTime(tableName, "created_at")
	_adminRole.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminRole.DeletedAt = field.NewField(tableName, "deleted_at")
	_adminRole.Name = field.NewString(tableName, "name")
	_adminRole.Desc = field.NewString(tableName, "desc")
	_adminRole.PermissionCode = field.NewField(tableName, "permission_code")

	_adminRole.fillFieldMap()

	return _adminRole
}

type adminRole struct {
	adminRoleDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Name           field.String
	Desc           field.String
	PermissionCode field.Field

	fieldMap map[string]field.Expr
}

func (a adminRole) Table(newTableName string) *adminRole {
	a.adminRoleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminRole) As(alias string) *adminRole {
	a.adminRoleDo.DO = *(a.adminRoleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminRole) updateTableName(table string) *adminRole {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Name = field.NewString(table, "name")
	a.Desc = field.NewString(table, "desc")
	a.PermissionCode = field.NewField(table, "permission_code")

	a.fillFieldMap()

	return a
}

func (a *adminRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminRole) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["name"] = a.Name
	a.fieldMap["desc"] = a.Desc
	a.fieldMap["permission_code"] = a.PermissionCode
}

func (a adminRole) clone(db *gorm.DB) adminRole {
	a.adminRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminRole) replaceDB(db *gorm.DB) adminRole {
	a.adminRoleDo.ReplaceDB(db)
	return a
}

type adminRoleDo struct{ gen.DO }

type IAdminRoleDo interface {
	gen.SubQuery
	Debug() IAdminRoleDo
	WithContext(ctx context.Context) IAdminRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminRoleDo
	WriteDB() IAdminRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminRoleDo
	Not(conds ...gen.Condition) IAdminRoleDo
	Or(conds ...gen.Condition) IAdminRoleDo
	Select(conds ...field.Expr) IAdminRoleDo
	Where(conds ...gen.Condition) IAdminRoleDo
	Order(conds ...field.Expr) IAdminRoleDo
	Distinct(cols ...field.Expr) IAdminRoleDo
	Omit(cols ...field.Expr) IAdminRoleDo
	Join(table schema.Tabler, on ...field.Expr) IAdminRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleDo
	Group(cols ...field.Expr) IAdminRoleDo
	Having(conds ...gen.Condition) IAdminRoleDo
	Limit(limit int) IAdminRoleDo
	Offset(offset int) IAdminRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleDo
	Unscoped() IAdminRoleDo
	Create(values ...*model.AdminRole) error
	CreateInBatches(values []*model.AdminRole, batchSize int) error
	Save(values ...*model.AdminRole) error
	First() (*model.AdminRole, error)
	Take() (*model.AdminRole, error)
	Last() (*model.AdminRole, error)
	Find() ([]*model.AdminRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRole, err error)
	FindInBatches(result *[]*model.AdminRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminRoleDo
	Assign(attrs ...field.AssignExpr) IAdminRoleDo
	Joins(fields ...field.RelationField) IAdminRoleDo
	Preload(fields ...field.RelationField) IAdminRoleDo
	FirstOrInit() (*model.AdminRole, error)
	FirstOrCreate() (*model.AdminRole, error)
	FindByPage(offset int, limit int) (result []*model.AdminRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminRoleDo) Debug() IAdminRoleDo {
	return a.withDO(a.DO.Debug())
}

func (a adminRoleDo) WithContext(ctx context.Context) IAdminRoleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminRoleDo) ReadDB() IAdminRoleDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminRoleDo) WriteDB() IAdminRoleDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminRoleDo) Session(config *gorm.Session) IAdminRoleDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminRoleDo) Clauses(conds ...clause.Expression) IAdminRoleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminRoleDo) Returning(value interface{}, columns ...string) IAdminRoleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminRoleDo) Not(conds ...gen.Condition) IAdminRoleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminRoleDo) Or(conds ...gen.Condition) IAdminRoleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminRoleDo) Select(conds ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminRoleDo) Where(conds ...gen.Condition) IAdminRoleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminRoleDo) Order(conds ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminRoleDo) Distinct(cols ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminRoleDo) Omit(cols ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminRoleDo) Join(table schema.Tabler, on ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminRoleDo) Group(cols ...field.Expr) IAdminRoleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminRoleDo) Having(conds ...gen.Condition) IAdminRoleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminRoleDo) Limit(limit int) IAdminRoleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminRoleDo) Offset(offset int) IAdminRoleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminRoleDo) Unscoped() IAdminRoleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminRoleDo) Create(values ...*model.AdminRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminRoleDo) CreateInBatches(values []*model.AdminRole, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminRoleDo) Save(values ...*model.AdminRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminRoleDo) First() (*model.AdminRole, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRole), nil
	}
}

func (a adminRoleDo) Take() (*model.AdminRole, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRole), nil
	}
}

func (a adminRoleDo) Last() (*model.AdminRole, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRole), nil
	}
}

func (a adminRoleDo) Find() ([]*model.AdminRole, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminRole), err
}

func (a adminRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRole, err error) {
	buf := make([]*model.AdminRole, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminRoleDo) FindInBatches(result *[]*model.AdminRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminRoleDo) Attrs(attrs ...field.AssignExpr) IAdminRoleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminRoleDo) Assign(attrs ...field.AssignExpr) IAdminRoleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminRoleDo) Joins(fields ...field.RelationField) IAdminRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminRoleDo) Preload(fields ...field.RelationField) IAdminRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminRoleDo) FirstOrInit() (*model.AdminRole, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRole), nil
	}
}

func (a adminRoleDo) FirstOrCreate() (*model.AdminRole, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRole), nil
	}
}

func (a adminRoleDo) FindByPage(offset int, limit int) (result []*model.AdminRole, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminRoleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminRoleDo) Delete(models ...*model.AdminRole) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminRoleDo) withDO(do gen.Dao) *adminRoleDo {
	a.DO = *do.(*gen.DO)
	return a
}
