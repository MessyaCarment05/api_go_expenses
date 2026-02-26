package models



type Expense struct {
	ExpensesID int    `json:"expenses_id"`
	Nominal    int    `json:"nominal"`
	Category   string `json:"category"`
	Note       string `json:"note"`
	Date       string `json:"date"`
}