package main

import (
	_ "database/sql"

	"net/http"

	"github.com/Alfeenn/web/controller"
)

func main() {
	http.HandleFunc("/", controller.UserIndex)
	http.HandleFunc("/search", controller.Search)
	http.HandleFunc("/user/get_form", controller.Getform)
	http.HandleFunc("/user/store", controller.Store)
	http.HandleFunc("/user/delete", controller.Delete)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("Assets/css"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("Assets/vendor"))))

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}

}
