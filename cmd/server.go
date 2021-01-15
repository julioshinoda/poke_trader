package main

import (
	"fmt"
	"net/http"
	"os"
	"poketrader/internal/trade"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		trade.Handler(r)
	})

	return r
}
