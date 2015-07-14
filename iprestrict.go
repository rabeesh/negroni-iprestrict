package iprestrict

import (
    "log"
    "os"
    "strings"
    "net/http"
)

type IpRestrict struct {
    Ips []string
    Logger  *log.Logger
}

func New(ips []string) *IpRestrict {
    return &IpRestrict{
        Ips: ips,
        Logger: log.New(os.Stdout, "[negroni ip restrict is enabled] ", 0),
    }
}

func (ir *IpRestrict) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    ip := strings.Split(req.RemoteAddr,":")[0]
    if ir.search(ip) {
        rw.WriteHeader(http.StatusForbidden);
        return
    }
    next(rw, req)
}

func (ir *IpRestrict) search(find string) bool {
    for _, b := range ir.Ips {
        if b == find {
            return true
        }
    }
    return false
}