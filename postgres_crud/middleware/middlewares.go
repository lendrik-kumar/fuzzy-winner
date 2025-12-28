package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lendrik-kumar/postgres-crud/models"
	_ "github.com/lib/pq"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("warning: .env not loaded, relying on environment:", err)
	}

	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		log.Fatal("POSTGRES_URL is not set in environment")
	}

	if !strings.Contains(connStr, "sslmode=") {
		if strings.Contains(connStr, "?") {
			connStr = connStr + "&sslmode=disable"
		} else {
			connStr = connStr + "?sslmode=disable"
		}
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error loading the url, is it present in env")
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("db connected")

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertId := insertStock(stock)

	res := Response{
		ID:      insertId,
		Message: "Successfully created",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stock, err := getStock(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("stock updation sucessfull. total rows affescted %v", updatedRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stocks deleted sucessfully. %v", deletedRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(s models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlq := `INSERT INTO stocks(name, price, company) VALUES ($1,$2,$3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlq, s.Name, s.Price, s.Company).Scan(&id)

	if err != nil {
		log.Fatalf("unable to execute %v", err)
	}

	fmt.Printf("Inserted a record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock models.Stock

	sqlq := `SELECT * FROM stocks WHERE stockid=$1`

	err := db.QueryRow(sqlq, id).Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	if err != nil {
		return stock, err
	}

	return stock, nil
}

func getAllStocks() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlq := `SELECT * FROM stocks`

	rows, err := db.Query(sqlq)

	if err != nil {
		return stocks, err
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("unable to scan the row %v", err)
			return stocks, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func updateStock(id int64, s models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlq := `UPDATE stocks SET name=$1, price=$2, company=$3 WHERE stockid=$4`

	res, err := db.Exec(sqlq, s.Name, s.Price, s.Company, id)

	if err != nil {
		log.Fatalf("unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error while checking the affected rows %v", err)
	}

	fmt.Printf("total rows affected %v", rowsAffected)

	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlq := `DELETE FROM stocks WHERE stockid=$1`

	res, err := db.Exec(sqlq, id)

	if err != nil {
		log.Fatalf("unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error while checking the affected rows %v", err)
	}

	fmt.Printf("total rows affected %v", rowsAffected)

	return rowsAffected
}
