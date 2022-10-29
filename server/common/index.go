package common

import (
	"encoding/json"
	"io"
)

func EncodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
