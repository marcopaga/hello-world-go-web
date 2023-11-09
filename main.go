package main

import (
    "fmt"
    "net/http"
    "math/big"
    "strconv"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

func calculateFactorial(n int64) *big.Int {
    if n < 0 {
        return big.NewInt(0)
    }
    
    result := big.NewInt(1)
    for i := int64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(i))
    }
    
    return result
}

func factorialHandler(w http.ResponseWriter, r *http.Request) {
    // Get the "number" query parameter from the URL, or use 10 as a default.
    param := r.URL.Query().Get("number")
    n, err := strconv.ParseInt(param, 10, 64)
    if err != nil {
        n = 100000
    }

    factorial := calculateFactorial(n)
    fmt.Fprintf(w, "%d! = %s\n", n, factorial.String())
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    http.HandleFunc("/factorial", factorialHandler)
    http.ListenAndServe(":8080", nil)
}
