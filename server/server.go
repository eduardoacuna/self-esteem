package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/satori/go.uuid"
	"goji.io"
	"goji.io/pat"

	"github.com/eduardoacuna/self-esteem/log"
	"github.com/eduardoacuna/self-esteem/tasks"
	"github.com/eduardoacuna/self-esteem/users"
)

var mux *goji.Mux

func init() {
	mux = goji.NewMux()
}

// SetupRoutes configures the multiplexer with the self-esteem API routes
func SetupRoutes(ctx context.Context) {
	log.Info(ctx, "routes setup")
	mux.HandleFunc(pat.Get("/users"), getAllUsers)
	mux.HandleFunc(pat.Get("/users/:user-id"), getUserByID)
	mux.HandleFunc(pat.Get("/users/:user-id/tasks"), getAllTasks)
	mux.HandleFunc(pat.Get("/users/:user-id/tasks/:task-id"), getTaskByID)

	// Middleware
	mux.Use(initializeContext)
	mux.Use(logAllRequests)
}

// ListenAndServe run the server
func ListenAndServe(ctx context.Context, bin, address, port string) {
	fmt.Println(`

███████╗███████╗██╗     ███████╗    ███████╗███████╗████████╗███████╗███████╗███╗   ███╗
██╔════╝██╔════╝██║     ██╔════╝    ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔════╝████╗ ████║
███████╗█████╗  ██║     █████╗█████╗█████╗  ███████╗   ██║   █████╗  █████╗  ██╔████╔██║
╚════██║██╔══╝  ██║     ██╔══╝╚════╝██╔══╝  ╚════██║   ██║   ██╔══╝  ██╔══╝  ██║╚██╔╝██║
███████║███████╗███████╗██║         ███████╗███████║   ██║   ███████╗███████╗██║ ╚═╝ ██║
╚══════╝╚══════╝╚══════╝╚═╝         ╚══════╝╚══════╝   ╚═╝   ╚══════╝╚══════╝╚═╝     ╚═╝`)
	log.Info(ctx, "server listening", "bin", bin, "address", address, "port", port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux)
}

////////////////////
// users handlers //
////////////////////

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := users.GetAllUsers(ctx)
	if err != nil {
		log.Error(ctx, "database problem", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "null")
		return
	}
	data, err := json.Marshal(users)
	if err != nil {
		log.Error(ctx, "json marshaling problem", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(data))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := strconv.Atoi(pat.Param(r, "user-id"))
	if err != nil {
		log.Error(ctx, "string to int problem", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "null")
		return
	}
	user, err := users.GetUserByID(ctx, userID)
	if err != nil {
		log.Error(ctx, "database problem", "error", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "null")
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Error(ctx, "json marshaling problem", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(data))
}

////////////////////
// tasks handlers //
////////////////////

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := strconv.Atoi(pat.Param(r, "user-id"))
	if err != nil {
		log.Error(ctx, "string to int problem", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "null")
		return
	}
	tasks, err := tasks.GetAllTasks(ctx, userID)
	if err != nil {
		log.Error(ctx, "database problem", "error", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "null")
		return
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		log.Error(ctx, "json marshaling problem", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(data))
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := strconv.Atoi(pat.Param(r, "user-id"))
	if err != nil {
		log.Error(ctx, "string to int problem", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "null")
		return
	}
	taskID, err := strconv.Atoi(pat.Param(r, "task-id"))
	if err != nil {
		log.Error(ctx, "string to int problem", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "null")
		return
	}
	task, err := tasks.GetTaskByID(ctx, userID, taskID)
	if err != nil {
		log.Error(ctx, "database problem", "error", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "null")
		return
	}
	data, err := json.Marshal(task)
	if err != nil {
		log.Error(ctx, "json marshaling problem", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(data))
}

////////////////
// middleware //
////////////////

func initializeContext(h http.Handler) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "ID", uuid.NewV4().String())
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(wrapper)
}

func logAllRequests(h http.Handler) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Info(ctx, fmt.Sprintf("%s %v", r.Method, r.URL))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(wrapper)
}
