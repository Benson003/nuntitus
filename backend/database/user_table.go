package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This ensures the created at always has a date
func (u *UserTable) BeforeCreate(tx *gorm.DB) (Error error) {
	u.UserID = uuid.Must(uuid.NewRandom())
	return nil
}

// If any record
func (u *UserTable) AfterUpdate(tx *gorm.DB) (Error error) {
	u.UpdatedAt = time.Now()
	return nil
}

func (db *DBObject) CreateUser(username string, hashed_password string, email string) (UserTable, error) {
	user := NewUserBuilder()
	user.SetUsername(username).SetEmail(email).SetPasswordHash(hashed_password)
	f_user := user.Build()

	if err := db.DB.Create(&f_user).Error; err != nil {
		return f_user, err
	}
	return f_user, nil
}

func (db *DBObject) GetUserByID(user_id uuid.UUID) (UserTable, error) {
	var user UserTable
	if err := db.DB.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *DBObject) GetUserByEmail(email string) (UserTable, error) {
	var user UserTable
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (db *DBObject) GetUserByUsername(username string) (UserTable, error) {
	var user UserTable
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *DBObject) DeleteUser(user_id uuid.UUID) (UserTable, error) {
	user, err := db.GetUserByID(user_id)
	if err != nil {
		return user, err
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *DBObject) UpdateUser(user_id uuid.UUID, username string, password_hash string, email string) (UserTable, error) {
	userBuilder := NewUserBuilder()

	user, err := userBuilder.FetchUserToUpdate(db.DB, user_id)
	if err != nil {
		return UserTable{}, err
	}
	if username != "" {
		user.SetUsername(username)
	}

	if password_hash != "" {
		user.SetPasswordHash(password_hash)
	}
	if email != "" {
		user.SetEmail(email)
	}
	f_user := user.Build()
	if err := db.DB.Save(&f_user).Error; err != nil {
		return UserTable{}, err
	}

	return f_user, nil
}
