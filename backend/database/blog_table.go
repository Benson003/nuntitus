package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (b *BlogTable) BeforeCreate(tx *gorm.DB) (err error) {
	b.BlogID = uuid.Must(uuid.NewRandom())
	return nil
}
func (u *BlogTable) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}

func (db *DBObject) CreateBlog(user_id uuid.UUID, title string, summary string, publish_time time.Time, visiblity bool) (*BlogTable, error) {
	blog := NewBlogBuilder()
	blog.SetTitle(title).
		SetSummary(summary).
		SetPublishTime(publish_time).
		SetVisiblity(visiblity).
		SetUserID(user_id)


	f_blog := blog.Build()
	if err := db.DB.Create(&f_blog).Error; err != nil {
		return nil, err
	}
	return &f_blog, nil
}

func (db *DBObject) GetBlogByID(user_id uuid.UUID, blog_id uuid.UUID) (*BlogTable, error) {
	var blog BlogTable
	if err := db.DB.Preload("User").Where("user_id = ? AND blog_id = ?", user_id, blog_id).First(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

func (db *DBObject) GetBlogsPaginated(user_id uuid.UUID, page int, count int) ([]*BlogTable, error) {
	var blogs []*BlogTable

	offset := (page - 1) * count
	if offset < 0 {
		offset = 0
	}

	err := db.DB.
		Preload("User").
		Where("user_id = ?", user_id).
		Limit(count).
		Offset(offset).
		Order("created_at DESC").
		Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (db *DBObject) UpdateBlog(user_id uuid.UUID, blog_id uuid.UUID, title string, summary string, visiblity bool, publish_time time.Time) (*BlogTable, error) {
	blogBuilder := NewBlogBuilder()

	blog, err := blogBuilder.FetchBlogToUpdate(db.DB, user_id, blog_id)
	if err != nil {
		return nil, err
	}
	if title != "" {
		blog.SetTitle(title)
	}
	if summary != "" {
		blog.SetSummary(summary)
	}
	blog.SetVisiblity(visiblity)
	blog.SetPublishTime(publish_time)
	f_blog := blog.Build()

	if err := db.DB.Save(&f_blog).Error; err != nil {
		return nil, err
	}

	return &f_blog, nil
}

func (db *DBObject) DeleteBlog(user_id uuid.UUID, blog_id uuid.UUID) (*BlogTable, error) {
	blog, err := db.GetBlogByID(user_id, blog_id)
	if err != nil {
		return nil, err
	}
	if err := db.DB.Delete(blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}
