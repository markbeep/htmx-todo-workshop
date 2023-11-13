package main

import (
	"coding-weekend/components"
	"coding-weekend/internal"
	"flag"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// command line flags
var (
	port = flag.String("port", "3000", "the port to serve the website on")
	host = flag.String("host", "localhost", "the host to serve the website on")
)

var (
	todoCounter = 0
	lock        sync.Mutex
)

func main() {
	flag.Parse() // parses command line flags defined above

	r := chi.NewRouter()
	r.Use(middleware.RequestID) // request IDs to distinguish different sessions
	r.Use(middleware.Logger)    // adds helpful logging

	todos := map[string][]internal.Todo{}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		requestId := extractRequestUser(middleware.GetReqID(r.Context()))
		userTodos, ok := todos[extractRequestUser(requestId)]
		if !ok {
			templ.Handler(components.Index([]internal.Todo{})).ServeHTTP(w, r)
		} else {
			templ.Handler(components.Index(userTodos)).ServeHTTP(w, r)
		}
	})

	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		requestId := extractRequestUser(middleware.GetReqID(r.Context()))

		r.ParseForm()
		text := r.PostFormValue("text")
		lock.Lock()
		defer lock.Unlock() // gets called even on panic
		todos[requestId] = append(todos[requestId], internal.Todo{Text: text, ID: todoCounter})
		todoCounter++
		templ.Handler(components.TodoList(todos[requestId])).ServeHTTP(w, r)
	})

	r.Delete("/todo/{id:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		requestId := extractRequestUser(middleware.GetReqID(r.Context()))

		// make sure the argument is a valid int (not too big)
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		todos[requestId] = slices.DeleteFunc[[]internal.Todo](todos[requestId], func(t internal.Todo) bool {
			return t.ID == id
		})
		templ.Handler(components.TodoList(todos[requestId])).ServeHTTP(w, r)
	})

	// starts listening
	fullHost := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("serving on %s", fullHost)
	if err := http.ListenAndServe(fullHost, r); err != nil {
		log.Fatal(err)
	}
}

func extractRequestUser(requestId string) string {
	// requestId is of format: Name/randomKey-00000x
	splitted := strings.Split(requestId, "-")
	return strings.Join(splitted[:len(splitted)-1], "")
}
