package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/model/vo"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PostService struct {
}

func (p *PostService) SelectOne(post *domain.Post) (*vo.VPosts, error) {
	db := global.DB.Model(&domain.Post{})

	if post.UserId != 0 {
		db.Where("user_id=?", post.UserId)
	}
	if post.ID != 0 {
		db.Where("id=?", post.ID).First(&post)
	}
	if db.Error != nil {
		return nil, db.Error
	}


	var tags []*domain.Tag
	var user *domain.User
	var postCoverImage *domain.File
	var userAvatar *domain.File
	var subject *domain.Subject
	var subjectAvatar *domain.File
	var subjectCoverImage *domain.File

	// 2、根据ID查询tags
	if err := global.DB.Table("tags").Joins("left join posts_tags on tags.id=posts_tags.tag_id").Where("posts_tags.post_id=?", post.ID).Find(&tags).Error; err != nil {
		return nil, err
	}

	// 3、查询用户信息
	if err := global.DB.Model(&domain.User{}).Where("id=?", post.UserId).First(&user).Error; err != nil {
		return nil, err
	}

	// 根据id查询封面图片
	if post.CoverImageId != 0 {
		if err := global.DB.Model(&domain.File{}).Where("id=?", post.CoverImageId).First(&postCoverImage).Error; err != nil {
			return nil, err
		}
	}

	// 查询用户头像
	if user != nil {
		if err := global.DB.Model(&domain.File{}).Where("id=?", user.Avatar).First(&userAvatar).Error; err != nil {
			return nil, err
		}
	}

	// 查询专题信息
	if post.SubjectId != 0 {
		if err := global.DB.Model(&domain.Subject{}).Where("id=?", post.SubjectId).First(&subject).Error; err != nil {
			return nil, err
		}
		if err := global.DB.Model(&domain.File{}).Where("id=?", subject.Avatar).First(&subjectAvatar).Error; err != nil {
			return nil, err
		}
		if err := global.DB.Model(&domain.File{}).Where("id=?", subject.CoverImage).First(&subjectCoverImage).Error; err != nil {
			return nil, err
		}
	}

	getPosts := &vo.VPosts{
		ID:              post.ID,
		Title:           post.Title,
		MarkdownContent: post.MarkdownContent,
		HtmlContent:     post.HtmlContent,
		CoverImageId:    post.CoverImageId,
		CoverImage:      postCoverImage,
		Description:     post.Description,
		Visibility:      post.Visibility,
		Status:          post.Status,
		SubjectId:       post.SubjectId,
		ImageIds:        post.ImageIds,
		Tags:            tags,
		Likes:           post.Likes,
		Views:           post.Views,
		UserId:          post.UserId,
		User: &vo.VUser{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Active:    user.Active,
			Email:     user.Email,
			Avatar:    userAvatar,
			CreatedAt: user.CreatedAt,
		},
		CreatedAt:   post.CreatedAt,
		PublishedAt: post.PublishedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	if subject != nil {
		getPosts.Subject = &vo.VSubject{
			ID:          subject.ID,
			Title:       subject.Title,
			Avatar:      subjectAvatar,
			CoverImage:  subjectCoverImage,
			Description: subject.Description,
			Visibility:  subject.Visibility,
			UserID:      subject.UserID,
			Views:       subject.Views,
			CreatedAt:   subject.CreatedAt,
		}
	}
	return getPosts, nil
}

//func (p *PostService) SelectOne(id int) (getPosts *vo.VPosts, err error) {
//	// 1、根据ID查询博客
//	var posts domain.Post
//	if err := global.DB.Model(&domain.Post{}).Where("id=?", id).First(&posts).Error; err != nil || err == gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	// 2、根据ID查询tags
//	var tags []*domain.Tag
//	if err := global.DB.Table("tags").Joins("left join posts_tags on tags.id=posts_tags.tag_id").Where("posts_tags.post_id=?", id).Find(&tags).Error; err != nil {
//		return nil, err
//	}
//
//	// 3、根据ID查询subject
//	var subject *domain.Subject
//	if posts.SubjectId != 0 {
//		if err := global.DB.Model(&domain.Subject{}).Where("id=?", posts.SubjectId).First(&subject).Error; err != nil {
//			return nil, err
//		}
//	}
//
//	// 4、根据ID查询user
//	var user *domain.User
//	if err := global.DB.Model(&domain.User{}).Where("id=?", posts.UserId).First(&user).Error; err != nil {
//		return nil, err
//	}
//
//	// 根据id查询封面图片
//	var postCoverImage *domain.File
//	if err := global.DB.Model(&domain.File{}).Where("id=?", posts.CoverImageId).First(&postCoverImage).Error; err != nil {
//		return nil, err
//	}
//
//	var userAvatar *domain.File
//	if err := global.DB.Model(&domain.File{}).Where("id=?", posts.CoverImageId).First(&userAvatar).Error; err != nil {
//		return nil, err
//	}
//
//	var subjectAvatar *domain.File
//	var subjectCoverImage *domain.File
//	if err := global.DB.Model(&domain.File{}).Where("id=?", posts.CoverImageId).First(&subjectAvatar).Error; err != nil {
//		return nil, err
//	}
//	if err := global.DB.Model(&domain.File{}).Where("id=?", posts.CoverImageId).First(&subjectCoverImage).Error; err != nil {
//		return nil, err
//	}
//
//	getPosts = &vo.VPosts{
//		ID:              posts.ID,
//		Title:           posts.Title,
//		MarkdownContent: posts.MarkdownContent,
//		HtmlContent:     posts.HtmlContent,
//		CoverImageId:    posts.CoverImageId,
//		CoverImage:      postCoverImage,
//		Description:     posts.Description,
//		Visibility:      posts.Visibility,
//		Status:          posts.Status,
//		SubjectId:       posts.SubjectId,
//		ImageIds:        posts.ImageIds,
//		Tags:            tags,
//		Likes:           posts.Likes,
//		Views:           posts.Views,
//		UserId:          posts.UserId,
//		User: &vo.VUser{
//			ID:       user.ID,
//			Username: user.Username,
//			Nickname: user.Nickname,
//			Active:   user.Active,
//			Email:    user.Email,
//			Avatar:   userAvatar,
//			Created:  user.CreatedAt,
//		},
//		CreatedAt:   posts.CreatedAt,
//		PublishedAt: posts.PublishedAt,
//		UpdatedAt:   posts.UpdatedAt,
//	}
//
//	if subject != nil {
//		getPosts.Subject = &vo.VSubject{
//			ID:          subject.ID,
//			Title:       subject.Title,
//			Avatar:      subjectAvatar,
//			CoverImage:  subjectCoverImage,
//			Description: subject.Description,
//			Visibility:  subject.Visibility,
//			UserID:      subject.UserID,
//			Views:       subject.Views,
//			CreatedAt:   subject.CreatedAt,
//		}
//	}
//
//	return getPosts, nil
//}

// 查询公开博客
func (p *PostService) SelectAllWeb(c *gin.Context, page *pager.Pager, filter *dto.ListPosts) error {
	var posts []*domain.Post
	var vPosts []*vo.VPosts
	var order domain.Dict

	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize
	var count int64

	if filter.OrderBy != 0 {
		if err := global.DB.Model(&domain.Dict{}).Where("name=? and code = ?", "order_by", filter.OrderBy).First(&order).Error; err != nil {
			return response.DatabaseSelectError.SetMsg("查询列表排序字典失败： %s", err)
		}
	} else {
		if err := global.DB.Model(&domain.Dict{}).Where("name=? and code = ?", "order_by", 1).First(&order).Error; err != nil {
			return response.DatabaseSelectError.SetMsg("查询列表排序字典失败： %s", err)
		}
	}

	db := global.DB.Debug().Model(&domain.Post{})
	if filter.Search != "" {
		db.Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", filter.Search)
	}

	if filter.SubjectId != 0 {
		db.Where("subject_id=?", filter.SubjectId)
	}

	// 判断是否登录
	isLogin, exists := c.Get("is_login")
	if exists {
		if !isLogin.(bool) {
			if filter.UserId == 0 {
				db.Where("status=2 and visibility=2")
			} else {
				db.Where("status=2 and user_id=? and visibility=2", filter.UserId)
			}
		} else {
			userId, e := c.Get("current_user_id")
			if e {
				if filter.UserId == 0 {
					db.Where("status=2 and visibility=2 or (user_id=? and visibility=?)", userId.(int), 1)
				} else {
					if filter.UserId == userId.(int) {
						db.Where("user_id=?", filter.UserId)
					} else {
						db.Where("user_id=? and visibility=2", filter.UserId)
					}
				}
			}
		}
	}

	if filter.TagId != 0 {
		db.Where("id in (?)", global.DB.Table("posts_tags").Select("post_id").Where("tag_id=?", filter.TagId))
	}

	if err := db.Count(&count).Error; err != nil {
		return err
	}
	if err := db.Order(order.Value).Offset(offset).Limit(limit).Select("ID").Find(&posts).Error; err != nil {
		return err
	}

	for _, post := range posts {
		if vp, err := p.SelectOne(post); err != nil {
			return err
		} else {
			vPosts = append(vPosts, vp)
		}
	}

	// 判断是否为空
	page.TotalRows = count
	if count == 0 {
		page.PageCount = 0
		page.List = make([]string, 0)
	} else {
		page.PageCount = int((count + int64(page.PageSize) - 1) / int64(page.PageSize))
		page.List = &vPosts
	}

	return nil
}

func (p *PostService) SelectAll(page *pager.Pager, filter *dto.ListPosts) error {
	var posts []*domain.Post
	var vPosts []*vo.VPosts
	var order domain.Dict

	if filter.OrderBy != 0 {
		if err := global.DB.Model(&domain.Dict{}).Where("name=? and code = ?", "order_by", filter.OrderBy).First(&order).Error; err != nil {
			return response.DatabaseSelectError.SetMsg("查询列表排序字典失败： %s", err)
		}
	} else {
		if err := global.DB.Model(&domain.Dict{}).Where("name=? and code = ?", "order_by", 1).First(&order).Error; err != nil {
			return response.DatabaseSelectError.SetMsg("查询列表排序字典失败： %s", err)
		}
	}

	// 1、判断是否是搜索
	db := global.DB.Model(&domain.Post{})
	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize
	var count int64

	if filter.Search != "" {
		db.Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", filter.Search).Order(order.Value)

	}

	if filter.UserId != 0 {
		db.Where("user_id=?", filter.UserId)
	}
	if filter.SubjectId != 0 {
		db.Where("subject_id=?", filter.SubjectId)
	}
	if filter.Visibility != 0 {
		db.Where("visibility=?", filter.Visibility)
	}
	if filter.Status != 0 {
		db.Where("status=?", filter.Status)
	}
	if filter.TagId != 0 {
		db.Where("id in (?)", global.DB.Table("posts_tags").Select("post_id").Where("tag_id=?", filter.TagId))
	}

	if err := db.Count(&count).Error; err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	if err := db.Order(order.Value).Offset(offset).Limit(limit).Select("ID").Find(&posts).Error; err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	for _, post := range posts {
		if vp, err := p.SelectOne(post); err != nil {
			return err
		} else {
			vPosts = append(vPosts, vp)
		}
	}

	page.PageNo = filter.PageNo
	page.PageSize = filter.PageSize
	page.TotalRows = count

	if count == 0 {
		page.PageCount = 0
		page.List = make([]string, 0)
	} else {
		page.PageCount = int((count + int64(page.PageSize) - 1) / int64(page.PageSize))
		page.List = &vPosts
	}

	return nil
}

func (p *PostService) DeleteOne(c *gin.Context, id int) error {
	var post *domain.Post

	// 1、获取user_id
	userId, _ := c.Get("current_user_id")
	db := global.DB.Model(&domain.Post{})

	// 2、根据user_id和id查询博客
	if err := db.Where("user_id=? and id=?", userId, id).First(&post).Error; err != nil {
		return err
	}

	// 2、根据id删除subject
	if err := db.Where("user_id=? and id=?", userId, id).Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

func (p *PostService) CreateOne(params *dto.AddPosts) (vPost *vo.VPosts, err error) {
	post := &domain.Post{
		Title:           params.Title,
		MarkdownContent: params.MarkdownContent,
		HtmlContent:     params.HtmlContent,
		CoverImageId:    params.CoverImageId,
		Description:     params.Description,
		Visibility:      params.Visibility,
		Status:          params.Status,
		UserId:          params.UserId,
		SubjectId:       params.SubjectId,
		ImageIds:        params.ImageIds,
		Views:           0,
		Likes:           0,
		CreatedAt:       params.CreatedAt,
		UpdatedAt:       time.Now(),
		PublishedAt:     time.Now(),
	}
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Post{}).Create(post).Error; err != nil {
			return err
		}

		for _, tag := range params.Tags {
			if err := tx.Model(&domain.PostsTags{}).Create(&domain.PostsTags{PostId: post.ID, TagId: tag.ID}); err != nil {
				tx.Callback()
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if vPosts, err := p.SelectOne(post); err != nil {
		return nil, err
	} else {
		return vPosts, nil
	}
}

func (p *PostService) UpdateOne(params *dto.PutPosts) (vPosts *vo.VPosts, err error) {
	// 查询post是否存在
	var post *domain.Post
	if err := global.DB.Model(&domain.Post{}).Where("id=? and user_id=?", params.ID, params.UserId).First(&post).Error; err != nil {
		return nil, err
	}

	post = &domain.Post{
		ID:              params.ID,
		Title:           params.Title,
		MarkdownContent: params.MarkdownContent,
		HtmlContent:     params.HtmlContent,
		CoverImageId:    params.CoverImageId,
		Description:     params.Description,
		Visibility:      params.Visibility,
		Status:          params.Status,
		UserId:          params.UserId,
		SubjectId:       params.SubjectId,
		ImageIds:        params.ImageIds,
		Likes:           params.Likes,
		Views:           params.Views,
		CreatedAt:       params.CreatedAt,
		UpdatedAt:       time.Now(),
		PublishedAt:     params.PublishedAt,
	}

	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 更新posts
		if err := tx.Model(&domain.Post{}).Where("id=?", params.ID).Updates(post).Error; err != nil {
			return err
		}
		if params.Tags != nil && len(params.Tags) > 0 {

		}
		// 删除原来的post-tag关系
		if err := tx.Model(&domain.PostsTags{}).Where("post_id=?", post.ID).Delete(&domain.PostsTags{PostId: post.ID}); err != nil {
			tx.Callback()
		}

		// 重新添加post-tag关系
		for _, tag := range params.Tags {
			if err := tx.Model(&domain.PostsTags{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&domain.PostsTags{PostId: post.ID, TagId: tag.ID}); err != nil {
				tx.Callback()
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if vPosts, err := p.SelectOne(post); err != nil {
		return nil, err
	} else {
		return vPosts, nil
	}
}

func (p *PostService) IncrementViews(id int) error {
	return global.DB.Model(&domain.Post{}).Omit("updated_at").Where("id=?", id).Update("views", gorm.Expr("views + 1")).Error
}

func (p *PostService) DecrementViews(id int) error {
	return global.DB.Model(&domain.Post{}).Omit("updated_at").Where("id=?", id).Update("views", gorm.Expr("views - 1")).Error
}

func (p *PostService) IncrementLikes(id int) error {
	return global.DB.Model(&domain.Post{}).Select("likes").Omit("updated_at").Where("id=?", id).Update("likes", gorm.Expr("likes + 1")).Error
}

func (p *PostService) DecrementLikes(id int) error {
	return global.DB.Model(&domain.Post{}).Select("likes").Omit("updated_at").Where("id=?", id).Update("likes", gorm.Expr("likes - 1")).Error
}

func NewPostService() *PostService {
	return &PostService{}
}
