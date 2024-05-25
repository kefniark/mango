package internal

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api/apiconnect"
	productsService "github.com/kefniark/go-web-server/internal/api/products"
	usersService "github.com/kefniark/go-web-server/internal/api/users"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/kefniark/go-web-server/internal/middlewares"
)

func registerRoutes(mux *http.ServeMux, options *core.ServerOptions) {
	interceptors := connect.WithInterceptors(middlewares.WithDevLogInterceptor(options))

	path, handler := apiconnect.NewUsersHandler(usersService.NewUserService(options), interceptors)
	mux.Handle(path, handler)

	path, handler = apiconnect.NewProductsHandler(productsService.NewProductService(options), interceptors)
	mux.Handle(path, handler)
}
