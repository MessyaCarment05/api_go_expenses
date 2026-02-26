package main

import (
	"apiexercises/database"
	"apiexercises/services"
	"log"
	"net/http"
)

func main() {
	log.Println("MAIN SUCCESSFULLY RUN")
	database.GetConnection()
	http.HandleFunc("/expenses/", func(w http.ResponseWriter, r *http.Request){
		log.Println("HIT /expenses", r.Method)
		switch r.Method {
		case http.MethodGet:
			services.GetExpenses(w,r)
		case http.MethodPost:
			services.CreateExpense(w,r)
		case http.MethodDelete:
			services.DeleteExpense(w,r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}