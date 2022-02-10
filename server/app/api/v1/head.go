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


	head := &domain.Head{ID: id}
	// 3、获取当前用户ID，判断是否是管理员
	userId, _ := c.Get("current_user_id")
	if api.CheckAdmin(c) {
		// 如果是管理员的话，查询所有博客
		head.UserID = 0
	} else {
		// 否则，只查询当前用户的博客
		head.UserID = userId.(int)
	}

	// 4、查询head
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
	if !api.CheckPermission(c, "posts", "list") {
		return nil, response.Forbidden.SetMsg("没有列表权限. ")
	}

	// 2、获取参数
	var query dto.Query
	if err := c.ShouldBindQuery(&query); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 3、获取当前用户ID，判断是否是管理员
	userId, _ := c.Get("current_user_id")
	if api.CheckAdmin(c) {
		// 如果是管理员的话，查询所有博客
		query.UserID = 0
	} else {
		// 否则，只查询当前用户的博客
		query.UserID = userId.(int)
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

// Post 增加一个head
func (h *Head) Post(c *gin.Context) (*response.Response, error) {
	// 1、检查是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、检查是否有权限
	if !api.CheckPermission(c, "posts", "add") {
		return nil, response.Forbidden
	}

	// 3、获取参数并结构化
	var head *domain.Head
	userID, _ := c.Get("current_user_id")
	if err := c.ShouldBindJSON(&head); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	head.UserID = userID.(int)

	// 4、新增
	if err := h.headService.CreateOne(head); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(""),nil
}

// Put 修改head信息
func (h *Head) Put(c *gin.Context) (*response.Response, error) {
	// 1、检查是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin
	}
	// 2、检查是否有更新权限
	if !api.CheckPermission(c, "posts", "update") {
		return nil, response.Forbidden
	}

	// 3、获取信息
	var head *domain.Head
	userID, _ := c.Get("current_user_id")
	if err := c.ShouldBindJSON(&head); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	head.UserID = userID.(int)

	// 4、更新

	return response.Success(""),nil
}

func (h *Head) Delete(c *gin.Context) (*response.Response, error) {
	return response.Success(""),nil
}
