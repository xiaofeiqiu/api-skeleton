package services

import (
	"fmt"
	"time"

	"github.com/xiaofeiqiu/api-skeleton/lib/logger"
)

type PizzaService struct {
}

type CreatePizzaRequest struct {
	Size int `json:"size"`
}

type CreatePizzaResonse struct {
	Size       int
	Price      string
	CreateDate time.Time
}

func (s *PizzaService) MakePizza(request CreatePizzaRequest) (*CreatePizzaResonse, error) {
	logger.Info("MakePizza", "Creating pizza")

	if request.Size > 10 {
		return nil, fmt.Errorf("pizza too big")
	}

	return &CreatePizzaResonse{
		Size:       request.Size,
		Price:      "$10",
		CreateDate: time.Now(),
	}, nil
}
