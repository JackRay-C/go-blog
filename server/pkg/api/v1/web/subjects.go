package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Subject struct {
	log            logger.Logger
	subjectService service.SubjectService
	postService service.PostService
}

func NewSubject() *Subject {
	return &Subject{
		log:            global.Log,
		subjectService: service.NewSubjectService(),
		postService: service.NewPostService(),
	}
}

func (s Subject) List(c *gin.Context) (*vo.Response, error) {
	params := po.Subject{}
	if err := c.ShouldBind(&params); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}
	p.MustSort(c)

	if err := s.subjectService.ISelectListWeb(c, &p, &params); err != nil {
		s.log.Errorf("分页查询失败： %s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	s.log.Infof("分页查询成功：%s", &p)
	return vo.Success(&p), nil
}

func (s Subject) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 查询subject
	subject := &po.Subject{ID: id}
	if err := s.subjectService.ISelectOneWeb(c, subject); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	} else {
		if subject.Visibility == 1 {
			// 判断是否登录
			get, exists := c.Get("is_login")
			if exists {
				if !get.(bool) {
					return nil, vo.Forbidden.SetMsg("该专题为私有，无权限查看! ")
				} else {
					// 如果登录，判断用户id是否为登录用户
					userId, e := c.Get("current_user_id")
					if e {
						if userId == subject.UserID {
							// 更新subject views + 1

							return vo.Success(subject), nil
						} else {
							return nil, vo.Forbidden.SetMsg("该专题为私有，无权限查看! ")
						}
					}
				}
			}
		}
		// 更新subject views + 1

		return vo.Success(subject), nil
	}
}