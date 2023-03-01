package controller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Alfeenn/web/Entities"
	"github.com/Alfeenn/web/model"
)

var userModel = model.New()

func UserIndex(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	var user Entities.Alluser
	var data map[string]interface{}
	name := r.URL.Query().Get("nama") //query parameter name pada url
	if name == "" {
		data = map[string]interface{}{
			"data": template.HTML(GetData()),
		}
	} else {
		err := userModel.Search(name, &user)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"user": user,
		}
	}

	temp, _ := template.ParseFiles("Assets/index.html")
	temp.ExecuteTemplate(w, "Index", data)
}

func GetData() string {
	buffer := &bytes.Buffer{}

	t, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("Assets/data.html")
	var table []Entities.Alluser
	var data map[string]interface{}
	err := userModel.FindAll(&table)

	if err != nil {
		panic(err)
	}
	data = map[string]interface{}{
		"Table": table,
	}
	t.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()

}

func Getform(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}
	var user Entities.Alluser
	if err != nil {
		data = map[string]interface{}{
			"title": "Tambahdata",
		}
	} else {
		err := userModel.Find(id, &user)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title": "Edit Data Mahasiswa",
			"user":  user,
		}
	}
	t := template.Must(template.ParseFiles("Assets/form.html"))
	t.ExecuteTemplate(w, "form.html", data)

}
func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		//Mengambil key value untuk fungsi create dan update dari form html
		var user Entities.Alluser
		user.Username = r.Form.Get("username")
		user.Nama = r.Form.Get("nama")
		user.Alamat = r.Form.Get("alamat")
		user.Umur = r.Form.Get("umur")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
		var data map[string]interface{}

		if err != nil {
			err := userModel.Create(&user)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil disimpan",
				"data":    template.HTML(GetData()),
			}

		} else {
			user.Id = id
			err := userModel.Update(id, &user)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil dibuah",
				"data":    template.HTML(GetData()),
			}

		}
		ResponseJson(w, http.StatusOK, data)
	}

}
func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = userModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data berhasil dihapus",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)

}

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user []Entities.Alluser
	var data map[string]interface{}
	name := r.Form.Get("Nama")
	if name == "" {

		data = map[string]interface{}{
			"data": template.HTML(GetData()),
		}
	} else {
		err := userModel.Search(name, &user)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"user": user,
		}
	}

	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
