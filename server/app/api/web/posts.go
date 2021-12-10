package web

import (
	"blog/app/api"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if one, err := p.postService.SelectOne(id); err != nil {
		p.log.Errorf("查询ID为【%d】的博客失败： %s", id, err)
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的博客失败： %s", id, err)
	} else {
		if one.Status == 1  {
			p.log.Errorf("查询ID为【%d】的博客为草稿", id)
			return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
		}
		if one.Visibility == 1 {
			if !api.CheckLogin(c) {
				p.log.Errorf("查询ID为【%d】的博客为私有，需要登录. ", id)
				return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
			} else {
				// 登录的话，判断该用户id和博客的id是否一致，不一致的话，返回not found
				userId, _ := c.Get("current_user_id")
				if one.UserId != userId {
					p.log.Errorf("查询ID为【%d】的博客为私有，登录用户不一致，没有该博客查询权限 ", id)
					return nil, response.RecordNotFound.SetMsg("没有该博客权限. ")
				}
			}
		}

		go func() {
			// 更新博客views
			params := &dto.PutPosts{
				ID:    id,
				Views: one.Views + 1,
			}
			_, _ = p.postService.UpdateOne(params)
		}()

		p.log.Infof("查询博客【%d】成功. ", id)
		return response.Success(one), nil
	}
}

func (p *Post) List(c *gin.Context) (*response.Response, error) {
	// 1、获取参数
	post := dto.ListPosts{}
	if err := c.ShouldBind(&post); err != nil {
		p.log.Errorf("绑定参数错误: error: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、查询博客
	page := pager.Pager{}
	if err := p.postService.SelectAll1(c, &page, &post); err != nil {
		p.log.Errorf("分页查询博客失败: %s", err)
		return nil, response.InternalServerError.SetMsg("分页查询博客失败：%s", err)
	}

	// 3、返回查询结果
	p.log.Infof("分页查询博客成功: [第 %d 页，总页数：%d, 总行数：%d]", page.PageNo, page.PageCount, page.TotalRows)
	return response.PagerResponse(&page), nil
}

func (p *Post) Like(c *gin.Context) (*response.Response, error){
	// 更新post的like
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if one, err := p.postService.SelectOne(id); err != nil {
		return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
	} else {
		if one.Status == 1  {
			return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
		}
		go func() {
			// 更新博客views
			params := &dto.PutPosts{
				ID:    id,
				Likes: one.Likes + 1,
			}
			_, _ = p.postService.UpdateOne(params)
		}()
	}

	return response.Success("success"),nil
}