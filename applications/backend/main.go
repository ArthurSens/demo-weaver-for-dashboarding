package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpconv "github.com/ArthurSens/demo-weaver-for-dashboarding/generated/client/go/http"
)

type Book struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Location   string    `json:"location"`
	BorrowedAt time.Time `json:"borrowed_at,omitzero"`
	ReturnedAt time.Time `json:"returned_at,omitzero"`
}

func main() {
	books := map[string]*Book{
		"dcc": {
			ID:         "dcc",
			Title:      "Dungeon Crawler Carl",
			Author:     "Matt Dinniman",
			Location:   "Vienna/01",
			ReturnedAt: time.Date(2025, 8, 13, 17, 42, 0, 0, time.UTC),
		},
		"o11y": {
			ID:         "o11y",
			Title:      "Using weaver for Observability",
			Author:     "PromCon Participants",
			Location:   "Munich/01",
			ReturnedAt: time.Date(2025, 4, 16, 17, 42, 0, 0, time.UTC),
		},
	}

	requestDurations := httpconv.NewServerRequestDuration()
	prometheus.Register(requestDurations)
	recordRequestDuration := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			requestDurations.With(httpconv.RequestMethodAttr(r.Method), httpconv.UrlSchemeAttr("http"), httpconv.RouteAttr(r.Pattern)).Observe(float64(time.Since(start).Seconds()))
		})
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", HandleGetBooks(books))
	mux.HandleFunc("POST /books/{id}/borrow", HandlePostBorrow(books))
	mux.HandleFunc("POST /books/{id}/return", HandlePostReturn(books))
	mux.HandleFunc("POST /books/", HandlePostBooks(books))
	mux.Handle("GET /metrics", promhttp.Handler())
	slog.Info("starting backend server", "addr", ":8080")
	http.ListenAndServe(":8080", recordRequestDuration(mux))
}

func HandleGetBooks(books map[string]*Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookList := make([]*Book, len(books))
		i := 0
		for _, b := range books {
			bookList[i] = b
			i++
		}
		json.NewEncoder(w).Encode(bookList)

	}
}

func HandlePostBorrow(books map[string]*Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, ok := books[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if !books[id].BorrowedAt.IsZero() {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		books[id].BorrowedAt = time.Now()
		w.WriteHeader(http.StatusAccepted)
	}
}
func HandlePostReturn(books map[string]*Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, ok := books[id]
		if !ok {
		}
		if books[id].BorrowedAt.IsZero() {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		books[id].BorrowedAt = time.Time{}
		books[id].ReturnedAt = time.Now()
		w.WriteHeader(http.StatusAccepted)
	}
}

func HandlePostBooks(books map[string]*Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			slog.Error("failed to decode request body", "err", err)
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}
		books[book.ID] = &book
		w.WriteHeader(http.StatusCreated)
	}
}
