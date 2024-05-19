package products

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
)

func (service *ProductService) Delete(
	ctx context.Context,
	req *connect.Request[api.ProductGetRequest],
) (*connect.Response[api.ProductEmptyResponse], error) {
	return nil, errors.New("ProductService.delete is not implemented")
}
