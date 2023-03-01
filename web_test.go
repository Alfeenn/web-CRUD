package web

import (
	_ "database/sql"

	"net/http"
	"testing"

	"github.com/Alfeenn/web/controller"
)

func TestWeb(t *testing.T) {
	http.HandleFunc("/", controller.UserIndex)
	http.HandleFunc("/user/get_form", controller.Getform)
	http.HandleFunc("/user/store", controller.Store)
	http.HandleFunc("/user/delete", controller.Delete)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}

}
