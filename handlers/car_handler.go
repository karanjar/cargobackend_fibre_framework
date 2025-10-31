package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/karanjar/cargobackend_fibre_framework.git/models"
)

var Mu sync.Mutex

func Carhandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			Createcar(w, r)
		} else {
			http.Error(w, "Incorrect Bad Request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w, "You entered bad get request", http.StatusBadRequest)
			return
		} else {
			id, _ := strconv.Atoi(entity)
			Getcar(w, r, id)

		}
	case "PUT":
		if entity == "" {
			http.Error(w, "You entered bad put request", http.StatusBadRequest)

		} else {
			id, _ := strconv.Atoi(entity)
			Updatecar(w, r, id)
		}

	case "DELETE":
		if entity == "" {
			http.Error(w, "You entered bad delete request", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			Deletecar(w, r, id)
		}
	default:
		http.Error(w, "we dont support this Method", http.StatusMethodNotAllowed)
	}

}
func Createcar(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}

	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
		return
	}

	if err := car.Insert(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Car created with the id:", car.Id)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}

func Getcar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{Id: id}

	if err := car.Get(); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Car not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Server error", http.StatusInternalServerError)
		fmt.Println("Error fetching car:", err)
		return
	}

	fmt.Println("Car found with the id:", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)

}

func Deletecar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()
	car := &models.Car{Id: id}

	car.Delete()

	fmt.Println("Car deleted with the id:", id)

	w.WriteHeader(http.StatusNoContent)

}
func Updatecar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()
	car := &models.Car{Id: id}
	if err := json.NewDecoder(r.Body).Decode(car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
		return
	}
	if err := car.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Car Updated with the id:", car.Id)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}
