package main

import (
	"net/http"
	"strings"
	"fmt"
	"learn/mysql/model"
	"encoding/json"
	"time"
)

type ActionResult struct {
	Msg string `"json:msg"`
	UserList [] model.User `"json:userList"`
}

func rootRequest(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Path
	msg = strings.TrimPrefix(msg, "/")
	msg = fmt.Sprintf("Hello %s", msg)
	userList := []model.User{{Name: "Thai", Birthday: time.Now()}}
	actionResult := ActionResult{Msg: msg, UserList: userList}
	json.NewEncoder(w).Encode(actionResult)
}

func main() {
	http.HandleFunc("/", rootRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}