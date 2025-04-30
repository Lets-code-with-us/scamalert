package main

import (
	"fmt"
	"net/http"
)

func main() {

	// instance of http server
	router := http.NewServeMux()

	
	// handler
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	//  intialise server port
	server := http.Server{
		Addr:    "8000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
