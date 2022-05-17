package v1

import (
	"blog/app/model/dto"
	"blog/app/model/vo"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Subject struct {
	log            logger.Logger
	subjectService *service.SubjectService
	postService    service.PostService
}

func NewSubject() *Subject {
	return &Subject{
		log:            global.Logger,
		subjectService: service.NewSubjectService(),
		postService:    service.NewPostService(),
	}
}

func (s *Subject) Get(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("获取ID为【%d】的专题失败：未登录. ",id)
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "read") {
		return nil, response.Forbidden.SetMsg("获取ID为【%d】的专题失败：没有权限. ", id)
	}

	// 4、查询专题
	if v, err := s.subjectService.SelectOneById(id); err != nil {
		s.log.Errorf("查询ID为【%d】的专题失败 : %s",id, err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	} else {
		return response.Success(v), nil
	}
}

func (s *Subject) List(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	params := dto.ListSubjects{}
	if err := c.ShouldBind(&params); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("获取专题列表失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "list") {
		return nil, response.Forbidden.SetMsg("获取专题列表失败：没有权限. ")
	}

	// 4、获取专题列表
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if err := s.subjectService.SelectAll(c, &p, &params); err != nil {
		s.log.Errorf("获取专题列表失败： %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	s.log.Infof("获取专题列表成功：%s", &p)
	return response.Success(&p), nil
}

func (s *Subject) Post(c *gin.Context) (*response.Response, error) {
	// 1、绑定参数
	subject := &dto.AddSubjects{}
	if err := c.ShouldBindJSON(&subject); err != nil {
		s.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("新建专题失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "add") {
		return nil, response.Forbidden.SetMsg("新建专题失败：没有权限. ")
	}

	// 4、创建专题
	one, err := s.subjectService.CreateOne(c, subject)
	if err != nil {
		s.log.Errorf("新建专题失败：error: %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(one), nil
}

func (s *Subject) Delete(c *gin.Context) (*response.Response, error) {
	// 1、绑定参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、获取当前用户
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("删除ID为【%d】的专题失败：未登录. ", id)
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "delete") {
		return nil, response.Forbidden.SetMsg("删除ID为【%d】的专题失败：没有权限. ", id)
	}

	// 4、删除该专题
	if err := s.subjectService.DeleteOne(c, id); err != nil {
		s.log.Infof("删除ID为【%d】的专题失败: %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success("delete success"), nil
}

func (s *Subject) Put(c *gin.Context) (*response.Response, error) {
	// 1、绑定参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var subject dto.PutSubjects
	if err := c.ShouldBindJSON(&subject); err != nil {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判读是否登录
	if !auth.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("更新专题失败：未登录. ")
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "subjects", "update") {
		return nil, response.Forbidden.SetMsg("更新专题失败：没有权限. ")
	}

	// 4、更新专题
	one, err := s.subjectService.SaveOne(c, &subject)
	if err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	byId, _ := s.subjectService.SelectOneById(one.ID)
	return response.Success(byId), nil
}

func (s *Subject) GetPosts(c *gin.Context) (*response.Response, error) {
	p := vo.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	post := dto.ListPosts{SubjectId: id}
	if isLogin, ok := c.Get("is_login"); ok {
		if !isLogin.(bool) {
			post.Visibility = 2
		}
	}
	if err := s.postService.ISelectAll(c, &p); err!=nil {
		return nil, err
	}

	s.log.Infof("根据专题查询博客成功")
	return response.Success(&p), nil
}

