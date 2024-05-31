package handlers

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/kefniark/mango/example/codegen/api"
	"github.com/kefniark/mango/example/codegen/database"
	"github.com/kefniark/mango/example/config"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	api.UnimplementedProductsServer
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func mapProductSQLToGrpc(product database.Product) *api.ProductData {
	return &api.ProductData{
		Id:   product.ID.String(),
		Name: product.Name,
	}
}

func (service *ProductService) Get(ctx context.Context, req *connect.Request[api.ProductGetRequest]) (*connect.Response[api.ProductData], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}

func (service *ProductService) Set(ctx context.Context, req *connect.Request[api.ProductSetRequest]) (*connect.Response[api.ProductData], error) {
	db := config.GetDB(ctx)

	id := getUUID(req.Msg.GetId())
	product, err := db.SetProduct(ctx, database.SetProductParams{
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

func (service *ProductService) Delete(ctx context.Context, req *connect.Request[api.ProductGetRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}
