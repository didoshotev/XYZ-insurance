package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	cc := newContractController()

	http.Handle("/contracts", *cc)
	http.Handle("/contract", *cc)
	http.Handle("/contract/", *cc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
