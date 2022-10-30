package customer

import (
	"net/http"

	common "github.com/didoshotev/XYZ-insurance/common/cors"
)

func RegisterCustomerControllers() {
	customerController := newCustomerController()

	http.Handle("/customers", common.CORSMiddleware(*customerController))
	http.Handle("/customer", common.CORSMiddleware(*customerController))
	http.Handle("/customer/", common.CORSMiddleware(*customerController))
}
