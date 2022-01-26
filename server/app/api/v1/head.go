package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Head struct {
	headService *service.HeadService
}

func NewHead() *Head {
	return &Head{
		headService: service.NewHeadService(),
	}
}

// Get 获取单个head
func (h *Head) Get(c *gin.Context) (*response.Response, error) {
	// 1、检查登录和权限
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("未登录. ")
	}
	if !api.CheckPermission(c, "posts", "read") {
		return nil, response.Forbidden.SetMsg("没有读权限. ")
	}

	// 2、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 3、查询head
	head := &domain.Head{ID: id}
	if err := h.headService.SelectOne(head); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(head), nil
}



// List 分页获取博客
func (h *Head) List(c *gin.Context) (*response.Response, error) {
	// 1、检查登录和权限
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("未登录. ")
	}
	if !api.CheckPermission(c, "posts", "lists") {
		return nil, response.Forbidden.SetMsg("没有列表权限. ")
	}

	// 2、获取参数
	var query dto.Query
	if err := c.ShouldBindQuery(&query); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	userId, _ := c.Get("current_user_id")
	if !api.CheckAdmin(c) {
		query.UserId = 0
	} else {
		query.UserId = userId.(int)
	}
	page := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	// 3、查询博客列表
	if err := h.headService.SelectList(&page, query); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&page), nil
}
