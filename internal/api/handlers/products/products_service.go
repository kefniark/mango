package products

import (
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/rs/zerolog"
)

type ProductService struct {
	db     *db.Queries
	logger zerolog.Logger
	api.UnimplementedProductsServer
}

func NewProductService(db *db.Queries, logger *zerolog.Logger) *ProductService {
	return &ProductService{
		db:     db,
		logger: logger.With().Str("service", "ProductService").Logger(),
	}
}

func mapProductSQLToGrpc(product db.Product) *api.ProductData {
	return &api.ProductData{
		Id:   product.ID.(string),
		Name: product.Name,
	}
}
