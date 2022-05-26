package todo_list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type TodoItem struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Competed bool   `json:"competed"`
	ListId   int    `json:"listId"`
}

func post1(w http.ResponseWriter, r *http.Request) {
	var params TodoItem
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(body, &params); err != nil {
		fmt.Println(err)
	}
	query := fmt.Sprintf("INSERT INTO todo_item VALUES (%d, '%s', %t, %d)", rand.Intn(10_000_000), params.Title, false, params.ListId)
	_, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	get(w, r)
}

func _delete1(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := fmt.Sprintf("DELETE FROM todo_item WHERE id=%s", id)
	db.Query(query)
	get(w, r)
}

func patch1(w http.ResponseWriter, r *http.Request) {

}

func TodoItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		post1(w, r)
	case "DELETE":
		_delete1(w, r)
	case "PATCH":
		patch1(w, r)
	}
}
