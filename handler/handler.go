package handler

import (
	"github.com/vxppjzuowmgnofwjawqk/crud/handler/todo_list"
	"net/http"
)

func GetMux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/list", todo_list.TodoListHandler)
	m.HandleFunc("/item", todo_list.TodoItemHandler)
	return m
}
