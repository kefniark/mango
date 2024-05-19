package products

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/moroz/uuidv7-go"
)

func (service *ProductService) Set(
	ctx context.Context,
	req *connect.Request[api.ProductSetRequest],
) (*connect.Response[api.ProductData], error) {
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
