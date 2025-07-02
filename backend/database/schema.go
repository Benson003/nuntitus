package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
This file defines the schema for my data base
Users Table and Blog Table
The init method spins up a sqlite db but
plans in later versions to add support for postgres
*/

//This is a ponter to the db object

type DBObject struct {
	DB    *gorm.DB
	Error error
}

// This is the user table for logins
// I plan to later intrgrate oauth
// but im stcking with jwt for now
type UserTable struct {
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;unique"`
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// The blogs table
// No real explaning to do here
// May later add metrics as a seprate table
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
	FilePath	string
}

// This inits the data base
// I swear i didnt know where else to put this
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
