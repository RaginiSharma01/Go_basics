package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http server has few tasks
	// 1. process dynamic req,
	// 2. serve static assets
	// 3 Accept connections

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "welcome to my server")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	fmt.Println("the server is runnig on local host 3000")
	http.ListenAndServe(":3000", nil)
}
