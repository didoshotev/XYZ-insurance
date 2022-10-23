package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/didoshotev/XYZ-insurance/models"
)

type contractController struct {
	contractIDPattern *regexp.Regexp
}

func (cc contractController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/contracts" {
		switch r.Method {
		case http.MethodGet:
			cc.getContracts(w, r)
		case http.MethodPost:
			cc.createContractHandler(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := cc.contractIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("cannot find user with this id"))
		}

		switch r.Method {
		case http.MethodGet:
			cc.getContract(id, w, r)
		case http.MethodPut:
			cc.editContractHandler(id, w, r)
		case http.MethodDelete:
			cc.deleteContractHandler(id, w, r)
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

func (cc *contractController) getContract(id int, w http.ResponseWriter, r *http.Request) {
	contract, err := models.GetContractByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(contract, w)
	w.WriteHeader(http.StatusOK)
}

func (cc *contractController) createContractHandler(w http.ResponseWriter, r *http.Request) {
	contract, err := cc.parseContractRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse user object"))
		return
	}
	contract, err = models.CreateContract(contract)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(contract, w)
	w.WriteHeader(http.StatusOK)
}

func (cc *contractController) editContractHandler(id int, w http.ResponseWriter, r *http.Request) {
	contract, err := cc.parseContractRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse user object"))
		return
	}
	if id != contract.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	contract, err = models.UpdateContract(contract)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(contract, w)
	w.WriteHeader(http.StatusOK)
}

func (cc *contractController) deleteContractHandler(id int, w http.ResponseWriter, r *http.Request) {
	err := models.DeleteContract(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *contractController) parseContractRequest(r *http.Request) (models.Contract, error) {
	dec := json.NewDecoder(r.Body)
	var contract models.Contract
	err := dec.Decode(&contract)
	if err != nil {
		return models.Contract{}, err
	}
	return contract, nil
}
