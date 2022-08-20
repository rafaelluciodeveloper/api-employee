package main

import (
	"fmt"
	"go-api/configs"
	"go-api/route"
	"net/http"
	"os"
)

// @title Employee API documentation
// @version 1.0.0
// @host https://api-employee-v1.herokuapp.com
// @BasePath /employee
func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := route.SetupRoute()

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)

}
