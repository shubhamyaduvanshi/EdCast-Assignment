package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// Database connection string
const dsn = "admin:admin12345@tcp(database-1.cldgjycdqfem.ap-south-1.rds.amazonaws.com)/Calculator"

/*type operands struct {
	Operand1 int `json:"operand1,omitempty"`
	Operand2 int `json:"operand2,omitempty"`
	Result   int `json:"response,omitempty"`
}*/

type Operands struct {
	Operand1 int `json:"operand1"`
	Operand2 int `json:"operand2"`
	Result   int `json:"response"`
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request Operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 + request.Operand2
	// Generate a random UUID for primary key
	id := uuid.New().String()

	// Insert the operands, method name, result, and generated UUID into the database
	err = insertOperation(id, request.Operand1, request.Operand2, response, "add")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request Operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 - request.Operand2
	// Generate a random UUID for primary key
        id := uuid.New().String()

        // Insert the operands, method name, result, and generated UUID into the database
        err = insertOperation(id, request.Operand1, request.Operand2, response, "subtract")
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
	json.NewEncoder(w).Encode(response)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request Operands
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "{}", http.StatusInternalServerError)
	}
	response := request.Operand1 * request.Operand2
        // Generate a random UUID for primary key
        id := uuid.New().String()

        // Insert the operands, method name, result, and generated UUID into the database
        err = insertOperation(id, request.Operand1, request.Operand2, response, "multiply")
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
	json.NewEncoder(w).Encode(response)
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

// Function to count the number of APIs hit so for---from DB
func countHandler(w http.ResponseWriter, r *http.Request) {
    // Open database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query to count the number of entries
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM `go-calc`").Scan(&count)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the count as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// Function to insert operation into the database
func insertOperation(id string, operand1, operand2, response int, methodName string) error {
	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	//Prepare the insert statement uuid, operand1, operand2, Method, Result id, operand1, operand2, method_name, result
	stmt, err := db.Prepare("INSERT INTO `go-calc` (uuid, operand1, operand2, Method, Result) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the insert statement
	_, err = stmt.Exec(id, operand1, operand2, methodName, response)
	if err != nil {
		return err
	}

	return nil
}



func main() {
	
	fmt.Println("Starting server on port 8080...")
        

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", status).Methods("GET")
        router.HandleFunc("/count", countHandler).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/subtract", subtract).Methods("POST")
	router.HandleFunc("/multiply", multiply).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
	
}
