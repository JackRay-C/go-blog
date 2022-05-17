package v1

import (
	"blog/app/model/po"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	repositoryService service.RepositoryService
}

func NewRepository() *Repository {
	return &Repository{
		repositoryService: service.NewRepositoryService(),
	}
}

func (r *Repository) List(c *gin.Context) (*response.Response, error)  {
	// 1、检查权限
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("未登录. ")
	}
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, response.Forbidden.SetMsg("没有列表权限. ")
	}

	// 2、获取分页参数和查询参数
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	repository := po.Repository{}
	userId, _ := c.Get("current_user_id")
	if auth.CheckAdmin(c) {
		// 如果是管理员的话，查询所有博客
		repository.UserId = 0
	} else {
		// 否则，只查询当前用户的博客
		repository.UserId = userId.(int)
	}

	if err := r.repositoryService.ISelectAll(&p, &repository); err != nil {
		return nil, err
	}

	return response.Success(&p), nil
}