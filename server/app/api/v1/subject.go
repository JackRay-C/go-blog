package v1

import (
	"blog/app/api"
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

type Subject struct {
	log            logger.Logger
	subjectService *service.SubjectService
	postService    *service.PostService
}

func NewSubject() *Subject {
	return &Subject{
		log:            global.Logger,
		subjectService: service.NewSubjectService(),
		postService:    service.NewPostService(),
	}
}

func (s *Subject) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if v, err := s.subjectService.SelectOneById(id); err != nil {
		s.log.Errorf("根据ID查询专题 : %s", err)
		return nil, err
	} else {
		s.log.Infof("根据ID查询专题成功: %s", v)
		return response.Success(v), nil
	}
}

func (s *Subject) List(c *gin.Context) (*response.Response, error) {
	params := dto.ListSubjects{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	p := pager.Pager{}
	s.log.Infof("分页查询专题")
	if err := s.subjectService.SelectAll(c, &p, &params); err != nil {
		s.log.Errorf("分页查询失败： %s", err)
		return nil, err
	}

	s.log.Infof("分页查询成功：%s", &p)
	return response.PagerResponse(&p), nil
}

func (s *Subject) Post(c *gin.Context) (*response.Response, error) {
	s.log.Infof("新建专题")

	subject := &domain.Subject{}
	if err := c.ShouldBindJSON(&subject); err != nil {
		s.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if currentUserId, ok := c.Get("current_user_id"); ok {
		subject.UserID = currentUserId.(int)
	}

	if err := s.subjectService.CreateOne(subject); err != nil {
		s.log.Errorf("新建专题失败：error: %s", err)
		return nil, err
	}

	s.log.Infof("新建专题成功")
	return response.Success(subject), nil
}

func (s *Subject) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、获取当前用户
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("新建专题失败：未登录. ")
	}

	// 3、判断是否有权限
	if !api.CheckPermission(c, "subjects", "delete") {
		return nil, response.Forbidden.SetMsg("新建博客失败：没有权限. ")
	}

	// 4、删除该专题
	if err := s.subjectService.DeleteOne(c, id); err != nil {
		s.log.Infof("删除ID为【%d】的专题失败: %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	s.log.Infof("删除ID为【%d】的专题成功")
	return response.Success("delete success"), nil
}

func (s *Subject) Patch(c *gin.Context) (*response.Response, error) {
	s.log.Infof("patch 更新专题")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	subject := &domain.Subject{}
	//var condition map[string]interface{}
	if err := c.ShouldBindJSON(subject); err != nil {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	subject.ID = id
	if currentUserId, ok := c.Get("current_user_id"); ok {
		subject.UserID = currentUserId.(int)
	}
	if err := s.subjectService.UpdateOne(subject); err != nil {
		s.log.Errorf("更新错误： %s", err)
		return nil, err
	}

	s.log.Infof("更新成功")
	return response.Success(subject), nil
}

func (s *Subject) Put(c *gin.Context) (*response.Response, error) {
	s.log.Infof("put更新subject")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var subject domain.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		s.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := s.subjectService.SaveOne(&subject); err != nil {
		return nil, err
	}

	if v, err := s.subjectService.SelectOneById(id); err != nil {
		return nil, err
	} else {
		return response.Success(v), nil
	}
}

func (s *Subject) GetPosts(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
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
	if err := s.postService.SelectAll(&p, &post); err!=nil {
		return nil, err
	}

	s.log.Infof("根据专题查询博客成功")
	return response.Success(&p), nil
}

func (s *Subject) DeletePosts(c *gin.Context) (*response.Response, error) {
	return response.Success("success"), nil
}

func (s *Subject) PutPosts(c *gin.Context) (*response.Response, error) {
	return response.Success("success"), nil
}
