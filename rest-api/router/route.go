package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"sync"
)

var lock = &sync.Mutex{}

var routeInstance *chi.Mux

// GetRoute Create Route Instance as a singleton.
func GetRoute() *chi.Mux {
	if routeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if routeInstance == nil {
			fmt.Printf("Creating route instance...")
			routeInstance = chi.NewRouter()

			routeInstance.Use(middleware.RequestID)
			routeInstance.Use(middleware.Logger)
		}
	}
	return routeInstance
}
