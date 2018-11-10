package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var salaries = map[string]string{
	"bob":   "$75,000",
	"alice": "$125,000",
	"fred":  "$500,000",
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/salary/{id}", salaryGet).Methods(http.MethodGet)
	router.Use(authorizer)
	log.Println("Serving on port 8000")
	http.ListenAndServe(":8000", router)
}

func salaryGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if sal, ok := salaries[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		bs, _ := json.Marshal(sal)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bs)
	}
}

func authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Implement authorization check here.
		// - Return 500 on error.
		// - Return 403 on authorization failure.
		// - Call next handler on autohrization success.

		next.ServeHTTP(w, r)
	})
}
