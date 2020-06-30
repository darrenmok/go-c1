package main

import (
	"fmt"
	"net/http"

	"github.com/darrenmok/go-c1/visitor"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", visitor.Get)
	fmt.Println("Server started at port 3000!")
	http.ListenAndServe(":3000", r)
}
