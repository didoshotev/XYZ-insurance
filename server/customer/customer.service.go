package customer

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/didoshotev/XYZ-insurance/common"
)

type customerController struct {
	customerIdPattern *regexp.Regexp
}

func (custController customerController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/customers" {
		switch r.Method {
		case http.MethodGet:
			custController.getCustomers(w, r)
		case http.MethodPost:
			custController.createCustomer(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := custController.customerIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		customerId, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("cannot find user with this id"))
		}
		switch r.Method {
		case http.MethodGet:
			custController.getCustomer(customerId, w, r)
		case http.MethodDelete:
			custController.deleteCustomer(customerId, w, r)
		case http.MethodPut:
			custController.updateCustomer(customerId, w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func newCustomerController() *customerController {
	return &customerController{
		customerIdPattern: regexp.MustCompile(`^/customer/(\d+)/?`),
	}
}

func (custController customerController) getCustomers(w http.ResponseWriter, r *http.Request) {
	customerList, err := getCustomerList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	customerListParsed, err := json.Marshal(customerList)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(customerListParsed)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (custController customerController) getCustomer(customerId int, w http.ResponseWriter, r *http.Request) {
	customer, err := getCustomer(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cannot find user with this id"))
		return
	}
	j, err := json.Marshal(customer)
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

func (custController customerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	customer, err := custController.parseCustomerBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse customer body"))
		return
	}
	_, err = createCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse customer body"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (custController customerController) updateCustomer(customerId int, w http.ResponseWriter, r *http.Request) {
	currCustomer, err := custController.parseCustomerBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse customer body"))
		return
	}
	if currCustomer.CustomerId != customerId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse customer body"))
		return
	}
	customer, err := updateCustomer(currCustomer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not update customer"))
		return
	}
	common.EncodeResponseAsJSON(customer, w)
}

func (custController customerController) deleteCustomer(customerId int, w http.ResponseWriter, r *http.Request) {
	err := removeCustomer(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (custController customerController) parseCustomerBody(r *http.Request) (Customer, error) {
	dec := json.NewDecoder(r.Body)
	var customer Customer
	err := dec.Decode(&customer)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}
