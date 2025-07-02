package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBObject struct {
	DB    *gorm.DB
	Error error
}

type UserTable struct {
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;unique"` // UUID
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BlogTable struct {
	BlogID      uuid.UUID `json:"blog_id" gorm:"type:uuid;primaryKey;unique"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Visibility  bool      `json:"visibility"`
	PublishTime time.Time `json:"publish_time"`
	UserID      uuid.UUID `json:"user_id"`
	User        UserTable `json:"user" gorm:"foreignKey:UserID;references:UserID"`
}

func InitDataBase(dbPath string) *DBObject {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return &DBObject{DB: nil, Error: err}
	}
	if err := db.AutoMigrate(&BlogTable{}, &UserTable{}); err != nil {
		return &DBObject{DB: nil, Error: err}
	}
	return &DBObject{DB: db, Error: nil}
}
