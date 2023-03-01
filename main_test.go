package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		//logic web
		fmt.Fprint(w, "Hello")
		//Fprintf digunakan untuk konversi writer type byte dan print writer ke web server
	}
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	mux.HandleFunc("/database", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Database page")
	})
	mux.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Database img")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: &mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestReq(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Method)
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: &mux,
	}

	server.ListenAndServe()

}
