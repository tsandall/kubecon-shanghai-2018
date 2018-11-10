package main

import (
	"context"
	"log"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/metrics"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
)

func main() {

	ctx := context.Background()

	query := ast.MustParseBody(`data.example.allow = x`)

	input := map[string]interface{}{
		"method": "GET",
		"path":   []string{"salary", "bob"},
		"user":   "alice",
	}

	store := inmem.NewFromObject(map[string]interface{}{
		"managers": map[string]interface{}{
			"bob":   []string{"alice", "fred"},
			"alice": []string{"fred"},
		},
	})

	compiler := ast.MustCompileModules(map[string]string{
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

	metrics := metrics.New()

	r := rego.New(
		rego.Metrics(metrics),
		rego.Trace(true),
		rego.Store(store),
		rego.Compiler(compiler),
		rego.ParsedQuery(query),
		rego.Input(input),
	)

	rs, err := r.Eval(ctx)

	//	rego.PrintTrace(os.Stdout, r)

	if err != nil {
		log.Fatal("error: ", err)
	}

	log.Printf("%+v\n", rs)
	log.Println(metrics)
}
