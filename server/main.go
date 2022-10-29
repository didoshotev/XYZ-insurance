package main

import (
	"net/http"

	"github.com/didoshotev/XYZ-insurance/contract"
)

func main() {
	contract.RegisterControllers()
	http.ListenAndServe(":5000", nil)
}
