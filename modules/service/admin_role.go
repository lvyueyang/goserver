package service

import (
	"server/dal/dao"
	"server/dal/dbtypes"
	"server/dal/model"
	"server/lib/errs"
	"server/types"
	"server/utils"
	"strings"
)

type AdminRoleService struct {
}

func NewAdminRoleService() *AdminRoleService {
	return new(AdminRoleService)
}

type FindAdminRoleListOption struct {
	types.Pagination
	types.Order
	Keyword string `json:"keyword" form:"keyword"`
}

func (s *AdminRoleService) FindList(query FindAdminRoleListOption) (utils.ListResult[[]*model.AdminRole], error) {
	result := utils.ListResult[[]*model.AdminRole]{}
	q := dao.AdminRole.Where(
		dao.AdminRole.Name.Like("%" + query.Keyword + "%"),
	)

	if query.OrderKey != "" {
		col, _ := dao.AdminRole.GetFieldByName(query.OrderKey)
		if strings.ToLower(query.OrderType) == "desc" {
			q = q.Order(col.Desc())
		} else {
			q = q.Order(col)
		}
	}

	if list, total, err := q.FindByPage(utils.PageTrans(query.Pagination)); err != nil {
		return result, err
	} else {
		result.List = list
		result.Total = total
	}

	return result, nil
}

func (s *AdminRoleService) Create(info model.AdminRole) (*model.AdminRole, error) {
	oldData, err := dao.AdminRole.Where(dao.AdminRole.Name.Eq(info.Name)).Or(dao.AdminRole.Code.Eq(info.Code)).Take()

	if err == nil {
		return oldData, errs.CreateServerError("角色名称/编码重复", err, nil)
	}

	data := &model.AdminRole{
		Name:            info.Name,
		Code:            info.Code,
		Desc:            info.Desc,
		PermissionCodes: info.PermissionCodes,
	}
	if err := dao.AdminRole.Create(data); err != nil {
		return oldData, err
	}

	return data, nil
}

func (s *AdminRoleService) Update(id uint, info model.AdminRole) (*model.AdminRole, error) {
	current, err := dao.AdminRole.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("角色不存在", err, nil)
	}
	data := &model.AdminRole{
		Name: info.Name,
		Code: info.Code,
		Desc: info.Desc,
	}

	if _, err := dao.AdminRole.Where(dao.AdminRole.ID.Eq(id)).Updates(&data); err != nil {
		return current, err
	}
	return data, nil
}

func (s *AdminRoleService) UpdatePermissionCode(id uint, codes dbtypes.StringArray) (*model.AdminRole, error) {
	current, err := dao.AdminRole.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("角色不存在", err, nil)
	}
	data := &model.AdminRole{
		PermissionCodes: codes,
	}

	if _, err := dao.AdminRole.Where(dao.AdminRole.ID.Eq(id)).Updates(&data); err != nil {
		return current, err
	}
	return data, nil
}

func (s *AdminRoleService) Delete(id uint) error {
	if _, err := dao.AdminRole.FindByID(id); err != nil {
		return errs.CreateServerError("角色不存在", err, nil)
	}

	if _, err := dao.AdminRole.Where(dao.AdminRole.ID.Eq(id)).Delete(); err != nil {
		return err
	}
	return nil
}
