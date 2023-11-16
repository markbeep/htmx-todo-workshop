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

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// command line flags
var (
	port = flag.String("port", "3000", "the port to serve the website on")
	host = flag.String("host", "0.0.0.0", "the host to serve the website on")
)

var (
	todos       = []internal.Todo{} // all todos are stored in a global array
	todoCounter = 0                 // used for assigning todo ids
)

func main() {
	flag.Parse() // parses command line flags defined above

	r := chi.NewRouter()
	r.Use(middleware.Logger) // adds helpful logging

	// Handles all GET requests to /
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(components.Index(todos)).ServeHTTP(w, r)
	})

	// Handles all POST requests to /todo made by htmx
	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() // REQUIRED for the r.PostFormValue to get filled with values
		text := r.PostFormValue("text")

		todos = append(todos, internal.Todo{Text: text, ID: todoCounter})
		todoCounter++

		templ.Handler(components.TodoList(todos)).ServeHTTP(w, r)
	})

	// Handles all DELETE requests to /todo/{id} where id is some number
	r.Delete("/todo/{id:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		// make sure the argument is a valid int (not too large for example)
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		todos = slices.DeleteFunc[[]internal.Todo](todos, func(t internal.Todo) bool {
			return t.ID == id
		})
		templ.Handler(components.TodoList(todos)).ServeHTTP(w, r)
	})

	// starts listening
	fullHost := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("serving on %s", fullHost)
	if err := http.ListenAndServe(fullHost, r); err != nil {
		log.Fatal(err)
	}
}
