package internal

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api/apiconnect"
	api "github.com/kefniark/go-web-server/internal/api"

	"github.com/kefniark/go-web-server/internal/core"
	"github.com/kefniark/go-web-server/internal/middlewares"
)

func registerAPIRoutes(mux *http.ServeMux, options *core.ServerOptions) {
	interceptors := connect.WithInterceptors(middlewares.WithDevLogInterceptor(options))

	path, handler := apiconnect.NewUsersHandler(api.NewUserService(options), interceptors)
	mux.Handle(path, handler)

	path, handler = apiconnect.NewProductsHandler(api.NewProductService(options), interceptors)
	mux.Handle(path, handler)
}
