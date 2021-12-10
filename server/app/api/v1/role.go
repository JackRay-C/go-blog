package v1

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Role struct {
	log           logger.Logger
	roleService *service.RoleService
}

func NewRole() *Role {
	return &Role{
		log: global.Logger,
		roleService: service.NewRoleService(),
	}
}

func (r *Role) Get(c *gin.Context) (*response.Response, error) {
	r.log.Infof("根据ID查询角色信息")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	role := &domain.Role{ID: id}

	if err := r.roleService.SelectOne(role); err != nil {
		r.log.Errorf("查询ID为【%d】的角色信息失败：%s",id, err)
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的角色信息失败：%s",id, err)
	}

	r.log.Infof("根据ID查询角色成功: %s", role)
	return response.Success(role), nil
}

func (r *Role) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if err := r.roleService.SelectAll(&p); err != nil {
		r.log.Errorf("分页查询角色失败： %s", err)
		return nil, err
	}

	r.log.Infof("分页查询角色成功：%s", &p)
	return response.PagerResponse(&p), nil
}

func (r *Role) Post(c *gin.Context) (*response.Response, error) {
	r.log.Infof("新建角色")

	role := &domain.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		r.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.CreateOne(role); err != nil {
		r.log.Errorf("新建角色失败：error: %s", err)
		return nil, err
	}

	r.log.Infof("新建角色成功")
	return response.Success(role), nil
}

func (r *Role) Delete(c *gin.Context) (*response.Response, error) {
	r.log.Infof("删除路由")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	role := &domain.Role{ID: id}
	if err := r.roleService.DeleteOne(role); err != nil {
		r.log.Errorf("删除出错： %s", err)
		return nil, err
	}
	r.log.Infof("删除路由成功!")
	return response.Success("删除成功"), nil
}

func (r *Role) Patch(c *gin.Context) (*response.Response, error) {
	r.log.Infof("修改角色信息.")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	role.ID = id
	if err := r.roleService.UpdateOne(&role); err != nil {
		return nil, err
	}

	r.log.Infof("修改角色信息成功.")
	return  response.Success(role), err
}

func (r *Role) Put(c *gin.Context) (*response.Response, error) {
	r.log.Infof("修改角色所有信息")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	role.ID = id

	if err := r.roleService.UpdateOne(&role); err != nil {
		return nil, err
	}
	r.log.Infof("修改角色信息成功.")
	return response.Success(role), nil
}

func (r *Role) GetMenus(c *gin.Context) (*response.Response, error)  {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.SelectMenus(&p, &domain.Role{ID: id}); err != nil {
		return nil, err
	}

	r.log.Infof("获取角色菜单地址成功.")
	return response.Success(p), nil
}

func (r *Role) PostMenus(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestMenus dto.AddRoleMenus
	if err := c.ShouldBindJSON(&requestMenus); err != nil {
		r.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.UpdateMenus(&domain.Role{ID: id}, requestMenus.Menus); err != nil {
		return nil, response.InternalServerError.SetMsg("授权ID为【%d】的角色菜单失败：%s", id, err)
	}

	r.log.Infof("为角色添加菜单成功")
	return response.Success(requestMenus),nil
}

//func (r *Role) DeleteMenus(c *gin.Context) (*response.Response, error) {
//	r.log.Infof("删除角色菜单")
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//	menuId, err := strconv.Atoi(c.Param("menu_id"))
//	if err != nil {
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//
//	if err := r.roleService.DeleteRoleMenus(&domain.Role{ID: id}, &domain.Menu{ID: menuId}); err != nil {
//		return nil, err
//	}
//	return response.Success("删除成功. "), nil
//}

func (r *Role) PutMenus(c *gin.Context) (*response.Response, error) {
	r.log.Infof("修改角色菜单")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestMenus dto.AddRoleMenus
	if err := c.ShouldBindJSON(&requestMenus); err != nil {
		r.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.UpdateMenus(&domain.Role{ID: id}, requestMenus.Menus); err != nil {
		return nil, err
	}
	return response.Success(requestMenus), nil
}

