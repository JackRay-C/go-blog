package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/bo"
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Subject struct {
	Log            logger.Logger
	subjectService service.SubjectService
	postService    service.PostService
}

func NewSubject() *Subject {
	return &Subject{
		Log:            global.Log,
		subjectService: service.NewSubjectService(),
		postService:    service.NewPostService(),
	}
}

func (s *Subject) Get(c *gin.Context) (*vo.Response, error) {
	// 1、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("获取ID为【%d】的专题失败：未登录. ", id)
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "read") {
		return nil, vo.Forbidden.SetMsg("获取ID为【%d】的专题失败：没有权限. ", id)
	}

	// 4、查询专题
	userId, _ := c.Get(global.SessionUserIDKey)
	subject := &po.Subject{ID: id, UserID: userId.(int)}
	if err := s.subjectService.ISelectOne(c, subject); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(subject), nil
}

func (s *Subject) List(c *gin.Context) (*vo.Response, error) {
	// 1、获取参数
	var params *po.Subject
	if err := c.ShouldBind(&params); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("获取专题列表失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "list") {
		return nil, vo.Forbidden.SetMsg("获取专题列表失败：没有权限. ")
	}

	// 4、获取专题列表
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if err := s.subjectService.ISelectList(c, &p, params); err != nil {
		s.Log.Errorf("获取专题列表失败： %s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	s.Log.Infof("获取专题列表成功：%s", &p)
	return vo.Success(&p), nil
}

func (s *Subject) Post(c *gin.Context) (*vo.Response, error) {
	// 1、绑定参数
	var subject *po.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("新建专题失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "add") {
		return nil, vo.Forbidden.SetMsg("新建专题失败：没有权限. ")
	}

	// 4、创建专题
	if err := s.subjectService.ICreateOne(c, subject); err != nil {
		return nil, err
	}

	return vo.Success(&subject), nil
}

func (s *Subject) Delete(c *gin.Context) (*vo.Response, error) {
	// 1、绑定参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、获取当前用户
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("删除ID为【%d】的专题失败：未登录. ", id)
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "delete") {
		return nil, vo.Forbidden.SetMsg("删除ID为【%d】的专题失败：没有权限. ", id)
	}

	// 4、删除该专题
	userId, _ := c.Get(global.SessionUserIDKey)
	subject := &po.Subject{ID: id, UserID: userId.(int)}
	if err := s.subjectService.IDeleteOne(c, subject); err != nil {
		return nil, err
	}

	return vo.Success(subject), nil
}

func (s *Subject) Put(c *gin.Context) (*vo.Response, error) {
	// 1、绑定参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	var subject dto.PutSubjects
	if err := c.ShouldBindJSON(&subject); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("更新专题失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "update") {
		return nil, vo.Forbidden.SetMsg("更新专题失败：没有权限. ")
	}

	// 4、更新专题
	updateSubject := &po.Subject{ID: id}

	if err := s.subjectService.IUpdateOne(c, &subject , updateSubject);err != nil {
		return nil, err
	}

	return vo.Success(updateSubject), nil
}

func (s *Subject) GetPosts(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		s.Log.Errorf("参数绑定错误：%s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	post := &bo.Post {Head: &po.Head{SubjectID: id}}
	if err := c.ShouldBind(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckLogin(c) {
		post.Head.Visibility = global.Public
		post.Head.Status = global.Publish
		if err := s.postService.ISelectListWeb(c, &p, post); err != nil {
			return nil, err
		}
		return vo.Success(p), nil
	}

	// 如果登录检查是否有查询所有博客的权限
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, vo.Forbidden
	}

	userId, _ := c.Get(global.SessionUserIDKey)
	post.Head.UserID = userId.(int)

	if err := s.postService.ISelectList(c, &p, post); err != nil {
		return nil, err
	}

	return vo.Success(&p), nil
}
