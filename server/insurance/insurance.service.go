package insurance

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/didoshotev/XYZ-insurance/common"
)

type InsuranceController struct {
	insuranceIdPattern *regexp.Regexp
}

func (ic InsuranceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/insurances" {
		switch r.Method {
		case http.MethodGet:
			ic.getInsurances(w, r)
		case http.MethodPost:
			ic.createInsurance(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := ic.insuranceIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		insuranceId, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("cannot find insurance with this id"))
		}
		switch r.Method {
		case http.MethodGet:
			ic.getInsurance(insuranceId, w, r)
		case http.MethodDelete:
			ic.deleteInsruance(insuranceId, w, r)
		case http.MethodPut:
			ic.updateInsurance(insuranceId, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func newInsuranceController() *InsuranceController {
	return &InsuranceController{
		insuranceIdPattern: regexp.MustCompile(`^/insurance/(\d+)/?`),
	}
}

func (ic InsuranceController) getInsurances(w http.ResponseWriter, r *http.Request) {
	insuranceList, err := getInsuranceList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	insuranceListParsed, err := json.Marshal(insuranceList)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(insuranceListParsed)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ic InsuranceController) getInsurance(insuranceId int, w http.ResponseWriter, r *http.Request) {
	insurance, err := getInsurance(insuranceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if insurance == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cannot find insurance with this id"))
		return
	}
	j, err := json.Marshal(insurance)
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

func (ic InsuranceController) createInsurance(w http.ResponseWriter, r *http.Request) {
	insurance, err := ic.parseInsuranceBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse insurance body"))
		return
	}
	_, err = createInsurance(insurance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse insurance body"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ic InsuranceController) updateInsurance(insuranceId int, w http.ResponseWriter, r *http.Request) {
	currInsurance, err := ic.parseInsuranceBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse insurance body"))
		log.Fatal(err)
		return
	}
	if currInsurance.InsuranceId != insuranceId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse insurance body"))
		log.Fatal(err)
		return
	}
	fmt.Println("currInsurance: ", currInsurance)
	insurance, err := updateInsurance(currInsurance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not update insurance"))
		log.Fatal(err)
		return
	}
	common.EncodeResponseAsJSON(insurance, w)
}

func (ic InsuranceController) deleteInsruance(insuranceId int, w http.ResponseWriter, r *http.Request) {
	err := removeInsurance(insuranceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ic InsuranceController) parseInsuranceBody(r *http.Request) (Insurance, error) {
	dec := json.NewDecoder(r.Body)
	var insurance Insurance
	err := dec.Decode(&insurance)
	if err != nil {
		return Insurance{}, err
	}
	return insurance, nil
}
