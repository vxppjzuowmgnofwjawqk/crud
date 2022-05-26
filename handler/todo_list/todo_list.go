package todo_list

import (
	"encoding/json"
	"fmt"
	"github.com/vxppjzuowmgnofwjawqk/crud/database"
	"io/ioutil"
	"net/http"
)

type TodoList struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	TodoItems []TodoItem `json:"todoItems"`
}

var db = database.GetDB()

func get(w http.ResponseWriter, r *http.Request) {
	var result []TodoList
	rows, err := db.Query("SELECT * FROM todo_list")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var todoList TodoList
		if err := rows.Scan(&todoList.Id, &todoList.Title); err != nil {
			fmt.Println(err)
		}
		rows1, err := db.Query(fmt.Sprintf("SELECT * FROM todo_item WHERE list_id=%d", todoList.Id))
		if err != nil {
			fmt.Println(err)
		}
		for rows1.Next() {
			var todoItem TodoItem
			if err := rows1.Scan(
				&todoItem.Id,
				&todoItem.Title,
				&todoItem.Competed,
				&todoItem.ListId,
			); err != nil {
				fmt.Println(err)
			}
			todoList.TodoItems = append(todoList.TodoItems, todoItem)
		}
		result = append(result, todoList)
	}
	b, err := json.Marshal(result)
	fmt.Fprintf(w, string(b))
}

func post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var params TodoList
	if err = json.Unmarshal(body, &params); err != nil {
		fmt.Println(err)
	}
	db.Query(fmt.Sprintf("INSERT INTO todo_list VALUES (%d, '%s')", params.Id, params.Title))
	get(w, r)
}

func _delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Query(fmt.Sprintf("DELETE FROM todo_list WHERE id=%s", id))
	get(w, r)
}

func patch(w http.ResponseWriter, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		fmt.Println(err)
	} else {
		var params TodoList
		if err = json.Unmarshal(body, &params); err != nil {
			fmt.Println(err)
		}
		db.Query(fmt.Sprintf("UPDATE todo_list SET title='%s' WHERE id=%d", params.Title, params.Id))
	}
	get(w, r)
}

func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		post(w, r)
	case "DELETE":
		_delete(w, r)
	case "PATCH":
		patch(w, r)
	}
}
