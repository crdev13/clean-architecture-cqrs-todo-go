package http

import "net/http"

func NewServer(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("POST /users/", func(w http.ResponseWriter, r *http.Request) {
		// handles POST /users/{id}/todos
		if r.Method == "POST" && len(r.URL.Path) > 7 {
			h.CreateTodo(w, r)
			return
		}
		w.WriteHeader(404)
	})

	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request) {
		// handles GET /users/{id}/todos
		if r.Method == "GET" && len(r.URL.Path) > 7 {
			h.ListUserTodos(w, r)
			return
		}
		w.WriteHeader(404)
	})
	mux.HandleFunc("POST /todos/", func(w http.ResponseWriter, r *http.Request) {
		// POST /todos/{id}/done
		h.MarkTodoDone(w, r)
	})

	return mux
}
