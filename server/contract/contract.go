package contract

import (
	"net/http"

	common "github.com/didoshotev/XYZ-insurance/common/cors"
)

func RegisterContractControllers() {
	contractController := newContractController()

	// http.Handle("/contracts", common.CORSMiddleware(common.AuthMiddleware(*contractController)))
	http.Handle("/contracts", common.CORSMiddleware(*contractController))
	http.Handle("/contract", common.CORSMiddleware(*contractController))
	http.Handle("/contract/", common.CORSMiddleware(*contractController))
}
