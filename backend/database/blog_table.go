package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (b *BlogTable) BeforeCreate(tx *gorm.DB) (Error error) {
	b.BlogID = uuid.Must(uuid.NewRandom())
	return nil
}
func (u *BlogTable) AfterUpdate(tx *gorm.DB) (Error error) {
	u.UpdatedAt = time.Now()
	return nil
}


func (db *DBObject) CreateBlog(user_id uuid.UUID,title string, summary string) error{

	return nil
}


func (db *DBObject) GetBlogs(user_id uuid.UUID, page int, count int) ([]BlogTable, error) {
	return []BlogTable{}, nil
}

func (db *DBObject) GetBlogByID(user_id uuid.UUID, blog_id uuid.UUID) (BlogTable, error) {
	return BlogTable{}, nil
}
