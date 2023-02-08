package db

import (
	"context"
	"database/sql"
	"errors"
)


type Repository struct {
	products []*Product
	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
		products: []*Product{},
		database: database,
	}
}


func (r *Repository) AddProduct(ctx context.Context, title string, description string) error {
	if title == "" {
		return errors.New("title is empty")
	}
	if description == "" {
		return errors.New("description is empty")
	}
	
	_, err := r.database.ExecContext(ctx, `
		insert into product (title, description)
		values ($1, $2)
	`, title, description)
	if err != nil {
		return err
	}
	return nil
}

// func (r *Repository) GetProducts(ctx context.Context) ([]*Product, error) {
// 	raws, err := r.database.QueryContext(ctx, `
// 		select id, title
// 		from product
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer raws.Close()

// 	var result []*Product
// 	for raws.Next() {
// 		p := new(Product)
// 		err = raws.Scan(&p.ID, &p.Title)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, p)
// 	}

// 	return result, nil
// }

// func (r *Repository) GetProduct(id int64) (*Product, bool) {
// 	for _, product := range r.products {
// 		if id == product.ID {
// 			return product, true
// 		}

// 	}
// 	return nil, false
// }
