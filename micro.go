package main

import (
	"net/http"
	"log"
	"os"
	"github.com/Prajapati31/goMicro"
)

func main() {
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	hh:=handlers.NewHello()
	http.HandleFunc()
	sm:=http.NewServeMux()
	sm.Handle("/",hh)

	http.ListenAndServe(":9090", nil)
}
