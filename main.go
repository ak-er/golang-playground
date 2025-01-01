package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct of data
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type SearchItemType struct {
	Name string `json:"name"`
}

// set dummy data
var items = []Item{
	{ID: 1, Name: "Pen"},
	{ID: 2, Name: "Copy"},
	{ID: 3, Name: "Mobile"},
	{ID: 4, Name: "Laptop"},
}

// helper function to respondError
func APIErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// get items handler
func GetItems(w http.ResponseWriter, r *http.Request) {
	// set the format content-type for http
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// create items handler
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	items = append(items, newItem)
	json.NewEncoder(w).Encode(items)
}

// update items handler
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateItem Item
	if len(items) == 0 {
		APIErrorResponse(w, http.StatusNotFound, "no items found")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, item := range items {
		if item.ID == updateItem.ID {
			items[i] = updateItem
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	APIErrorResponse(w, http.StatusBadRequest, "something went wrong")
}

// delete items handler
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var deleteItem Item
	if len(items) == 0 {
		APIErrorResponse(w, http.StatusNotFound, "items not found")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&deleteItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, item := range items {
		if item.ID == deleteItem.ID {
			items = append(items[:i], items[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	APIErrorResponse(w, http.StatusBadRequest, "something went wrong")
}

// search items handler
func SearchItem(w http.ResponseWriter, r *http.Request) {
	var searchItem SearchItemType
	if len(items) == 0 {
		APIErrorResponse(w, http.StatusBadRequest, "items not found")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&searchItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, item := range items {
		if item.Name == searchItem.Name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("SearchItem Not found:: %v", searchItem.Name))
}

const PORT = ":8080"

func main() {
	// setup routes
	http.HandleFunc("/items", GetItems)
	http.HandleFunc("/items/create", CreateItem)
	http.HandleFunc("/items/update", UpdateItem)
	http.HandleFunc("/items/delete", DeleteItem)
	http.HandleFunc("/items/search", SearchItem)
	// starting server
	fmt.Printf("Server is starting at http::/127.0.0.1:%s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Error Occur while server", err)
	}
}
