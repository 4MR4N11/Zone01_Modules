package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/api"
	"forum/config"
	"forum/handlers"
	"forum/utils"
)

func main() {
	if err := utils.InitServices(); err != nil {
		log.Fatal(err)
	}
	if err := utils.InitTables(); err != nil {
		log.Fatal(err)
	}
	defer config.DB.Close()
	http.HandleFunc("/static/", handlers.ServeStatic)
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/api/register", api.RegisterApi)
	http.HandleFunc("/api/post", api.PostApi)
	http.HandleFunc("/api/react", api.ReactToPostHandler)
	http.HandleFunc("/api/add/comment", api.AddComment)
	http.HandleFunc("/api/like/comment", api.HandleLikeComment)
	http.HandleFunc("/api/login", api.LoginApi)
	http.HandleFunc("/api/filter", api.FilterApi)
	http.HandleFunc("/api/post/{id}", api.GetPostApi)
	http.HandleFunc("/api/logout", api.LogoutApi)
	http.HandleFunc("/ws", api.WsEndpoint)
	http.HandleFunc("/api/msg", api.MessageApi)
	fmt.Printf("Server running on http://localhost%v", config.ADDRS)
	err := http.ListenAndServe(config.ADDRS, nil)
	if err != nil {
		log.Fatal(err)
	}
}
