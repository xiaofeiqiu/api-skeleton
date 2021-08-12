package restutils

import (
	"encoding/json"
	"net/http"
)

func JsonResp(w http.ResponseWriter, code int, body interface{}) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshaling body to json"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonBody)
	return
}
