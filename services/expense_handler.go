package services

import (
	"apiexercises/database"
	"apiexercises/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	script := "SELECT expenses_id, nominal, category, note, date FROM expenses"
	rows, err := database.DB.QueryContext(ctx, script)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var e models.Expense
		rows.Scan(&e.ExpensesID, &e.Nominal, &e.Category, &e.Note, &e.Date)
		expenses = append(expenses, e)
	}

	json.NewEncoder(w).Encode(expenses)
	log.Println("GET SUCCESS")
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	// decode request JSON ke http dulu
	var temp models.Expense
	reqErr := json.NewDecoder(r.Body).Decode(&temp)
	if reqErr != nil {
		http.Error(w, "invalid body", 400)
		return
	}

	ctx := context.Background()
	script := "INSERT INTO expenses (nominal, category, note,date) VALUES(?,?,?,?)"

	_, err := database.DB.ExecContext(ctx, script, temp.Nominal, temp.Category, temp.Note, temp.Date)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("POST SUCCESS")
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	// ambil id dari request
	id_input := strings.TrimSpace(r.URL.Query().Get("id"))

	id, err := strconv.Atoi(id_input)

	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}
	ctx := context.Background()
	script := "DELETE FROM expenses WHERE expenses_id = ?"
	_, err2 := database.DB.ExecContext(ctx, script, id)
	if err2 != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	log.Println("DELETE SUCCESS")
}
