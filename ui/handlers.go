package ui

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func IndexHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("index handler")
	})
}

func RegisterHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		if r.Method != "POST" {
			return
		}

		fmt.Println("registrationHandler")
		// parsing
		if err := r.ParseForm(); err != nil {
			panic(err.Error())
		}
		login := r.Form.Get("login")
		password := r.Form.Get("password")
		email := r.Form.Get("email")
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			age = -1
		}
		fmt.Println(login, password, email, age)
		// TODO: validation, hashing password

		user := &model.User{
			Login:login,
			Password: password,
			Email: email,
			Age: age}

		lastId, err := m.InsertUser(user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("last_id = %d", lastId)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}

func ShowRegisterFormHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		// Обработка формы
		b, err := ioutil.ReadFile("assets/views/registration/index.html")
		if err != nil {
			panic(err)
		}

		// показываем нашу страницу
		w.Write(b)
	})
}
