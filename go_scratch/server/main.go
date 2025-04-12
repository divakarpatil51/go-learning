package main

import (
	"log"
	"net/http"
)

type app struct {
	addr string
}

func (a *app) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Post call")
	w.Write([]byte("test"))
}

func (a *app) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Write call")
	w.Write([]byte("test"))
	w.WriteHeader(http.StatusOK)
}
func main() {
	app := &app{addr: ":8080"}

	mux := http.NewServeMux()

	serv := &http.Server{Addr: app.addr, Handler: mux}

	mux.HandleFunc("GET /users", app.getUserHandler)
	mux.HandleFunc("POST /users", app.createUserHandler)

	err := serv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
