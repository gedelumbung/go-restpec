package main

import (
	"net/http"

	"github.com/gedelumbung/go-restpec"
)

func main() {

	http.HandleFunc("/200", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Ok(rw,
			"Well done!!",
			[]User{
				User{
					ID:   1,
					Name: "Gede",
				},
				User{
					ID:   2,
					Name: "Ayu",
				},
			},
			map[string]interface{}{
				"pagination": NewPagination(40, 1, 5),
			})
	})
	http.HandleFunc("/400", func(rw http.ResponseWriter, r *http.Request) {
		restpec.BadRequest(rw,
			"Your request is invalid",
			nil,
			nil)
	})
	http.HandleFunc("/401", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Unauthorized(rw,
			"Error code response for missing or invalid authentication token",
			nil,
			nil)
	})
	http.HandleFunc("/403", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Forbidden(rw,
			"You cannot access this route",
			nil,
			nil)
	})
	http.HandleFunc("/404", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Forbidden(rw,
			"Your request is not found",
			nil,
			nil)
	})
	http.HandleFunc("/405", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Forbidden(rw,
			"You method is not allowed",
			nil,
			nil)
	})
	http.HandleFunc("/422", func(rw http.ResponseWriter, r *http.Request) {
		restpec.Forbidden(rw,
			"We cannot process your entity",
			nil,
			nil)
	})

	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int
	Name string
}

type Pagination struct {
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

func NewPagination(total, page, limit int) *Pagination {
	lastPage := calculateLastPage(limit, total)
	return &Pagination{Total: total, Limit: limit, CurrentPage: page, LastPage: lastPage}
}

func calculateLastPage(perPage, total int) int {
	pages := total / perPage
	if total%perPage > 0 {
		return pages + 1
	}
	return pages
}
