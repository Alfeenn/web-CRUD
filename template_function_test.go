package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (mypage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + mypage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNC").Parse(`{{.SayHello "Budi"}}`))
	//merender parameter budi milik struct mypage yang mengimplement func sayHello
	t.ExecuteTemplate(w, "FUNC", MyPage{
		Name: "John",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:3000", nil)
	recorder := httptest.NewRecorder()
	TemplateFunction(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse("{{len .Name}}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "John",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()
	TemplateFunctionGlobal(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	//membuat function global menggunakan template.funcs
	//sebelum memparsing value, register function global terlebih dahulu
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse("{{upper .Name}}"))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "yeti",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()
	TemplateFunctionCreateGlobal(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	//function pipeline -> mengirim hasil function ke function berikutnya
	//function pipeline menggunakan simbol |
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse("{{sayHello .Name | upper}}"))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "yeti",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()
	TemplateFunctionPipeline(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))

}
