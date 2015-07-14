# Iprestrict Negroni middleware

Iprestrict middleware to restrict access for specific IP addresses

## Usage

~~~ go
package main

import (
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/rabeesh/negroni-iprestrict"
    "net/http"
)


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(rw, "Welcome to the home page!")
    })

    rips := []string{"122.12.12.12", "122.12.2.22"}

    n := negroni.New()
    n.Use(iprestrict.New(rips))
    n.UseHandler(mux)
    n.Run(":5000")
}
~~~

## Authors

- Rab [@rabeesh](http://stackoverflow.com/users/1722625/rab)