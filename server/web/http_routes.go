package web

import (
	"net/http"

	_ "clothshare/web/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter()
	fe := router.PathPrefix("/api").Subrouter()

	// product endpoints
	fe.Path("/products/{id}/image").HandlerFunc(app.ProductStoreImage).Methods(http.MethodPost)
	fe.Path("/products/{id}/image").HandlerFunc(app.ProductGetImage).Methods(http.MethodGet)
	fe.Path("/products/user/{id}").HandlerFunc(app.ProductGetListByUser).Methods(http.MethodGet)
	fe.Path("/products/{id}").HandlerFunc(app.ProductDelete).Methods(http.MethodDelete)
	fe.Path("/products/{id}").HandlerFunc(app.ProductGet).Methods(http.MethodGet)
	fe.Path("/products").HandlerFunc(app.ProductGetList).Methods(http.MethodGet)
	fe.Path("/products").HandlerFunc(app.ProductCreate).Methods(http.MethodPost)

	// user endpoints
	fe.Path("/users/{id}/image").HandlerFunc(app.UserStoreImage).Methods(http.MethodPost)
	fe.Path("/users/{id}").HandlerFunc(app.UserDelete).Methods(http.MethodDelete)
	fe.Path("/users/{id}").HandlerFunc(app.UserGet).Methods(http.MethodGet)
	fe.Path("/users").HandlerFunc(app.UserGetList).Methods(http.MethodGet)
	fe.Path("/users").HandlerFunc(app.UserCreate).Methods(http.MethodPost)

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	return router
}
