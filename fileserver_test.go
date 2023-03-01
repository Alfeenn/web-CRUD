package web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed Assets
var Assets embed.FS

func TestFileServer(t *testing.T) {
	directory, _ := fs.Sub(Assets, "Assets")
	Fileserver := http.FileServer(http.FS(directory))
	//fs.Sub -> mengakses/masuk ke dalam directory agar file directory tidak menjadi path url
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", Fileserver)) //strip prefix untuk menghapus prefix url agar dapat membaca path
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
