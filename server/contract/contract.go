package contract

import (
	"net/http"

	common "github.com/didoshotev/XYZ-insurance/common/cors"
)

func RegisterControllers() {
	cc := newContractController()

	http.Handle("/contracts", common.CORSMiddleware(*cc))
	http.Handle("/contract", common.CORSMiddleware(*cc))
	http.Handle("/contract/", common.CORSMiddleware(*cc))
}
