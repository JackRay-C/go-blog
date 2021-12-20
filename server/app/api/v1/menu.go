package v1

import (
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
	log logger.Logger
	menuService *service.MenuService
}

func NewMenu() *Menu {
	return &Menu{
		log: global.Logger,
		menuService: service.NewMenuService(),
	}
}

func (m *Menu) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	m.log.Infof("根据ID查询菜单: ID[%d]", id)
	menu := &domain.Menu{ID: id}

	if err := m.menuService.SelectOne(menu); err != nil {
		m.log.Errorf("根据ID查询菜单 : %s", err)
		return nil, err
	}

	m.log.Infof("根据ID查询菜单成功: %s", menu)
	return response.Success(menu), nil
}

func (m *Menu) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	m.log.Infof("分页查询菜单")

	if err := m.menuService.SelectAll(&p, &domain.Menu{}); err != nil {
		m.log.Errorf("分页查询菜单失败： %s", err)
		return nil, err
	}

	m.log.Infof("分页查询菜单成功：%s", &p)
	return response.Success(&p), nil
}

func (m *Menu) Post(c *gin.Context) (*response.Response, error) {
	m.log.Infof("新建菜单")

	menu := &domain.Menu{}
	if err := c.ShouldBindJSON(&menu); err != nil {
		m.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := m.menuService.CreateOne(menu); err != nil {
		m.log.Errorf("新建菜单失败：error: %s", err)
		return nil, err
	}

	m.log.Infof("新建菜单成功")
	return response.Success(menu), nil
}

func (m *Menu) Delete(c *gin.Context) (*response.Response, error) {
	m.log.Infof("删除菜单")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	menu := &domain.Menu{ID: id}
	if err := m.menuService.DeleteOne(menu); err != nil {
		m.log.Errorf("删除出错： %s", err)
		return nil, err
	}
	m.log.Infof("删除菜单成功!")
	return response.Success("删除成功"), nil
}

func (m *Menu) Patch(c *gin.Context) (*response.Response, error) {
	m.log.Infof("修改菜单信息.")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	var menu domain.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	menu.ID = id
	if err := m.menuService.UpdateOne(&menu); err != nil {
		return nil, err
	}

	m.log.Infof("修改菜单信息成功.")
	return  response.Success(menu), err
}

func (m *Menu) Put(c *gin.Context) (*response.Response, error) {
	m.log.Infof("修改菜单所有信息")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var menu domain.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	menu.ID = id

	if err := m.menuService.SaveOne(&menu); err != nil {
		return nil, err
	}
	m.log.Infof("修改路由信息成功.")
	return response.Success(menu), nil
}


