package routers

import(
	"clientes/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	
)

func GetClienteRouter() (*chi.Mux) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/clientes", handlers.HandleGetClientes)
	return router
}