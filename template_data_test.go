package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func templateData(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/name.gohtml")) //membuat template.New dan parse file
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Title template",
		"Name":  "Mike",
		"Address": map[string]interface{}{
			"Street":     "No Street",
			"PostalCode": 42162,
		},
	}) //merender setiap field dalam gohtml menggunakan map template
}

func TestDataMapTemplate(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()
	templateData(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

type Address struct {
	Street     string
	PostalCode bool
}
type Page struct {
	Name    string
	Title   string
	Address Address
}

func templateDataSturct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data struct",
		Name:  "Elton",
		Address: Address{
			Street: "Jalan ....",
		},
	})
}

func TestStructTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()

	templateDataSturct(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}
