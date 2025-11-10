package models

import (
	"errors"
	"fmt"

	"github.com/karanjar/cargobackend_fibre_framework.git/config"
	"gorm.io/gorm"
)

type Car struct {
	Id    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Year  int64   `json:"year"`
	Price float64 `json:"price"`
}

type Error struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func (c *Car) Insert() error {

	if err := config.Db.Create(&c).Error; err != nil {
		fmt.Printf("error inserting car %v", err)
		return err
	}

	return nil
}
func (c *Car) Get() error {

	if err := config.Db.First(c, c.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("error getting car")
			return err
		}
	}

	return nil
}
func (c *Car) Update() error {

	if err := config.Db.Save(c).Error; err != nil {
		fmt.Printf("error updating car %v", err)
		return err
	}

	return nil
}
func (c *Car) Delete() error {

	if err := config.Db.Delete(c).Error; err != nil {
		fmt.Printf("error deleting car %v", err)
		return err
	}

	return nil
}
