package gin_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/base-media-cloud/base-test/internal/adapters/delivery/royalmail"
	"github.com/base-media-cloud/base-test/internal/adapters/delivery/ups"
	"github.com/base-media-cloud/base-test/internal/adapters/delivery/yodal"
	handler "github.com/base-media-cloud/base-test/internal/adapters/handlers/gin"
	"github.com/base-media-cloud/base-test/internal/domain/models/product"
	"github.com/base-media-cloud/base-test/internal/domain/ports"
	"github.com/base-media-cloud/base-test/internal/domain/services"
	"github.com/base-media-cloud/base-test/internal/logger"
)

func TestHandler_get(t *testing.T) {
	tests := map[string]struct {
		deliSvc ports.DeliverySvc
	}{
		"success - ups delivery service": {
			deliSvc: ups.New(10),
		},
		"success - royal mail delivery service": {
			deliSvc: royalmail.New(10),
		},
		"success - yodal delivery service": {
			deliSvc: yodal.New(10),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			prods := product.Products{
				{
					Name:   "mock-product-1",
					Price:  1.00,
					Weight: 1.00,
				},
				{
					Name:   "mock-product-2",
					Price:  2.00,
					Weight: 2.00,
				},
				{
					Name:   "mock-product-3",
					Price:  3.00,
					Weight: 3.00,
				},
				{
					Name:   "mock-product-4",
					Price:  4.00,
					Weight: 4.00,
				},
			}

			mockSvc := services.New(test.deliSvc)
			h := handler.New(mockSvc, prods, logger.New())

			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)

			req := &http.Request{
				URL:    &url.URL{Scheme: "http", Host: "mock-host.com"},
				Header: make(http.Header),
				Method: http.MethodGet,
			}

			c.Request = req

			h.Get(c)

			expected := `[
        {
          "name": "mock-product-1",
          "product_price": "1.00",
          "delivery_price": "10.00",
          "total_price": "11.00"
        },
        {
          "name": "mock-product-2",
          "product_price": "2.00",
          "delivery_price": "20.00",
          "total_price": "22.00"
        },
        {
          "name": "mock-product-3",
          "product_price": "3.00",
          "delivery_price": "30.00",
          "total_price": "33.00"
        },
        {
          "name": "mock-product-4",
          "product_price": "4.00",
          "delivery_price": "40.00",
          "total_price": "44.00"
        }
      ]`

			assert.JSONEq(t, expected, rec.Body.String())
		})
	}
}
