package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrJsonUnmarshal     = errors.New("failed to parse product from request body")
	ErrProductIdMismatch = errors.New("product ID in path does not match product ID in body")
)

type Products struct {
	store Store
}

func NewProductsDomain(s Store) *Products {
	return &Products{
		store: s,
	}
}

func (d *Products) GetProduct(ctx context.Context, id string) (*Product, error) {
	product, err := d.store.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return product, nil
}

func (d *Products) AllProducts(ctx context.Context, next *string) (ProductRange, error) {
	if next != nil && strings.TrimSpace(*next) == "" {
		next = nil
	}

	productRange, err := d.store.All(ctx, next)
	if err != nil {
		return productRange, fmt.Errorf("%w", err)
	}

	return productRange, nil
}

func (d *Products) PutProduct(ctx context.Context, id string, body []byte) (*Product, error) {
	product := Product{Id: id}
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	productModel := ProductModel{
		Pk:    "product",
		Sk:    "product_id::" + product.Id,
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
	err := d.store.Put(ctx, productModel)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &product, nil
}

func (d *Products) DeleteProduct(ctx context.Context, id string) error {
	err := d.store.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
