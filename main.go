package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/xiaofeiqiu/api-skeleton/handlers"
	"github.com/xiaofeiqiu/api-skeleton/lib/logger"
	"github.com/xiaofeiqiu/api-skeleton/services"
)

const (
	Timeout  = 60
	Throttle = 10
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(logger.NewMiddlewareLogger()) //  you dont have to have this
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(Timeout * time.Second))
	r.Use(middleware.Throttle(Throttle))

	pizzaHandler := handlers.PizzaHandler{
		PizzaService: services.PizzaService{},
	}

	// health check
	r.Get("/example/health", handlers.Health)

	// resources for pizza
	r.Post("/example/createPizza", pizzaHandler.CreatePizza)
	r.Put("/example/updatePizza", pizzaHandler.UpdatePizza)
	r.Delete("/example/deletePizza", pizzaHandler.DeletePizza)
	r.Get("/example/getPizza", pizzaHandler.GetPizza)

	http.ListenAndServe(":8080", r)
	logger.Info("Init", "Server started")
}
