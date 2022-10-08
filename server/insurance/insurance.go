package insurance

import (
	"net/http"

	common "github.com/didoshotev/XYZ-insurance/common/cors"
)

func RegisterInsuranceControllers() {
	ic := newInsuranceController()

	http.Handle("/insurances", common.CORSMiddleware(*ic))
	http.Handle("/insurance", common.CORSMiddleware(*ic))
	http.Handle("/insurance/", common.CORSMiddleware(*ic))
}
