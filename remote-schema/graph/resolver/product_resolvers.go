package resolver

import (
	"context"
	"remote-schema/graph/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	p := &model.Product{
		ID:   len(r.products) + 1,
		Name: input.Name,
	}
	r.products = append(r.products, p)

	return p, nil
}

// MaskingProducts is the resolver for the masking_products field.
func (r *queryResolver) MaskingProducts(ctx context.Context) ([]*model.Product, error) {
	products := r.products
	for _, product := range products {
		product.Name = "xxxxxxxxx"
	}

	return products, nil
}
