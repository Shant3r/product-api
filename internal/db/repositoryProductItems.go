package db

import (
	"context"
	"database/sql"
	"errors"
)


type RepositoryItem struct {
	productsItems []*ProductItem
	database *sql.DB
}

func NewItems(database *sql.DB) *RepositoryItem {
	return &RepositoryItem{
		productsItems: []*ProductItem{},
		database: database,
	}
}

func (r *RepositoryItem) AddProductItem(ctx context.Context, sku string, material string, IdProductItem int64) error {
	if sku == "" {
		return errors.New("sku is empty")
	}
	if material == "" {
		return errors.New("material is empty")
	}
	_, err := r.database.ExecContext(ctx, `
	insert into product (sku, material,IdProductItem)
	values ($1, $2)
`, sku, material, IdProductItem)
if err != nil {
	return err
}

	return nil
}