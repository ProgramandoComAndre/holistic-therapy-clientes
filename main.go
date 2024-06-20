package main

import (
	
	"net/http"
	"clientes/routers"
)

func main() {

	
	router := routers.GetClienteRouter()

	http.ListenAndServe(":3333", router)



}