package web



//type Tag struct {
//	logs         logs.Logger
//	tagService  *service.TagService
//	postService service.PostService
//}
//
//func NewTag() *Tag {
//	return &Tag{
//		logs:         global.Logger,
//		tagService:  service.NewTagService(),
//		postService: service.NewPostService(),
//	}
//}
//
//func (t *Tag) Get(c *gin.Context) (*response.Response, error) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, response.InvalidParams.SetMsg("ID is required. ")
//	}
//
//	tag := &domain.Tag{ID: id}
//	if err := t.tagService.SelectOne(tag); err != nil {
//		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的标签失败 : %s", id, err)
//	}
//	return response.Success(tag), nil
//}
//
//func (t *Tag) List(c *gin.Context) (*response.Response, error) {
//	p := pager.Pager{
//		PageNo:   request.GetPageNo(c),
//		PageSize: request.GetPageSize(c),
//	}
//
//	if err := t.tagService.SelectAll(c, &p); err != nil {
//		t.logs.Errorf("分页查询标签失败： %s", err)
//		return nil, response.InternalServerError.SetMsg("%s", err)
//	}
//
//	t.logs.Infof("分页查询标签成功：%s", &p)
//	return response.Success(&p), nil
//}