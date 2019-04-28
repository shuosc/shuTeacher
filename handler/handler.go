package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shuTeacher/model"
	"shuTeacher/service/token"
)

func getTeacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenInHeader := r.Header.Get("Authorization")
	if tokenInHeader == "" {
		w.WriteHeader(401)
		return
	}
	tokenString := tokenInHeader[7:]
	if token.ValidateToken(tokenString) == token.Invalid {
		w.WriteHeader(403)
		return
	}
	id := r.URL.Query().Get("id")
	teacher, err := model.Get(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	data, _ := json.Marshal(teacher)
	_, _ = w.Write(data)
}

func setTeacherHandler(w http.ResponseWriter, r *http.Request) {
	tokenInHeader := r.Header.Get("Authorization")
	if tokenInHeader == "" {
		w.WriteHeader(401)
		return
	}
	tokenString := tokenInHeader[7:]
	if token.ValidateToken(tokenString) != token.Admin {
		w.WriteHeader(403)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var input model.Teacher
	_ = json.Unmarshal(body, &input)
	model.Save(input)
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTeacherHandler(w, r)
	case "POST":
		setTeacherHandler(w, r)
	case "PUT":
		setTeacherHandler(w, r)
	}
}

func PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
