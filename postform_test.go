package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormParse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//method otomatis parsing form -> r.PostFormValue("value")
	firstName := r.PostForm.Get("FirstName")
	lastName := r.PostForm.Get("LastName")
	fmt.Fprintf(w, "Hello %s %s ", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("FirstName=Rey&LastName=Mysterio") //mengirim body untuk parameter body io.reader
	//query parameter mengirim value di URL
	//formm post mengirim value di body dan menambahkan header content-type
	request := httptest.NewRequest(http.MethodPost, "localhost:3000", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded") //value standar form post
	recorder := httptest.NewRecorder()

	FormParse(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	result := string(body)
	fmt.Println(result)

}
