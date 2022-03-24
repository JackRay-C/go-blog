package v1

import (
	"blog/app/model/po"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"github.com/gin-gonic/gin"
)

type History struct {
	historyService *service.HistoryService
}

func NewHistory() *History {
	return &History{
		historyService: service.NewHistoryService(),
	}
}

// List 获取博客编辑的历史记录
func (h *History) List(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、判断是否有list权限
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, response.Forbidden
	}

	// 3、获取查询参数
	var history *po.History
	if err := c.ShouldBindJSON(&history); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 4、设置history的用户ID
	userId, _ := c.Get("current_user_id")
	history.UserID = userId.(int)

	// 5、查询history
	var histories []*po.History
	if err := h.historyService.SelectList(history, &histories); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	// 6、返回数据
	return response.Success(&histories), nil
}

// Post 新建历史记录
func (h *History) Post(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、判断是否有新建的权限
	if !auth.CheckPermission(c, "posts", "add") {
		return nil, response.Forbidden
	}

	// 3、绑定参数
	var history *po.History
	if err := c.ShouldBindJSON(&history); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	// 4、设置history的用户ID
	userId, _ := c.Get("current_user_id")
	history.UserID = userId.(int)


	// 4、创建
	if err := h.historyService.CreateOne(history); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	// 5、返回结果
	return response.Success(history), nil
}

// Get 获取单条历史记录
func (h *History) Get(c *gin.Context) (*response.Response, error) {
	return nil, nil
}

// Put 更新历史记录， example：提交、暂存、发布
func (h *History) Put(c *gin.Context) (*response.Response, error) {
	return nil, nil
}

func (h *History) Delete(c *gin.Context) (*response.Response, error) {
	return nil, nil
}

