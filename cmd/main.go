package main

import (
	"github.com/vxppjzuowmgnofwjawqk/crud/handler"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:1440", handler.GetMux())
}
