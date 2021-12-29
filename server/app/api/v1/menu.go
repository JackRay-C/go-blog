package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Menu struct {
	log             logger.Logger
	menuService     *service.MenuService
	roleMenuService *service.RolesMenusService
}

func NewMenu() *Menu {
	return &Menu{
		log:             global.Logger,
		menuService:     service.NewMenuService(),
		roleMenuService: service.NewRolesMenusService(),
	}
}

func (m *Menu) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID错误. ")
	}

	if !api.CheckPermission(c, "menus", "read") {
		return nil, response.Forbidden.SetMsg("查询菜单信息失败： 没有权限. ")
	}

	menu := &domain.Menu{}
	if err := m.menuService.SelectOne(menu); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(menu), err
}

func (m *Menu) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if !api.CheckPermission(c, "menus", "list") {
		return nil, response.Forbidden.SetMsg("查询菜单列表失败： 没有权限. ")
	}

	if err := m.menuService.SelectAll(&p, &domain.Menu{}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&p), nil
}

func (m *Menu) Post(c *gin.Context) (*response.Response, error) {
	m.log.Infof("新建菜单")

	menu := &domain.Menu{}
	if err := c.ShouldBindJSON(&menu); err != nil {
		m.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !api.CheckPermission(c, "menus", "add") {
		return nil, response.Forbidden.SetMsg("新建菜单失败：没有权限. ")
	}

	if err := m.menuService.CreateOne(menu); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	m.log.Infof("新建菜单成功")
	return response.Success(menu), nil
}

func (m *Menu) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if !api.CheckPermission(c, "menus", "delete") {
		return nil, response.Forbidden.SetMsg("删除菜单失败：没有权限. ")
	}

	menu := &domain.Menu{ID: id}
	if err := m.menuService.DeleteOne(menu); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	m.log.Infof("删除菜单成功!")
	return response.Success("删除成功"), nil
}

func (m *Menu) Put(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var menu domain.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	menu.ID = id

	if !api.CheckPermission(c, "menus", "update") {
		return nil, response.Forbidden.SetMsg("更新菜单信息失败： 没有权限. ")
	}

	if err := m.menuService.SaveOne(&menu); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	m.log.Infof("修改路由信息成功.")
	return response.Success(menu), nil
}
