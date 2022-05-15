package main

import (
	"github.com/guilhermesoaresmarques/gin-api-rest/database"
	"github.com/guilhermesoaresmarques/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
