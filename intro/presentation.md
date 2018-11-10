## Hello World

- Invoking OPA
    * Passing context
    * Passing simple parameters
- Handling evaluation results

## Rego parameters

- Salary example
- [Reference](https://godoc.org/github.com/open-policy-agent/opa/rego)

## Debugging

- Enabling tracing
- Reading the trace

## Performance

- Metrics
- Cache the parsed query
- Cache the compiler

## gorilla/mux Example

- Input construction

    ```golang
    input := map[string]interface{}{
        "method": r.Method,
        "path":   strings.Split(strings.Trim(r.URL.Path, "/"), "/"),
        "user":   r.Header.Get("Authorization"),
    }
    ```
