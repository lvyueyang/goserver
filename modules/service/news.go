package service

import (
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/types"
	"server/utils"
	"strconv"
	"strings"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return new(NewsService)
}

type FindNewsListOption struct {
	types.Pagination
	types.Order
	Keyword string `json:"keyword" form:"keyword"`
}

func (s *NewsService) FindList(query FindNewsListOption) (utils.ListResult[[]*model.News], error) {
	result := utils.ListResult[[]*model.News]{}
	n := dao.News
	q := n.Omit(n.Content).Where(
		n.Title.Like("%" + query.Keyword + "%"),
	)
	if id, err := strconv.ParseUint(query.Keyword, 10, 64); err == nil {
		q = q.Or(n.ID.Eq(uint(id)))
	}

	if query.OrderKey != "" {
		col, _ := n.GetFieldByName(query.OrderKey)
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

func (s *NewsService) FindDetail(id uint) (*model.News, error) {
	current, err := dao.News.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("新闻不存在", err, nil)
	}
	return current, err
}

func (s *NewsService) Create(info model.News) (*model.News, error) {
	oldData, err := dao.News.Where(dao.News.Title.Eq(info.Title)).Take()

	if err == nil {
		return oldData, errs.CreateServerError("新闻名称重复", err, nil)
	}

	data := &model.News{
		Cover:     info.Cover,
		Title:     info.Title,
		Desc:      info.Desc,
		Content:   info.Content,
		PushDate:  info.PushDate,
		Recommend: info.Recommend,
		AuthorID:  info.AuthorID,
	}
	if err := dao.News.Create(data); err != nil {
		return oldData, err
	}

	return data, nil
}

func (s *NewsService) Update(id uint, info model.News) (*model.News, error) {
	current, err := dao.News.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("新闻不存在", err, nil)
	}
	data := &model.News{
		Cover:     info.Cover,
		Title:     info.Title,
		Desc:      info.Desc,
		Content:   info.Content,
		PushDate:  info.PushDate,
		Recommend: info.Recommend,
	}

	if _, err := dao.News.Where(dao.News.ID.Eq(id)).Updates(&data); err != nil {
		return current, err
	}
	return data, nil
}

func (s *NewsService) Delete(id uint) error {
	if _, err := dao.News.FindByID(id); err != nil {
		return errs.CreateServerError("新闻不存在", err, nil)
	}

	if _, err := dao.News.Where(dao.News.ID.Eq(id)).Delete(); err != nil {
		return err
	}
	return nil
}
