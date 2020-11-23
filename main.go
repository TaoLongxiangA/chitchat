package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("ChitChat", version(), "started at", config.Address)

	// handle static file
	//mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// all route patterns matched here

	// index
	http.HandleFunc("/", index)
	// error
	http.HandleFunc("/err", err)

	// defined in route_auth.go
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signup_account", signupAccount)
	http.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	http.HandleFunc("/thread/new", newThread)
	http.HandleFunc("/thread/create", createThread)
	http.HandleFunc("/thread/post", postThread)
	http.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        nil,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}
