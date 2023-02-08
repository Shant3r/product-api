package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shant3r/product-api/internal/db"
	"github.com/satori/go.uuid"
)

type HandlerItem struct {
	repo *db.RepositoryItem
}

func NewItems(repository *db.RepositoryItem) *HandlerItem {
	return &HandlerItem{repo: repository}
}

func (h *HandlerItem) AddProductItem(ctx context.Context, c *gin.Context) {
	reqItem := new(AddProductItemRequest)
	err := c.BindJSON(reqItem)
	if err != nil {
		internalError(c, err)
		return
	}

	if reqItem.Material == "" {
		badRequest(c)
		return
	}

	reqSku:= uuid.NewV4()

	err = h.repo.AddProductItem(ctx, reqItem.Material, reqSku)
	if err != nil {
		internalError(c, err)
		return
	}
}
