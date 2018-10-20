
# RestPec
Simple library to handle HTTP Status Code, because we need more respect to HTTP Status Code :p

## Example

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
