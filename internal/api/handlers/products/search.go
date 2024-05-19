package products

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
)

func (service *ProductService) Search(
	ctx context.Context,
	req *connect.Request[api.ProductSearchRequest],
) (*connect.Response[api.ProductSearchResponse], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}
