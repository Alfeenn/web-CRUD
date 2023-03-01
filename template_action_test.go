package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CheckName(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Title": "Simmple",
		"Name":  "Johnny",
	})

}

func TestTemplateAction(t *testing.T) {
	//template action merender text dari perulangan dalam file gohtml
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()

	CheckName(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func checkValue(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/operator.gohtml"))

	t.ExecuteTemplate(w, "operator.gohtml", map[string]interface{}{
		"Title":      "Simple",
		"finalValue": 70,
	})
}

func TestTemplateOperator(t *testing.T) {
	//Melakukan pengecekan operator perbandingan dala file gohtml dan merender value
	//operator dalam template bukan == , !== , <= , < , >= , > tapi eq,lt,le,ge,gt
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()

	checkValue(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func Hobbies(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{"read"},
	})
}

func TestRangeTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()

	Hobbies(recorder, request)
	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Name":    "Bryan",
		"Address": map[string]interface{}{},
	})
}

func TestWithTemplates(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func LayoutTemplates(w http.ResponseWriter, r *http.Request) {
	//membuat template yg dapat menginclude beberapa layout dalam satu file gohtml
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/layout.gohtml", "./templates/footer.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Name":  "Uke",
		"Title": "Template Layout",
	})
}

func TestLayoutTempllate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()
	LayoutTemplates(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}
