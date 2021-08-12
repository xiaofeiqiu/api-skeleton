package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xiaofeiqiu/api-skeleton/lib/logger"
	rr "github.com/xiaofeiqiu/api-skeleton/lib/restutils"
	"github.com/xiaofeiqiu/api-skeleton/services"
)

type PizzaHandler struct {
	PizzaService services.PizzaService

	// if you want do talk to db, you can add db instance here
}

func (s *PizzaHandler) CreatePizza(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("CreatePizza", err, "error reading body")
		rr.JsonResp(w, 500, rr.Message{Message: "error reading request body", Error: err.Error()})
		return
	}

	var req services.CreatePizzaRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		rr.JsonResp(w, 500, rr.Message{Message: "unmarshal request failed", Error: err.Error()})
		return
	}

	resp, err := s.PizzaService.MakePizza(req)
	if err != nil {
		rr.JsonResp(w, 400, rr.Message{Message: "create pizza failed", Error: err.Error()})
		return
	}

	rr.JsonResp(w, 200, resp)
	return
}

func (s *PizzaHandler) UpdatePizza(w http.ResponseWriter, r *http.Request) {

	rr.JsonResp(w, 200, rr.Message{Message: "updated"})
	return
}

func (s *PizzaHandler) DeletePizza(w http.ResponseWriter, r *http.Request) {

	rr.JsonResp(w, 200, rr.Message{Message: "deleted"})
	return
}

func (s *PizzaHandler) GetPizza(w http.ResponseWriter, r *http.Request) {

	rr.JsonResp(w, 200, rr.Message{Message: "successful"})
	return
}
