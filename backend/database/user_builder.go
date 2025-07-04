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

// Returns the UserBuilder
func NewUserBuilder() UserBuilder {
	return UserBuilder{
		User: UserTable{},
	}
}

// Fetches a user from the DB
// Use only for updates

func (b *UserBuilder) FetchUserToUpdate(db *gorm.DB, user_id uuid.UUID) (*UserBuilder, error) {
	var user UserTable
	if err := db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		return &UserBuilder{}, err
	}
	return &UserBuilder{User: user}, nil
}

// Set username who would have thunk it
func (b *UserBuilder) SetUsername(username string) *UserBuilder {
	b.User.Username = username
	return b
}

/*
Please pass in a hash
I do not have the mental capacity to keep it synced every where
 without double or triple hashing
 */
func (b *UserBuilder) SetPasswordHash(password_hash string) *UserBuilder {
	b.User.PasswordHash = password_hash
	return b
}

//Set the email prety standard stuff
func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.User.Email = email
	return b
}

//Set the first name
func (b *UserBuilder) SetFirstName(first_name string) *UserBuilder {
	b.User.FirstName = first_name
	return b

}
// Set the last name
func (b *UserBuilder) SetLastName(last_name string) *UserBuilder {
	b.User.LastName = last_name
	return b

}

//Builds and returns the user object for creating or pdating
func (b *UserBuilder) Build() UserTable {
	return b.User
}
