package web

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func GetName(w http.ResponseWriter, r *http.Request) {

	Name := r.URL.Query().Get("Name")
	if Name == "" {
		http.ServeFile(w, r, "Assets/notfound.html") //Menggunakan salah satu file sebagai response di dalam writer
	} else {
		http.ServeFile(w, r, "Assets/")
	}
}

func TestServeFile(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(GetName),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed Assets/notfound.html
var Noresource string

//go:embed Assets
var OK embed.FS

func ServeEmbedFIle(w http.ResponseWriter, r *http.Request) {

	Name := r.URL.Query().Get("Name")
	if Name == "" {
		fmt.Fprint(w, Noresource)
	} else {
		fmt.Fprint(w, OK)
	}
}

func TestServe(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(ServeEmbedFIle),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
