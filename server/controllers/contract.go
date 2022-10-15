package controllers

import (
	"net/http"
	"regexp"

	"github.com/didoshotev/XYZ-insurance/models"
)

type contractController struct {
	contractIDPattern *regexp.Regexp
}

func (cc contractController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/contract" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("get contract"))
		case http.MethodPost:
			cc.createContractHandler(w, r)
		case http.MethodPut:
			w.Write([]byte("update contract"))
		case http.MethodDelete:
			w.Write([]byte("delete contract"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
		// } else if r.URL.Path == "/contracts" {
	} else {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Get all contracts"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// Constructor
func newContractController() *contractController {
	return &contractController{
		contractIDPattern: regexp.MustCompile(`^/contract/(\d+)/?`),
	}
}

// cc methods
func (cc *contractController) getContracts(w http.ResponseWriter, r *http.Request) {
	contracts := models.GetContracts()
	encodeResponseAsJSON(contracts, w)
}

func (cc *contractController) createContractHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create contract handler"))
}
