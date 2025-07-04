package database

import (
	"time"

	"github.com/Benson003/nuntius/tools/files"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogBuilder struct {
	Blog BlogTable
}

func NewBlogBuilder() BlogBuilder {
	return BlogBuilder{
		Blog: BlogTable{},
	}
}

func (b *BlogBuilder) FetchBlogToUpdate(db *gorm.DB, user_id uuid.UUID, blog_id uuid.UUID) (*BlogBuilder, error) {
	var blog BlogTable
	if err := db.Where("user_id = ? AND blog_id = ?", user_id, blog_id).First(&blog).Error; err != nil {
		return &BlogBuilder{}, err
	}
	return &BlogBuilder{Blog: blog}, nil
}

func (b *BlogBuilder) SetTitle(title string) *BlogBuilder {
	b.Blog.Title = title
	return b
}
func (b *BlogBuilder) SetSummary(summary string) *BlogBuilder {
	b.Blog.Summary = summary
	return b
}

func (b *BlogBuilder) SetVisiblity(visibility bool) *BlogBuilder {
	b.Blog.Visibility = visibility
	return b
}
func (b *BlogBuilder) SetPublishTime(publish_time time.Time) *BlogBuilder {
	b.Blog.PublishTime = publish_time
	return b
}
func (b *BlogBuilder) SetUserID(user_id uuid.UUID) *BlogBuilder {
	b.Blog.UserID = user_id
	return b
}

func (b *BlogBuilder) SetContent(content string) *BlogBuilder {
	files.UploadFile(b.Blog.UserID, []byte(content))
	return b
}

func (b *BlogBuilder) Build() BlogTable {
	return b.Blog
}
