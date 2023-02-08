package db

type Product struct {
	ID    int64
	Title string
	Desription string
}

type ProductItem struct {
	ID            int64
	Sku           string
	Material      string
	IdProductItem int64
}
