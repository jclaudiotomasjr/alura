package main

import (
	"fmt"

	"github.com/jclaudiotomasjr/alura/api-go-rest/database"
	"github.com/jclaudiotomasjr/alura/api-go-rest/routes"
)

func main() {

	database.ConectaComBanco()
	fmt.Println("iniciando servidor Rest com Go!")
	routes.HandleRequest()
}
