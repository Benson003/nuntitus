package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBuilder struct {
	User UserTable
}
/*
This is the User Builder
It is used to make the user creation and updataing logic simpler
you intialise a new UserBuilder
user := NewUserBuilder()
then set methods like
user.SetPasswordHash("...")
then the build method
user.Build() which returns a UserTable Struct
*/



func NewUserBuilder() UserBuilder {
	return UserBuilder{
		User: UserTable{},
	}
}
func (b *UserBuilder) FetchUserToUpdate(db *gorm.DB, user_id uuid.UUID) (*UserBuilder,error) {
	var user UserTable
	if err := db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		return &UserBuilder{}, err
	}
	return &UserBuilder{User: user}, nil
}

func (b *UserBuilder)SetUsername(username string)(*UserBuilder){
	b.User.Username = username
	return b
}

func (b *UserBuilder)SetPasswordHash(password_hash string)(*UserBuilder){
	b.User.PasswordHash  = password_hash
	return b
}
func (b *UserBuilder)SetEmail(email string)(*UserBuilder){
	b.User.Email = email
	return b
}

func (b *UserBuilder)Build()(UserTable){
	return b.User
}
