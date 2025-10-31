package models

import (
	"database/sql"
	"fmt"

	"github.com/karanjar/cargobackend_fibre_framework.git/config"
)

type Car struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Year  int64   `json:"year"`
	Price float64 `json:"price"`
}

func (c *Car) Insert() error {
	query := `INSERT INTO cars (name,model,year,price) VALUES ($1,$2,$3,$4) RETURNING id`
	err := config.Db.QueryRow(query, c.Name, c.Model, c.Year, c.Price).Scan(&c.Id)
	if err != nil {
		return fmt.Errorf("error inserting car to the database ", err)
	}
	return nil
}
func (c *Car) Get() error {
	query := `SELECT name,model,year,price FROM cars WHERE id = $1`
	err := config.Db.QueryRow(query, c.Id).Scan(&c.Name, &c.Model, &c.Year, &c.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no car found with id %d", c.Id)
		}
		return fmt.Errorf("error fetching car: %v", err)
	}
	return nil
}
func (c *Car) Update() error {
	query := `UPDATE cars SET name = $1, model = $2, year = $3, price = $4 WHERE id = $5`

	_, err := config.Db.Exec(query, c.Name, c.Model, c.Year, c.Price, c.Id)
	if err != nil {
		return fmt.Errorf("Error updating car: %v", err)
	}
	return nil
}
func (c *Car) Delete() {
	query := `DELETE FROM cars WHERE id = $1`
	_, err := config.Db.Exec(query, c.Id)
	if err != nil {
		fmt.Printf("Error deleting car with id: %v, error %v", c.Id, err)
	}
}
