package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)

	cookie.Name = "X-PZN-Name"
	cookie.Value = r.URL.Query().Get("Name")
	cookie.Path = "/"

	if cookie.Value == "" {
		fmt.Fprint(w, "Please enter cookie")
	} else {

		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Set cookie success")
	}
}

func RequestCookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprintln(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}

}
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", RequestCookies)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000/?Name=Alfian", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies() //menampilkan response cookies dari http.cookies

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value) //Mengambil cookie dari field http.cookies
	}

}

func TestGetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:3000", nil)
	cookies := new(http.Cookie)
	cookies.Name = "X-PZN-Name"
	cookies.Value = "Johns"
	request.AddCookie(cookies) //mengirim cookie ke server

	recorder := httptest.NewRecorder()

	RequestCookies(recorder, request)

	result := recorder.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))

}
