package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") //query parameter name pada url
	if name == "" {
		fmt.Fprintln(w, "Hello world")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParams(t *testing.T) {
	request1 := httptest.NewRequest(http.MethodGet, "localhost:3000/?name", nil)
	request2 := httptest.NewRequest(http.MethodGet, "localhost:3000/?name=john", nil)
	recorder := httptest.NewRecorder() //methood recorder yang me return response
	Handler(recorder, request1)
	Handler(recorder, request2)
	response := recorder.Result() //return response yg digenerate oleh handler
	body, _ := io.ReadAll(response.Body)
	result := string(body)

	fmt.Println(result)

}
func MultipleQueryParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"] //mapping query multpile value
	Multiple := strings.Join(names, " ")
	if Multiple == "" {
		fmt.Fprintln(w, "Hello guest!")

	} else {
		fmt.Fprintf(w, "Hello %s !", Multiple)
	}
}
func TestNoParams(t *testing.T) {
	request1 := httptest.NewRequest(http.MethodGet, "localhost:3000/?name", nil)
	request2 := httptest.NewRequest(http.MethodGet, "localhost:3000/?name=john&name=Mike&name=Jake", nil)
	recorder := httptest.NewRecorder() //methood recorder yang me return response
	MultipleQueryParams(recorder, request1)
	MultipleQueryParams(recorder, request2)
	response := recorder.Result() //return response yg digenerate oleh handler
	body, _ := io.ReadAll(response.Body)
	result := string(body)

	fmt.Println(result)

}
