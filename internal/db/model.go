package db

import "errors"

type Product struct {
	ID         int64
	Title      string
	Desription string
}

type ProductItem struct {
	ID        int64
	Sku       string
	Material  string
	ProductID int64
}

var ErrProductNotFound = errors.New("product not found")

type ErrNotFound struct {
	messange string
}

func (e *ErrNotFound) Error() string {
	return e.messange
}
