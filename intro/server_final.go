package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
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

var store = inmem.NewFromObject(map[string]interface{}{
	"managers": map[string]interface{}{
		"bob":   []string{"alice", "fred"},
		"alice": []string{"fred"},
	},
})

var compiler = ast.MustCompileModules(map[string]string{
	"example.rego": `
		package example

		import data.managers

		default allow = false

		allow = true {
		  input.method = "GET"
		  input.path = ["salary", employee_id]
		  input.user = employee_id
		}

		allow = true {
		  input.method = "GET"
		  input.path = ["salary", employee_id]
		  input.user = managers[employee_id][_]
		}
	`,
})

var query = ast.MustParseBody(`data.example.allow = true`)

func authorizer(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		input := map[string]interface{}{
			"method": r.Method,
			"path":   strings.Split(strings.Trim(r.URL.Path, "/"), "/"),
			"user":   r.Header.Get("Authorization"),
		}

		eval := rego.New(
			rego.Compiler(compiler),
			rego.Store(store),
			rego.Input(input),
			rego.ParsedQuery(query),
		)

		rs, err := eval.Eval(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if len(rs) == 0 {
			w.WriteHeader(http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
