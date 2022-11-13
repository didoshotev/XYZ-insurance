package contract

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/didoshotev/XYZ-insurance/common"
)

type contractController struct {
	contractIdPattern *regexp.Regexp
}

func (contrController contractController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/contracts" {
		switch r.Method {
		case http.MethodGet:
			contrController.getContracts(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := contrController.contractIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		contractId, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("cannot find user with this id"))
		}
		switch r.Method {
		case http.MethodGet:
			contrController.getContract(contractId, w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
			return
		case http.MethodPut:
			contrController.updateContract(contractId, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func newContractController() *contractController {
	return &contractController{
		contractIdPattern: regexp.MustCompile(`^/contract/(\d+)/?`),
	}
}

func (contController contractController) getContracts(w http.ResponseWriter, r *http.Request) {
	contractsList, err := getContractList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	contractsListParsed, err := json.Marshal(contractsList)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(contractsListParsed)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (contController contractController) getContract(contractId int, w http.ResponseWriter, r *http.Request) {
	contract, err := getContract(contractId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if contract == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cannot find contract with this id"))
		return
	}
	j, err := json.Marshal(contract)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (contController contractController) updateContract(contractId int, w http.ResponseWriter, r *http.Request) {
	currContract, err := contController.parseContractBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse contract body"))
		return
	}
	if currContract.ContractId != contractId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse contract body"))
		return
	}
	contract, err := updateContract(currContract)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not update contract"))
		return
	}
	common.EncodeResponseAsJSON(contract, w)
}

func (contController contractController) parseContractBody(r *http.Request) (Contract, error) {
	dec := json.NewDecoder(r.Body)
	var contract Contract
	err := dec.Decode(&contract)
	if err != nil {
		return Contract{}, err
	}
	return contract, nil
}
