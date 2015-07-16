package iprestrict

import (
    "net/http"
    "net/http/httptest"
    "fmt"
    "testing"
)


func Test_ValidIps(t *testing.T) {

    rips := []string{"122.12.12.12", "122.12.2.22"}
    ir := New(rips)

    rw := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "http://localhost/foobar", nil)

    req.RemoteAddr = "122.10.10.10"
    if err != nil {
        t.Fatal(err)
    }

    ir.ServeHTTP(rw, req, func (rw http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(rw, "")
    })

    if rw.Code == http.StatusForbidden {
        t.Error("Expected a valid response code")
    }
}



func Test_IpRestrict(t *testing.T) {

    rips := []string{"122.12.12.12", "122.12.2.22"}
    ir := New(rips)

    rw := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "http://localhost/foobar", nil)

    req.RemoteAddr = "122.12.12.12"
    if err != nil {
        t.Fatal(err)
    }

    ir.ServeHTTP(rw, req, func (rw http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(rw, "")
    })

    if rw.Code != http.StatusForbidden {
        t.Error("Expected a valid forbidden response code")
    }
}

