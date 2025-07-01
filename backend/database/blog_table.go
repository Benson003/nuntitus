package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)






func (b *BlogTable) BeforeCreate(tx *gorm.DB)(Error error){
	b.BlogID = uuid.Must(uuid.NewRandom())
	return nil
}
