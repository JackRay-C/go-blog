package console

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Head struct {
	headService service.HeadService
}

func NewHead() *Head {
	return &Head{
		headService: service.NewHeadService(),
	}
}

// Get 获取单个head
func (h *Head) Get(c *gin.Context) (*vo.Response, error) {
	// 1、检查登录和权限
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("未登录. ")
	}
	if !auth.CheckPermission(c, "posts", "read") {
		return nil, vo.Forbidden.SetMsg("没有读权限. ")
	}

	// 2、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	head := &po.Head{ID: id}
	// 3、获取当前用户ID，判断是否是管理员
	userId, _ := c.Get("current_user_id")

	// 否则，只查询当前用户的博客
	head.UserID = userId.(int)

	// 4、查询head
	if err := h.headService.ISelectOne(c, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(head), nil
}

// List 分页获取博客
func (h *Head) List(c *gin.Context) (*vo.Response, error) {
	// 1、检查登录和权限
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("未登录. ")
	}
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, vo.Forbidden.SetMsg("没有列表权限. ")
	}

	// 2、获取参数
	var head po.Head
	if err := c.ShouldBindQuery(&head); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 3、获取当前用户ID，判断是否是管理员
	userId, _ := c.Get("current_user_id")
	head.UserID = userId.(int)

	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	// 3、查询博客列表
	if err := h.headService.ISelectList(c, &p, &head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

// Post 增加一个head
func (h *Head) Post(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "posts", "add") {
		return nil, vo.Forbidden
	}

	// 3、获取参数并结构化
	var head *po.Head
	userID, _ := c.Get("current_user_id")
	if err := c.ShouldBindJSON(&head); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	head.UserID = userID.(int)

	// 4、新增
	if err := h.headService.ICreateOne(c, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Success(head), nil
}

// Put 修改head信息
func (h *Head) Put(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}
	// 2、检查是否有更新权限
	if !auth.CheckPermission(c, "posts", "update") {
		return nil, vo.Forbidden
	}

	// 3、获取信息
	var head *po.Head
	userID, _ := c.Get("current_user_id")
	if err := c.ShouldBindJSON(&head); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	head.UserID = userID.(int)

	// 4、更新
	if err := h.headService.IUpdateOne(c, head, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(head), nil
}

// Delete 删除单条记录信息
func (h *Head) Delete(c *gin.Context) (*vo.Response, error) {
	// 1、检查登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、检查是否有更新权限
	if !auth.CheckPermission(c, "posts", "delete") {
		return nil, vo.Forbidden
	}

	// 3、获取信息
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 获取当前用户ID
	userId, _ := c.Get("current_user_id")
	head := &po.Head{ID: id, UserID: userId.(int)}

	// 4、删除
	if err := h.headService.IDeleteOne(c, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(nil), nil
}
