package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/readfile", readFileHandler)
	http.ListenAndServe(":8080", nil) // INSECURE: no TLS
}

// Vulnerable login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// INSECURE: Hardcoded credentials
	username := "admin"
	passwordHash := md5.Sum([]byte("123456")) // INSECURE: weak hashing

	// INSECURE: direct SQL concatenation (SQL injection)
	userInput := r.FormValue("user")
	passInput := r.FormValue("pass")
	query := fmt.Sprintf("SELECT id FROM users WHERE user='%s' AND pass='%s'", userInput, passInput)

	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1)/demo") // INSECURE: hardcoded credentials
	defer db.Close()

	_, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(w, "Login failed")
		return
	}

	fmt.Fprintf(w, "Login OK for %s", username)
}

// File read handler
func readFileHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	// INSECURE: No validation, can read arbitrary system files
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(w, "Could not read: %v", err)
		return
	}

	fmt.Fprintf(w, "Content:\n%s", content)

	// INSECURE: Logging sensitive data
	log.Printf("User read file: %s", file)
}

// INSECURE: No error checking, wrong permissions, prints all environment variables
func dumpEnv() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}