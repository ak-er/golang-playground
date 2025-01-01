package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// struct of data
type Item struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB
var items []Item

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatalf("Error DB setup:: %v", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatalf("Error Creating table:: %v", err)
	}
}

// helper function to respondError
func APIErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// get items handler
func GetItems(w http.ResponseWriter, r *http.Request) {
	row, err := db.Query(`SELECT * FROM items`)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Database Error-getitems:: %v", err.Error()))
		return
	}
	defer row.Close()
	var response []Item
	for row.Next() {
		var item Item
		err := row.Scan(&item.ID, &item.Name)
		if err != nil {
			APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Database Error-scanitem:: %v", err.Error()))
			return
		}
		response = append(response, item)
	}
	if response == nil {
		APIErrorResponse(w, http.StatusNotFound, "items not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// create items handler
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("JSONDecodeError:: %v", err.Error()))
		return
	}
	stmt, err := db.Prepare(`INSERT INTO items (name) VALUES(?)`)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("DatabaseErrorCreateItems:: %v", err.Error()))
		return
	}
	defer stmt.Close()
	resp, err := stmt.Exec(newItem.Name)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("ExecuteQueryCreateItem:: %v", err.Error()))
		return
	}
	insertedId, err := resp.LastInsertId()
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("FailedToGetLastInsertedId:: %v", err.Error()))
		return
	}
	newItem.ID = insertedId
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

// update items handler
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateItem Item
	err := json.NewDecoder(r.Body).Decode(&updateItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	stmt, err := db.Prepare(`UPDATE items SET name=(?) WHERE id=(?)`)
	if err != nil {
		APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("FailedToUpdateQuery:: %v", err.Error()))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(updateItem.Name, updateItem.ID)
	if err != nil {
		APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("FailedToExecuteUpdateQuery:: %v", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateItem)
}

// delete items handler
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var deleteItem Item
	err := json.NewDecoder(r.Body).Decode(&deleteItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	stmt, err := db.Prepare(`DELETE FROM items WHERE id=(?)`)
	if err != nil {
		APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("FailedToDeleteQuery:: %v", err.Error()))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(deleteItem.ID)
	if err != nil {
		APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("FailedToExecuteDeleteQuery:: %v", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	respData := map[string]interface{}{
		"message":    "Successfully Deleted",
		"deleted_id": deleteItem.ID,
	}
	json.NewEncoder(w).Encode(respData)
}

// search items handler
func SearchItem(w http.ResponseWriter, r *http.Request) {
	var searchItem Item
	err := json.NewDecoder(r.Body).Decode(&searchItem)
	if err != nil {
		APIErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	row, err := db.Query(`SELECT * FROM items WHERE name LIKE ?`, "%"+searchItem.Name+"%")
	if err != nil {
		APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("FailedToSearchQuery:: %v", err.Error()))
		return
	}
	defer row.Close()
	var response []Item
	for row.Next() {
		var item Item
		err := row.Scan(&item.ID, &item.Name)
		if err != nil {
			APIErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Database Error-searchscanitem:: %v", err.Error()))
			return
		}
		response = append(response, item)
	}
	if response == nil {
		APIErrorResponse(w, http.StatusNotFound, "items not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
