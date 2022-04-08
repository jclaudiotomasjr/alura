package main

import (
	"github.com/jclaudiotomasjr/alura/api-go-gin/database"
	"github.com/jclaudiotomasjr/alura/api-go-gin/routes"
)

func main() {
	database.ConectaComBanco()
	routes.HandleRequests()
}
