package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000/hello", nil)
	recorder := httptest.NewRecorder()

	Hellohandler(recorder, request)

	response := recorder.Result()        //menampilkan result dari response di handler
	body, _ := io.ReadAll(response.Body) //membaca body menggunakan io.ReadAll->returnnya berupa []byte,error
	result := string(body)               //konversi value menjadi string
	fmt.Println(result)

}
