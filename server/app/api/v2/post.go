package v2
//
//import (
//	"blog/app/api"
//	"blog/app/domain"
//	"blog/app/model/po"
//	"blog/app/model/vo"
//	"blog/app/response"
//	"blog/app/service"
//	"blog/app/utils/auth"
//	"blog/core/global"
//	"blog/core/logs"
//	"github.com/gin-gonic/gin"
//	"strconv"
//	"time"
//)
//
//type Post struct {
//	logs               logs.Logger
//	postService       service.PostService
//	repositoryService service.RepositoryService
//	headService       service.HeadService
//	historyService    service.HistoryService
//}
//
//func NewPost() *Post {
//	return &Post{
//		logs:               global.Logger,
//		postService:       service.NewPostService(),
//		repositoryService: service.NewRepositoryService(),
//		headService:       service.NewHeadService(),
//		historyService:    service.NewHistoryService(),
//	}
//}
//
//// Initialize
//// 创建博客并初始化一个博客repository和history
//// example
///*
//POST /posts HTTP/1.0
//
//{
//
//		"id": 0,
//		"title": "",
//		"markdown_content": "",
//		"html_content": "",
//		"description": "",
//		"user_id": 0,
//		"image_ids": "",
//		"created_at": "",
//		"updated_at": "",
//		"deleted_at": ""
//
//}
//
//Response
//{
//	"head": {
//		"id": 1,
//		"repository_id": 1,
//		"visibility": 2,
//		"status": 1,
//		"subject_id": 0,
//		"cover_image_id": 0,
//		"user_id": 1,
//		"created_at": "2022-01-10 15:00:00",
//		"updated_at": "2022-01-10 15:00:00",
//		"deleted_at": "null"
//	},
//	"history": [
//		{
//			"head_id": 1,
//			"repository_id": 1,
//			"staged_at": "2022-01-10 15:00:00"
//			"commited_at": "2022-01-10 15:00:00",
//			"published_at": "2022-01-10 15:0:00"
//		}
//	],
//	"repository": {}
//}
//*/
//func (p *Post) Initialize(c *gin.Context) (*response.Response, error) {
//	if !auth.CheckLogin(c) {
//		return nil, response.NotLogin.SetMsg("未登录. ")
//	}
//	if !auth.CheckPermission(c, "posts", "add") {
//		return nil, response.Forbidden.SetMsg("没有新建博客的权限. ")
//	}
//	userId, _ := c.Get("current_user_id")
//
//	// 1、保存repository，获取ID
//	var repository *po.Repository
//	if err := c.ShouldBindJSON(&repository); err != nil {
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//
//	// 2、创建repository
//	repository.UserId = userId.(int)
//	if err := p.repositoryService.CreateOne(repository); err != nil {
//		return nil, response.InternalServerError.SetMsg("%s", err)
//	}
//
//	// 3、创建head头，指向repository id
//	now := time.Now()
//	head := &domain.Head{
//		ID:           0,
//		RepositoryID: repository.ID,
//		Visibility:   2,
//		Status:       1,
//		SubjectID:    0,
//		CoverImageId: 0,
//		UserID:       userId.(int),
//		CreatedAt:    now,
//		UpdatedAt:    now,
//	}
//	if err := p.headService.CreateOne(head); err != nil {
//		return nil, response.InternalServerError.SetMsg("%s", err)
//	}
//
//	// 4、保存history，设置为已提交
//	history := &domain.History{
//		HeadID:           head.ID,
//		RepositoryID:     repository.ID,
//		PrevRepositoryID: 0,
//		StagedAt:         now,
//		CommitedAt:       now,
//	}
//	if err := p.historyService.CreateOne(history); err != nil {
//		return nil, response.InternalServerError.SetMsg("%s", err)
//	}
//
//	posts := &vo.VPosts{
//
//	}
//
//	return response.Success(posts), nil
//}
//
//// Pull 根据博客ID返回当前博客的repository
///**
//GET /1 HTTP/1.1
//
//Response
//{
//	"head": {
//		"id": 0,
//		"repository_id": 0,
//		"visibility": 2,
//		"status": 0,
//		"subject_id": 0,
//		"user_id": 0,
//		"tags":[
//			{
//				"id": 1
//			},
//			{
//				"id": 2
//			}
//		]
//		"created_at": "2022-01-01 10:00:01",
//		"updated_at": "",
//		"deleted_at": "null"
//	},
//	"history": [
//		{},
//		{}
//	],
//	"repository": {
//
//	}
//}
//*/
//func (p *Post) Pull(c *gin.Context) (*response.Response, error) {
//	// 1、检查登录及权限
//	if !api.CheckLogin(c) {
//		return nil, response.NotLogin.SetMsg("未登录. ")
//	}
//	if !api.CheckPermission(c, "posts", "read") {
//		return nil, response.Forbidden.SetMsg("没有读权限. ")
//	}
//
//	// 2、获取参数
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	// 3、根据ID查询Head
//	head := &domain.Head{ID: id}
//	if err := p.headService.SelectOne(head); err != nil {
//		return nil, response.InternalServerError.SetMsg("%s", err)
//	}
//
//	// 4、根据id查询history
//	//histories := make([]*domain.History, 0)
//
//
//	// 5、根据head查询当前repository
//
//
//	// 6、组装数据返回
//	post := &vo.VPost{
//
//	}
//
//	return response.Success(post), nil
//}
//
//// Staged 暂存博客
//func (p *Post) Staged(c *gin.Context) {
//	// 只提交到repository，不保存记录，不发布
//}
//
//// Commit
//// parse post data from gin.Context then save to database
//// return commit data
//func (p *Post) Commit(c *gin.Context) (*response.Response, error) {
//	// 只提交到repository和history中，不发布
//
//	return response.Success(""), nil
//}
//
//// Publish 发布博客
//func (p *Post) Publish(c *gin.Context) (*response.Response, error) {
//	// 获取博客ID及RepositoryID，将博客当前的ID指向RepositoryID，并保存到history表中
//
//	return response.Success(""), nil
//}
//
//// Repositories 获取所有博客元信息
//func (p *Post) Repositories(c *gin.Context) (*response.Response, error)  {
//	// 获取所有博客的愿信息
//
//	return response.Success(""), nil
//}