package resolver

import (
	"gorm.io/gorm"
	"remote-schema/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []*model.Product
	DB       *gorm.DB
}
