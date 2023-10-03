package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/base-media-cloud/base-test/internal/domain/models/product"
	"github.com/base-media-cloud/base-test/internal/domain/ports"
	"github.com/base-media-cloud/base-test/utils"
)

type Handler struct {
	svcs   ports.Services
	prods  product.Products
	logger zerolog.Logger
}

func New(svcs ports.Services, prods product.Products, l zerolog.Logger) *Handler {
	return &Handler{
		svcs:   svcs,
		prods:  prods,
		logger: l,
	}
}

func (h *Handler) Get(c *gin.Context) {
	res := make([]product.ResProd, len(h.prods))
	for i, p := range h.prods {
		delPrice := h.svcs.DeliveryPrice(p.Weight)
		res[i] = product.ResProd{
			Name:          p.Name,
			ProductPrice:  utils.ToDecimalString(p.Price),
			DeliveryPrice: utils.ToDecimalString(delPrice),
			TotalPrice:    utils.ToDecimalString(p.Price + delPrice),
		}
	}

	delType := h.svcs.DeliveryType()
	h.logger.Info().
		Str("provider", delType).
		Msg("successfully got the prices of the products")

	c.AbortWithStatusJSON(http.StatusOK, res)
}
