package web

import (
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
	postService *service.PostService
}

func NewSubject() *Subject {
	return &Subject{
		log:            global.Logger,
		subjectService: service.NewSubjectService(),
		postService: service.NewPostService(),
	}
}

func (s Subject) List(c *gin.Context) (*response.Response, error) {
	params := dto.ListSubjects{}
	if err := c.ShouldBind(&params); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if err := s.subjectService.SelectAllWeb(c, &p, &params); err != nil {
		s.log.Errorf("分页查询失败： %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	s.log.Infof("分页查询成功：%s", &p)
	return response.Success(&p), nil
}

func (s Subject) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 查询subject
	if vSubject, err := s.subjectService.SelectOneById(id); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	} else {
		if vSubject.Visibility == 1 {
			// 判断是否登录
			get, exists := c.Get("is_login")
			if exists {
				if !get.(bool) {
					return nil, response.Forbidden.SetMsg("该专题为私有，无权限查看! ")
				} else {
					// 如果登录，判断用户id是否为登录用户
					userId, e := c.Get("current_user_id")
					if e {
						if userId == vSubject.UserID {
							// 更新subject views + 1
							go func() {
								_ = s.subjectService.IncrementViews(id)
							}()
							return response.Success(vSubject), nil
						} else {
							return nil, response.Forbidden.SetMsg("该专题为私有，无权限查看! ")
						}
					}
				}
			}
		}
		// 更新subject views + 1
		go func() {
			_ = s.subjectService.IncrementViews(id)
		}()
		return response.Success(vSubject), nil
	}
}

func (s Subject) GetPosts(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	params := dto.ListPosts{
		PageNo:    request.GetPageNo(c),
		PageSize:  request.GetPageSize(c),
		SubjectId: id,
		OrderBy:   0,
		TagId:     0,
		Search:    "",
	}
	p := pager.Pager{}

	if err := s.postService.SelectAllWeb(c, &p, &params); err != nil {
		return nil, err
	}

	return response.Success(&p), nil
}
