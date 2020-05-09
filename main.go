package main

import (
	"net/http"

	"github.com/benacook/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
