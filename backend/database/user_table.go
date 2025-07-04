package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This ensures the created at always has a date
func (u *UserTable) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.Must(uuid.NewRandom())
	return nil
}
func (u *UserTable) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Where("user_id = ?", u.UserID).Delete(&BlogTable{}).Error
}

// If any record Updates
func (u *UserTable) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}

func (db *DBObject) CreateUser(username string, hashed_password string, email string) (*UserTable, error) {
	user := NewUserBuilder()
	user.SetUsername(username).
		SetEmail(email).
		SetPasswordHash(hashed_password)

	f_user := user.Build()

	if err := db.DB.Create(&f_user).Error; err != nil {
		return nil, err
	}
	return &f_user, nil
}

func (db *DBObject) GetUserByID(user_id uuid.UUID) (*UserTable, error) {
	var user UserTable
	if err := db.DB.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DBObject) GetUserByEmail(email string) (*UserTable, error) {
	var user UserTable
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (db *DBObject) GetUserByUsername(username string) (*UserTable, error) {
	var user UserTable
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (db *DBObject) GetUserByUsernameOrEmail(username string, email string) (*UserTable, error) {
	var user UserTable
	if err := db.DB.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DBObject) DeleteUser(user_id uuid.UUID) (*UserTable, error) {
	user, err := db.GetUserByID(user_id)
	if err != nil {
		return nil, err
	}
	if err := db.DB.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DBObject) UpdateUser(user_id uuid.UUID, username string, password_hash string, email string, first_name string, last_name string) (*UserTable, error) {
	userBuilder := NewUserBuilder()

	user, err := userBuilder.FetchUserToUpdate(db.DB, user_id)
	if err != nil {
		return nil, err
	}
	if username != "" {
		user.SetUsername(username)
	}
	if first_name != "" {
		user.SetFirstName(first_name)
	}
	if last_name != "" {
		user.SetLastName(last_name)
	}

	if password_hash != "" {
		user.SetPasswordHash(password_hash)
	}
	if email != "" {
		user.SetEmail(email)
	}
	f_user := user.Build()
	if err := db.DB.Save(&f_user).Error; err != nil {
		return nil, err
	}

	return &f_user, nil
}
