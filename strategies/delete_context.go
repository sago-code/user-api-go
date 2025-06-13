package strategies

import (
	"github.com/sago-code/user-api-go/models"
	"gorm.io/gorm"
)

type DeleteContext struct {
	db       *gorm.DB
	strategy DeleteStrategy
}

func NewDeleteContext(db *gorm.DB, strategy DeleteStrategy) *DeleteContext {
	return &DeleteContext{
		db:       db,
		strategy: strategy,
	}
}

func (c *DeleteContext) SetStrategy(strategy DeleteStrategy) {
	c.strategy = strategy
}

func (c *DeleteContext) DeleteUser(user *models.User) error {
	return c.strategy.Delete(c.db, user)
}
