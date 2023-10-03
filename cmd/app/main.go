package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/base-media-cloud/base-test/config"
	"github.com/base-media-cloud/base-test/internal/adapters/delivery/royalmail"
	"github.com/base-media-cloud/base-test/internal/adapters/delivery/ups"
	"github.com/base-media-cloud/base-test/internal/adapters/delivery/yodal"
	"github.com/base-media-cloud/base-test/internal/adapters/handlers/gin"
	"github.com/base-media-cloud/base-test/internal/app"
	"github.com/base-media-cloud/base-test/internal/domain/models/product"
	"github.com/base-media-cloud/base-test/internal/domain/ports"
	"github.com/base-media-cloud/base-test/internal/domain/services"
	"github.com/base-media-cloud/base-test/internal/logger"
)

func main() {
	l := logger.New()

	f, err := os.ReadFile("./products.json")
	if err != nil {
		l.Fatal().Err(err).Msg("unable to read products file")
	}

	var prods product.Products
	if err := json.Unmarshal(f, &prods); err != nil {
		l.Fatal().Err(err).Msg("unable to unmarshal products")
	}

	cfg, err := config.New()
	if err != nil {
		l.Fatal().Err(err).Msg("unable to create config")
	}

	var delSvc ports.DeliverySvc
	switch cfg.DeliverySvc {
	case ups.Type:
		delSvc = ups.New(cfg.Price)
	case royalmail.Type:
		delSvc = royalmail.New(cfg.Price)
	case yodal.Type:
		delSvc = yodal.New(cfg.Price)
	default:
		l.Fatal().
			Err(errors.New("invalid delivery service provider")).
			Str("delivery_service", cfg.DeliverySvc).
			Msg("invalid delivery service")
	}

	svc := services.New(delSvc)
	h := gin.New(svc, prods, l)

	a := app.New(cfg, h, l)
	if err = a.Start(); err != nil {
		l.Fatal().Err(err).Msg("server closed")
	}
}
