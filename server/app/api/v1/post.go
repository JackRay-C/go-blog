package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Post struct {
	log         logger.Logger
	postService *service.PostService
}

func NewPost() *Post {
	return &Post{
		log:         global.Logger,
		postService: service.NewPostService(),
	}
}


func (p *Post) Get(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("获取ID为【%d】的博客失败：未登录. ",id)
	}

	// 3、判断是否有权限
	if !api.CheckPermission(c, "posts", "read") {
		return nil, response.Forbidden.SetMsg("获取ID为【%d】的博客失败：没有权限. ", id)
	}

	userId, _ := c.Get("current_user_id")
	v, err := p.postService.SelectOne(&domain.Post{ID: id, UserId: userId.(int)})
	if err != nil {
		p.log.Errorf("查询ID为【%d】的博客失败： %s", id, err)
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的博客失败： %s", id, err)
	}
	return response.Success(v), nil
}


// 查询当前用户的博客、包括已发布、草稿、公开、私有的
func (p *Post) List(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	post := dto.ListPosts{}
	if err := c.ShouldBind(&post); err != nil {
		p.log.Errorf("绑定参数错误: error: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、检查权限，如果有posts-list权限，则查询所有博客，否则只查询当前用户的所有博客
	userId, _ := c.Get("current_user_id")
	if post.UserId == 0 {
		if api.CheckAdmin(c) || api.CheckPermission(c, "posts", "list") {
			post.UserId = 0
		} else {
			post.UserId = userId.(int)
		}
	}
	if post.UserId != userId {
		if !api.CheckAdmin(c) || !api.CheckPermission(c, "posts", "list") {
			// 判断是否是管理员，不是管理员，判断是否有posts-list权限，没有的话，返回错误
			return nil, response.Forbidden.SetMsg("查询博客失败：没有权限. ")
		}
	}

	// 3、查询博客
	page := pager.Pager{}
	if err := p.postService.SelectAll(&page, &post); err != nil {
		p.log.Errorf("分页查询博客失败: %s", err)
		return nil, response.InternalServerError.SetMsg("分页查询博客失败：%s", err)
	}

	// 4、返回查询结果
	p.log.Infof("分页查询博客成功: [第 %d 页，总页数：%d, 总行数：%d]", page.PageNo, page.PageCount, page.TotalRows)
	return response.PagerResponse(&page), nil
}

func (p *Post) Post(c *gin.Context) (*response.Response, error) {
	params := dto.AddPosts{}

	// 1、绑定参数
	if err := c.ShouldBindJSON(&params); err != nil {
		p.log.Errorf("新建博客失败: 参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("新建博客失败：%s. ", err)
	}

	// 2、获取当前用户
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("新建博客失败：未登录. ")
	}
	if currentUserId, ok := c.Get("current_user_id"); ok {
		params.UserId = currentUserId.(int)
	}

	// 3、判断是否有权限
	if !api.CheckPermission(c, "posts", "add") {
		return nil, response.Forbidden.SetMsg("新建博客失败：没有权限. ")
	}

	// 4、创建博客
	if post, err := p.postService.CreateOne(&params); err != nil {
		p.log.Errorf("新建博客失败：error: %s", err)
		return nil, response.InternalServerError.SetMsg("新建博客失败：error: %s", err)
	} else {
		p.log.Errorf("新建博客【%d】成功! ", post.ID)
		return response.Success(post), nil
	}
}

func (p *Post) Delete(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("删除ID为【%d】的博客失败： %s ", id, err)
	}

	// 2、获取当前用户
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("删除ID为【%d】的博客失败：未登录. ", id)
	}

	// 3、判断是否有权限
	if !api.CheckPermission(c, "posts", "delete") {
		return nil, response.Forbidden.SetMsg("删除ID为【%d】的博客失败：没有权限. ", id)
	}

	// 4、删除博客
	if err := p.postService.DeleteOne(c, id); err != nil {
		return nil, response.InternalServerError.SetMsg("删除ID为【%d】的博客失败: %s", id, err)
	}
	p.log.Infof("删除博客【%d】成功.", id)
	return response.Success("delete success. "), nil
}

// 更新博客
func (p *Post) Put(c *gin.Context) (*response.Response, error) {
	params := dto.PutPosts{}

	// 1、获取要更新博客的ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	p.log.Infof("更新博客【%d】", id)

	// 2、获取当前用户
	if !api.CheckLogin(c) {
		p.log.Infof("更新博客【%d】失败：未登录", id)
		return nil, response.NotLogin.SetMsg("更新博客【%d】失败：未登录", id)
	}
	if currentUserId, ok := c.Get("current_user_id"); ok {
		params.UserId = currentUserId.(int)
	}

	// 3、判断是否有权限
	if !api.CheckPermission(c, "posts", "update") {
		p.log.Infof("更新博客【%d】失败：没有权限", id)
		return nil, response.Forbidden.SetMsg("更新博客【%d】失败：没有权限", id)
	}

	// 4、绑定body参数
	if err := c.ShouldBindJSON(&params); err != nil {
		p.log.Infof("更新博客【%d】失败：%s", id, err)
		return nil, response.InvalidParams.SetMsg("更新博客【%d】失败：%s", id, err)
	}

	// 根据路由上传递的ID，而不是传递的ID
	params.ID = id
	if post, err := p.postService.UpdateOne(&params); err != nil {
		p.log.Infof("更新博客【%d】失败：%s", id, err)
		return nil, response.InternalServerError.SetMsg("更新博客【%d】失败：%s", id, err)
	} else {
		p.log.Infof("更新博客【%d】成功", id)
		return response.Success(post), nil
	}
}

func (p *Post) GetComments(c *gin.Context) (*response.Response, error) {
	// 获取一个博客下所有的评论
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	p.log.Infof("查询博客 [%d] 的所有评论", id)
	var comments []*domain.Comment
	if err := p.postService.SelectPostComments(&domain.Post{ID: id}, &comments); err != nil {
		return nil, err
	}

	p.log.Infof("查询评论成功: %s", comments)
	return response.Success(&comments), nil
}

func (p *Post) PostComments(c *gin.Context) (*response.Response, error) {
	// 添加评论
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if comment.PostId == 0 {
		comment.PostId = id
	}

	// 判断是否登录，如果登录，设置userid，否则必须有昵称和邮箱
	if currentUserId, ok := c.Get("current_user_id"); ok {
		comment.UserID = currentUserId.(int)
	} else {
		p.log.Infof("%s", comment)
		if comment.Nickname == "" || comment.Email == "" {
			return nil, response.InvalidParams.SetMsg("昵称和邮箱是必须的.")
		}
	}

	p.log.Infof("给博客 [%d] 的添加评论: %s", id, comment)

	if err := p.postService.InsertPostComment(&domain.Post{ID: id}, &comment); err != nil {
		p.log.Errorf("给博客 [%d] 的添加评论失败: %s ", id, err)
		return nil, err
	}
	p.log.Infof("添加评论成功: %s", &comment)
	return response.Success(comment), nil
}

func (p *Post) PutComments(c *gin.Context) (*response.Response, error) {
	p.log.Infof("修改评论")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 判断是否登录，如果登录，设置userid, 没有则返回登录
	if currentUserId, ok := c.Get("current_user_id"); ok {
		comment.UserID = currentUserId.(int)
	} else {
		return nil, response.NotLogin
	}

	comment.ID = commentId
	comment.PostId = id

	p.log.Infof("%s", &comment)
	if err := p.postService.UpdatePostComment(&comment); err != nil {
		return nil, err
	}
	p.log.Infof("修改评论成功. ")

	return response.Success("success"), nil
}

func (p *Post) DeleteComments(c *gin.Context) (*response.Response, error) {
	p.log.Infof("修改评论")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		p.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 判断是否登录 没有则返回登录
	if _, ok := c.Get("current_user_id"); !ok {
		return nil, response.NotLogin
	}

	if err := p.postService.DeletePostComment(&domain.Comment{ID: commentId, PostId: id}); err != nil {
		return nil, err
	}

	return response.Success("delete success"), nil
}
