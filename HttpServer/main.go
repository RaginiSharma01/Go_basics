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

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to my server!"))
	})

	fmt.Println("the server is runnig on local host 3000")
	http.ListenAndServe(":3000", nil)
}
