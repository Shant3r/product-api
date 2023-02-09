package api

type AddProductRequest struct {
	Title      string `json:"title"`
	Desription string `json:"description"`
}

type AddProductItemRequest struct {
	Material  string `json:"material"`
	ProductID int64  `json: "productID"`
}
