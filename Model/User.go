package model

import (
	"database/sql"

	"github.com/Alfeenn/web/Entities"
	"github.com/Alfeenn/web/db"
	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	db *sql.DB
}

func New() *UserModel {
	db, err := db.Conn()
	if err != nil {
		panic(err)
	}
	return &UserModel{db: db}
}

func (u *UserModel) FindAll(users *[]Entities.Alluser) error {

	Query := "SELECT id,username,nama,alamat,umur FROM user "
	row, err := u.db.Query(Query)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		var data Entities.Alluser
		row.Scan(&data.Id,
			&data.Username,
			&data.Nama,
			&data.Alamat,
			&data.Umur)

		*users = append(*users, data)

	}
	return nil
}

func (u *UserModel) Delete(id int64) error {

	_, err := u.db.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (u *UserModel) Search(name string, user *[]Entities.Alluser) error {
	row, err := u.db.Query("SELECT id,username,nama,alamat,umur FROM user WHERE name=?", name)

	if err != nil {
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		var data Entities.Alluser
		row.Scan(
			&data.Id,
			&data.Username,
			&data.Nama,
			&data.Alamat,
			&data.Umur,
		)
		*user = append(*user, data)
	}
	return nil
}

func (u *UserModel) Find(id int64, user *Entities.Alluser) error {
	return u.db.QueryRow("SELECT id,username,nama,alamat,umur FROM user WHERE id=?", id).Scan(
		&user.Id,
		&user.Username,
		&user.Nama,
		&user.Alamat,
		&user.Umur,
	)

}

func (u *UserModel) Create(user *Entities.Alluser) error {
	result, err := u.db.Exec("INSERT INTO user (username,nama,alamat,umur) values(?,?,?,?)",
		user.Username, user.Nama, user.Alamat, user.Umur)
	if err != nil {
		panic(err)
	}
	lastInsertId, _ := result.LastInsertId()
	user.Id = lastInsertId
	return nil
}

func (u *UserModel) Update(id int64, user *Entities.Alluser) error {
	_, err := u.db.Exec("Update user set username=?, nama=?, alamat=?, umur=? WHERE id =?",
		user.Username, user.Nama, user.Alamat, user.Umur, user.Id)
	if err != nil {
		panic(err)
	}
	return nil
}
