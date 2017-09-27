package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var host string
var port string

func init() {
	flag.StringVar(&host, "Binding host address ie: 127.0.0.1", "127.0.0.1", "-host=127.0.0.1")
	flag.StringVar(&port, "Binding port", "3000", "-port=3000")

	flag.Parse()
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users/{userid}/log", logThing).Methods("POST")
	router.HandleFunc("/users", createUser).Methods("POST")

	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), n)
}
