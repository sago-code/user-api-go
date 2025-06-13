package strategies

import (
	"github.com/sago-code/user-api-go/models"
	"gorm.io/gorm"
)

type DeleteStrategy interface {
	Delete(db *gorm.DB, user *models.User) error
}

type SoftDeleteStrategy struct{}

func (s *SoftDeleteStrategy) Delete(db *gorm.DB, user *models.User) error {
	return db.Delete(user).Error
}

type HardDeleteStrategy struct{}

func (s *HardDeleteStrategy) Delete(db *gorm.DB, user *models.User) error {
	return db.Unscoped().Delete(user).Error
}
