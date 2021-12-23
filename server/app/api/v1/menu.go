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

//func (m *Menu) Get(c *gin.Context) (*response.Response, error) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	var menus []*domain.Menu
//	roles, exists := c.Get("current_user_roles")
//	if exists {
//		r := roles.([]*domain.Role)
//		// 获取角色的所有菜单
//		if err := m.roleMenuService.SelectMenusByRoles(&menus, r...); err != nil {
//			return nil, err
//		}
//	}
//
//	for _, menu := range menus {
//		if menu.ID == id {
//			return response.Success(menu), nil
//		}
//	}
//
//	return nil, response.RecordNotFound.SetMsg("查询菜单信息失败：没有该菜单. ")
//}
//
//func (m *Menu) List(c *gin.Context) (*response.Response, error) {
//	p := pager.Pager{
//		PageNo:   request.GetPageNo(c),
//		PageSize: request.GetPageSize(c),
//	}
//
//	var menus []*domain.Menu
//	roles, exists := c.Get("current_user_roles")
//	if exists {
//		r := roles.([]*domain.Role)
//		// 获取角色的所有菜单
//		if err := m.roleMenuService.SelectMenusByRoles(&menus, r...); err != nil {
//			return nil, err
//		}
//	}
//
//	p.TotalRows = int64(len(menus))
//	p.PageCount = int((p.TotalRows + int64(p.PageSize) - 1) / int64(p.PageSize))
//
//	var start,end int
//
//	if (p.PageNo-1)*p.PageSize > p.PageCount {
//		start = (p.PageCount - 1) * p.PageSize
//		p.PageNo = p.PageCount
//	}  else {
//		start = (p.PageNo-1)*p.PageSize
//	}
//
//	if (start + p.PageSize) <= len(menus) {
//		end = start + p.PageSize
//	} else {
//		end = len(menus)
//	}
//
//	p.List = menus[start:end]
//	return response.Success(&p), nil
//}

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
