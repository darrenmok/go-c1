package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	fmt.Println("Server started at port 3000!")
	http.ListenAndServe(":3000", r)
}
