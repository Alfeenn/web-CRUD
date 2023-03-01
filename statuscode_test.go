package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseStatus(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(400) //Bad Request
		fmt.Fprintf(w, "Name is empty")
	} else {

		fmt.Fprintf(w, "Hello %s", name)
	}

}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000/?name=", nil)
	recorder := httptest.NewRecorder()

	ResponseStatus(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	result := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(result)

}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000/?name=Lort", nil)
	recorder := httptest.NewRecorder()

	ResponseStatus(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	result := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(result)

}
