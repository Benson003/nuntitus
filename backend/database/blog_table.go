package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




func (b *BlogTable) BeforeCreate(tx *gorm.DB)(Error error){
	b.BlogID = uuid.Must(uuid.NewRandom())
	return nil
}
func (u *UserTable) AfterUpdate(tx *gorm.DB) (Error error) {
	u.UpdatedAt = time.Now()
	return nil
}

