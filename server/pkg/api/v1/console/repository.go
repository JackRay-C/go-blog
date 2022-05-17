package console

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
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

func (r *Repository) List(c *gin.Context) (*vo.Response, error)  {
	// 1、检查权限
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("未登录. ")
	}
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, vo.Forbidden.SetMsg("没有列表权限. ")
	}

	// 2、获取分页参数和查询参数
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
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

	if err := r.repositoryService.ISelectList(c, &p, &repository); err != nil {
		return nil, err
	}

	return vo.Success(&p), nil
}