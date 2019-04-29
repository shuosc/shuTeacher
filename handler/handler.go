package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"shuTeacher/model"
	"shuTeacher/service/crawl"
)

func getTeacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	teacher, err := model.Get(id)
	if err != nil {
		teacher, err = crawl.TeacherInfo(id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(404)
			return
		}
	}
	data, _ := json.Marshal(teacher)
	_, _ = w.Write(data)
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTeacherHandler(w, r)
	default:
		w.WriteHeader(405)
	}
}

func PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
