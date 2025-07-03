package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
This file defines the schema for the database:
- Users table
- Blogs table

Currently uses SQLite for simplicity,
with plans to add Postgres support later.
*/

// DBObject wraps the GORM DB pointer and any initialization error.
// This makes it easy to pass around and check DB health.
type DBObject struct {
	DB    *gorm.DB
	Error error
}

// UserTable represents registered users in the system.
// Sticking with JWT for now, may add OAuth in the future.
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

// BlogTable represents individual blogs created by users.
// May later add metrics as a separate table if needed.
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

// InitDataBase inits the SQLite database and auto-migrates tables.
// Use this on app startup, or you won't get a usable DB (trust me).
// It returns a DBObject that wraps the GORM DB and the error (if any).
// Yeah, I return errors this way. Feels natural, don’t ask.
// TODO: Write a proper way of making this integrate with multiple DataBases(including wrappers for clod providers if possible)
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
