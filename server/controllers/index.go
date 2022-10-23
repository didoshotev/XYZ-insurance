package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/didoshotev/XYZ-insurance/middlewares"
)

func RegisterControllers() {
	cc := newContractController()

	http.Handle("/contracts", middlewares.BaseMiddlewareHandler(*cc))
	http.Handle("/contract", *cc)
	http.Handle("/contract/", *cc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
