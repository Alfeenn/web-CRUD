package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeaderRequest(w http.ResponseWriter, r http.Request) {
	header := r.Header.Get("content-type")
	fmt.Fprint(w, header)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	HeaderRequest(recorder, *request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	result := string(body)

	fmt.Println(result)

}

func ResponseHeader(w http.ResponseWriter, r http.Request) {
	w.Header().Add("X-Powered-By", "Mike Phelan")
	fmt.Fprintf(w, "OK")

}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, *request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	result := string(body)
	fmt.Println(result)
	poweredBY := recorder.Header().Get("X-powered-by") //mendapatkan response dari header

	fmt.Println(poweredBY)
}
