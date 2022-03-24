package web

//type Post struct {
//	logs         logs.Logger
//	postService service.PostService
//	commentService service.CommentService
//}
//
//func NewPost() *Post {
//	return &Post{
//		logs:         global.Logger,
//		postService: service.NewPostService(),
//		commentService: service.NewCommentService(),
//	}
//}
//
//func (p *Post) Get(c *gin.Context) (*response.Response, error) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	if err := p.postService.ISelectOneWeb(c, &po.Post{ID: id}); err != nil {
//		p.logs.Errorf("查询ID为【%d】的博客失败： %s", id, err)
//		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的博客失败： %s", id, err)
//	}
//
//	return response.Success("success"), nil
//	//else {
//	//	if one.Status == 1  {
//	//		p.logs.Errorf("查询ID为【%d】的博客为草稿", id)
//	//		return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
//	//	}
//	//	if one.Visibility == 1 {
//	//		if !api.CheckLogin(c) {
//	//			p.logs.Errorf("查询ID为【%d】的博客为私有，需要登录. ", id)
//	//			return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
//	//		} else {
//	//			// 登录的话，判断该用户id和博客的id是否一致，不一致的话，返回not found
//	//			userId, _ := c.Get("current_user_id")
//	//			if one.UserId != userId {
//	//				p.logs.Errorf("查询ID为【%d】的博客为私有，登录用户不一致，没有该博客查询权限 ", id)
//	//				return nil, response.RecordNotFound.SetMsg("没有该博客权限. ")
//	//			}
//	//		}
//	//	}
//	//
//	//	go func() {
//	//		// 更新博客views
//	//		_ = p.postService.IncrementViews(id)
//	//	}()
//	//
//	//	p.logs.Infof("查询博客【%d】成功. ", id)
//	//	return response.Success(one), nil
//	//}
//}
//
//func (p *Post) List(c *gin.Context) (*response.Response, error) {
//	// 1、获取参数
//	post := po.Post{}
//	if err := c.ShouldBind(&post); err != nil {
//		p.logs.Errorf("绑定参数错误: error: %s", err)
//		return nil, response.InvalidParams.SetMsg("%s", err)
//	}
//
//	// 2、查询博客
//	page := pager.Pager{
//		PageNo:   request.GetPageNo(c),
//		PageSize: request.GetPageSize(c),
//	}
//	if err := p.postService.ISelectAllWeb(c, &page, &post); err != nil {
//		p.logs.Errorf("分页查询博客失败: %s", err)
//		return nil, response.InternalServerError.SetMsg("分页查询博客失败：%s", err)
//	}
//
//	// 3、返回查询结果
//	p.logs.Infof("分页查询博客成功: [第 %d 页，总页数：%d, 总行数：%d]", page.PageNo, page.PageCount, page.TotalRows)
//	return response.Success(&page), nil
//}
//
//// todo: 增加用户like表，管理用户喜欢的文章
//func (p *Post) Like(c *gin.Context) (*response.Response, error){
//	// 更新post的like
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	if err := p.postService.ISelectOne(c, &po.Post{ID: id}); err != nil {
//		return nil, response.RecordNotFound.SetMsg("该博客不存在. ")
//	}
//
//	return response.Success("success"),nil
//}