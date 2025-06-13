package strategies

import (
	"github.com/sago-code/user-api-go/models"
	"gorm.io/gorm"
)

// contexto que utiliza una estrategia de eliminación
type DeleteContext struct {
	db       *gorm.DB
	strategy DeleteStrategy
}

// crea un nuevo contexto con la estrategia especificada
func NewDeleteContext(db *gorm.DB, strategy DeleteStrategy) *DeleteContext {
	return &DeleteContext{
		db:       db,
		strategy: strategy,
	}
}

// establece una nueva estrategia de eliminación
func (c *DeleteContext) SetStrategy(strategy DeleteStrategy) {
	c.strategy = strategy
}

// elimina un usuario utilizando la estrategia actual
func (c *DeleteContext) DeleteUser(user *models.User) error {
	return c.strategy.Delete(c.db, user)
}
