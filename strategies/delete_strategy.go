package strategies

import (
	"github.com/sago-code/user-api-go/models"
	"gorm.io/gorm"
)

// define la interfaz para las estrategias de eliminación
type DeleteStrategy interface {
	Delete(db *gorm.DB, user *models.User) error
}

// implementa la estrategia de borrado lógico
type SoftDeleteStrategy struct{}

// realiza un borrado lógico del usuario
func (s *SoftDeleteStrategy) Delete(db *gorm.DB, user *models.User) error {
	return db.Delete(user).Error
}

// implementa la estrategia de borrado físico o permanente
type HardDeleteStrategy struct{}

// realiza un borrado físico o permanente del usuario
func (s *HardDeleteStrategy) Delete(db *gorm.DB, user *models.User) error {
	return db.Unscoped().Delete(user).Error
}
