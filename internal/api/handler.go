package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/shant3r/product-api/internal/db"
)

type Handler struct {
	repo *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{repo: repository}
}

func (h *Handler) AddProduct(ctx context.Context, c *gin.Context) {
	req := new(AddProductRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}
	if req.Title == "" {
		badRequest(c)
		return
	}
	if req.Desription == "" {
		badRequest(c)
		return
	}
	err = h.repo.AddProduct(ctx, req.Title, req.Desription)
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) AddProductItem(ctx context.Context, c *gin.Context) {
	req := new(AddProductItemRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}

	if req.Material == "" {
		badRequest(c)
		return
	}
	if req.ProductID <= 0 {
		badRequest(c)
		return
	}

	reqSku := uuid.NewV4()

	err = h.repo.AddProductItem(ctx, reqSku.String(), req.Material, req.ProductID)
	if err != nil {
		_,ok := err.(*db.ErrNotFound)
		if ok {
			badRequest(c)
			return
		}
		
		if err == db.ErrProductNotFound {
			badRequest(c)
			return
		}
		internalError(c, err)
		return
	}
}

// func (h *Handler) GetProducts(ctx context.Context, c *gin.Context) {
// 	idString := c.Request.URL.Query().Get("id")
// 	if idString != "" {
// 		id, err := strconv.ParseInt(idString, 10, 64)
// 		if err != nil {
// 			badRequst(c)
// 			return
// 		}
// 		product, ok := h.getProduct(id)
// 		if ok {
// 			statusOk(c, product)
// 		} else {
// 			notFound(c)
// 		}
// 		return
// 	}
// 	products, err := h.r.GetProducts(ctx)
// 	if err != nil {
// 		return
// 	}

// 	c.JSON(http.StatusOK, convertToProducts(products))

// }

// func (h *Handler) getProduct(id int64) (*Product, bool) {
// 	product, ok := h.r.GetProduct(id)
// 	if ok {
// 		return convertToProduct(product), true
// 	}
// 	return nil, false
// }

// func convertToProduct(p *db.Product) *Product {
// 	return &Product{
// 		Identity: p.ID,
// 		Name:     p.Title,
// 			}
// }

// func convertToDBProduct(p *Product) *db.Product {
// 	return &db.Product{
// 		Title: p.Name,
// 		ID:    p.Identity,
// 		}
// }

// func convertToProducts(products []*db.Product) []*Product {
// 	res := make([]*Product, 0, len(products))
// 	for _, p := range products {
// 		res = append(res, convertToProduct(p))
// 	}
// 	return res

// }

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}

// func notFound(c *gin.Context) {
// 	c.JSON(http.StatusNotFound, "not found")
// }

// func statusOk(c *gin.Context, val any) {
// 	c.JSON(http.StatusOK, val)
// }
