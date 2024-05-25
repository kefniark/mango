package products

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/moroz/uuidv7-go"
	"github.com/rs/zerolog"
)

type ProductService struct {
	db     *db.Queries
	logger zerolog.Logger
	api.UnimplementedProductsServer
}

func NewProductService(options *core.ServerOptions) *ProductService {
	return &ProductService{
		db:     options.DB,
		logger: options.Logger.With().Str("service", "ProductService").Logger(),
	}
}

func mapProductSQLToGrpc(product db.Product) *api.ProductData {
	return &api.ProductData{
		Id:   product.ID.(string),
		Name: product.Name,
	}
}

func (service *ProductService) Get(ctx context.Context, req *connect.Request[api.ProductGetRequest]) (*connect.Response[api.ProductData], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}

func (service *ProductService) Set(ctx context.Context, req *connect.Request[api.ProductSetRequest]) (*connect.Response[api.ProductData], error) {
	var id string
	if req.Msg.Id == nil {
		id = uuidv7.Generate().String()
	} else {
		id = req.Msg.GetId()
	}

	product, err := service.db.SetProduct(ctx, db.SetProductParams{
		ID:   id,
		Name: req.Msg.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(mapProductSQLToGrpc(product)), nil
}

func (service *ProductService) Search(
	ctx context.Context,
	req *connect.Request[api.ProductSearchRequest],
) (*connect.Response[api.ProductSearchResponse], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}

func (service *ProductService) Delete(ctx context.Context, req *connect.Request[api.ProductGetRequest]) (*connect.Response[api.ProductEmptyResponse], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}
