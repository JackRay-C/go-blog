package v1

//import (
//	"blog/app/api"
//	"blog/app/model/dto"
//	"blog/app/model/po"
//	"blog/app/pager"
//	"blog/app/request"
//	"blog/app/response"
//	"blog/app/service"
//	"blog/core/global"
//	"blog/core/logs"
//	"github.com/gin-gonic/gin"
//	"strconv"
//)
//
//type Post struct {
//	logs         logs.Logger
//	postService service.PostService
//}
//
//func NewPost() *Post {
//	return &Post{
//		logs:         global.Logger,
//		postService: service.NewPostService(),
//	}
//}
//
//// Get 查询单条博客
//func (p *Post) Get(c *gin.Context) (*response.Response, error) {
//	// 1、获取参数
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	// 2、判断是否登录
//	if !api.CheckLogin(c) {
//		return nil, response.NotLogin.SetMsg("获取ID为【%d】的博客失败：未登录. ", id)
//	}
//
//	// 3、判断是否有权限
//	if !api.CheckPermission(c, "posts", "read") {
//		return nil, response.Forbidden.SetMsg("获取ID为【%d】的博客失败：没有权限. ", id)
//	}
//
//	userId, _ := c.Get("current_user_id")
//
//	// 4、查询一条记录
//	post := &po.Post{ID: id, UserId: userId.(int)}
//	if err := p.postService.ISelectOne(c, post); err != nil {
//		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的博客失败： %s", id, err)
//	}
//
//	return response.Success(post), nil
//}
//
//// List 查询当前用户的博客、包括已发布、草稿、公开、私有的
//func (p *Post) List(c *gin.Context) (*response.Response, error) {
//	// 1、获取参数
//	post := dto.ListPosts{}
//	if err := c.ShouldBind(&post); err != nil {
//		p.logs.Errorf("绑定参数错误: error: %s", err)
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//
//	// 2、检查是否登录
//	if !api.CheckLogin(c) {
//		return nil, response.NotLogin.SetMsg("未登录. ")
//	}
//
//	// 2、检查权限，是否有post-list的权限
//	if !api.CheckPermission(c,"posts", "list") {
//		return nil, response.Forbidden.SetMsg("没有权限. ")
//	}
//
//	// 3、检查权限，如果有posts-list权限，则查询所有博客，否则只查询当前用户的所有博客
//	userId, _ := c.Get("current_user_id")
//	post.UserId = userId.(int)
//
//	// 4、查询博客
//	page := pager.Pager{
//		PageNo:   request.GetPageNo(c),
//		PageSize: request.GetPageSize(c),
//	}
//	if err := p.postService.ISelectAll(c, &page); err != nil {
//		p.logs.Errorf("分页查询博客失败: %s", err)
//		return nil, response.InternalServerError.SetMsg("分页查询博客失败：%s", err)
//	}
//
//	// 5、返回查询结果
//	p.logs.Infof("分页查询博客成功: [第 %d 页，总页数：%d, 总行数：%d]", page.PageNo, page.PageCount, page.TotalRows)
//	return response.Success(&page), nil
//}
//
//func (p *Post) Post(c *gin.Context) (*response.Response, error) {
//	params := po.Post{}
//
//	// 1、绑定参数
//	if err := c.ShouldBindJSON(&params); err != nil {
//		p.logs.Errorf("新建博客失败: 参数绑定错误: %s", err)
//		return nil, response.InvalidParams.SetMsg("新建博客失败：%s. ", err)
//	}
//
//	// 2、获取当前用户
//	if !api.CheckLogin(c) {
//		return nil, response.NotLogin.SetMsg("新建博客失败：未登录. ")
//	}
//	if currentUserId, ok := c.Get("current_user_id"); ok {
//		params.UserId = currentUserId.(int)
//	}
//
//	// 3、判断是否有权限
//	if !api.CheckPermission(c, "posts", "add") {
//		return nil, response.Forbidden.SetMsg("新建博客失败：没有权限. ")
//	}
//
//	// 4、创建博客
//	if err := p.postService.ICreateOne(c, &params); err != nil {
//		p.logs.Errorf("新建博客失败：error: %s", err)
//		return nil, response.InternalServerError.SetMsg("新建博客失败：error: %s", err)
//	} else {
//		p.logs.Errorf("新建博客【%d】成功! ", params.ID)
//		return response.Success(params), nil
//	}
//}
//
//func (p *Post) Delete(c *gin.Context) (*response.Response, error) {
//	// 1、获取参数
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("删除ID为【%d】的博客失败： %s ", id, err)
//	}
//
//	// 2、判断是否有权限
//	if !api.CheckPermission(c, "posts", "delete") {
//		return nil, response.Forbidden.SetMsg("删除ID为【%d】的博客失败：没有权限. ", id)
//	}
//
//	// 3、删除博客
//	if err := p.postService.IDeleteOne(c, &po.Post{ID: id}); err != nil {
//		return nil, response.InternalServerError.SetMsg("删除ID为【%d】的博客失败: %s", id, err)
//	}
//	p.logs.Infof("删除博客【%d】成功.", id)
//	return response.Success("delete success. "), nil
//}
//
//// 更新博客
//func (p *Post) Put(c *gin.Context) (*response.Response, error) {
//	params := po.Post{}
//
//	// 1、获取要更新博客的ID
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//
//	// 2、获取当前用户id
//	currentUserId, _ := c.Get("current_user_id")
//	params.UserId = currentUserId.(int)
//	params.ID = id
//
//	// 3、判断是否有权限
//	if !api.CheckPermission(c, "posts", "update") {
//		return nil, response.Forbidden.SetMsg("更新博客【%d】失败：没有权限", id)
//	}
//
//	// 4、绑定json到结构体上
//	if err := c.ShouldBindJSON(&params); err != nil {
//		return nil, response.InvalidParams.SetMsg("更新博客【%d】失败：%s", id, err)
//	}
//
//	// 5、更新博客
//	if err := p.postService.IUpdateOne(c,&params); err != nil {
//		return nil, response.InternalServerError.SetMsg("更新博客【%d】失败：%s", id, err)
//	} else {
//		return response.Success("post"), nil
//	}
//}