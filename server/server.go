package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/eduardoacuna/self-esteem/tasks"
	"github.com/eduardoacuna/self-esteem/users"
)

var mux *goji.Mux

func init() {
	mux = goji.NewMux()
}

// SetupRoutes configures the multiplexer with the self-esteem API routes
func SetupRoutes() {
	mux.HandleFunc(pat.Get("/users"), getAllUsers)
	mux.HandleFunc(pat.Get("/users/:user-id"), getUserByID)
	mux.HandleFunc(pat.Get("/users/:user-id/tasks"), getAllTasks)
	mux.HandleFunc(pat.Get("/users/:user-id/tasks/:task-id"), getTaskByID)

	// Middleware
	mux.Use(logAllRequests)
}

// ListenAndServe run the server
func ListenAndServe() {
	fmt.Println(`

███████╗███████╗██╗     ███████╗    ███████╗███████╗████████╗███████╗███████╗███╗   ███╗
██╔════╝██╔════╝██║     ██╔════╝    ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔════╝████╗ ████║
███████╗█████╗  ██║     █████╗█████╗█████╗  ███████╗   ██║   █████╗  █████╗  ██╔████╔██║
╚════██║██╔══╝  ██║     ██╔══╝╚════╝██╔══╝  ╚════██║   ██║   ██╔══╝  ██╔══╝  ██║╚██╔╝██║
███████║███████╗███████╗██║         ███████╗███████║   ██║   ███████╗███████╗██║ ╚═╝ ██║
╚══════╝╚══════╝╚══════╝╚═╝         ╚══════╝╚══════╝   ╚═╝   ╚══════╝╚══════╝╚═╝     ╚═╝
                            (listening on port 8080)                                    `)
	http.ListenAndServe("localhost:8080", mux)
}

////////////////////
// users handlers //
////////////////////

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := users.GetAllUsers()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Fprintf(w, string(data))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user-id")
	user, err := users.GetUserByID(userID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Fprintf(w, string(data))
}

////////////////////
// tasks handlers //
////////////////////

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user-id")
	tasks, err := tasks.GetAllTasks(userID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Fprintf(w, string(data))
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user-id")
	taskID := pat.Param(r, "task-id")
	task, err := tasks.GetTaskByID(userID, taskID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}
	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Fprintf(w, string(data))
}

////////////////
// middleware //
////////////////

func logAllRequests(h http.Handler) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %v\n", r.URL)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(wrapper)
}
