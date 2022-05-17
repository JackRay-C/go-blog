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

type Head struct {
	log logger.Logger
	headService service.HeadService

}

func NewHead() *Head  {
	return &Head{
		log:         global.Log,
		headService: service.NewHeadService(),
	}
}

func (h *Head) Get(c *gin.Context) (*vo.Response, error)  {
	h.log.Info("根据ID获取head信息")

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams
	}

	head := &po.Head{ID:ID}

	if err := h.headService.ISelectOneWeb(c, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	if head.Visibility == global.Public {
		return vo.Success(head), nil
	} else {
		if UserID, ok := c.Get(global.SessionUserIDKey);ok {
			if head.UserID != UserID.(int) {
				return nil, vo.Forbidden.SetMsg("该博客为私有博客，无法查看！")
			}
			return vo.Success(head), nil
		}
		return nil, vo.Forbidden.SetMsg("没有权限，请先登录!")
	}
}

func (h *Head) List(c *gin.Context) (*vo.Response, error)  {
	head := &po.Head{
		Visibility: global.Public,
		Status:     global.Staged,
	}

	if userID, ok := c.Get(global.SessionUserIDKey); ok {
		head.UserID = userID.(int)
		head.Visibility = 0
	}

	pager := &vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if err := h.headService.ISelectListWeb(c, pager, head); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(pager), nil
}