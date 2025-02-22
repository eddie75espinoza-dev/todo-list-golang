package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Estructura para una tarea
type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var tasks []Task

// Obtener todas las tareas
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Crear una nueva tarea
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	// Tareas de ejemplo
	tasks = append(tasks, Task{ID: "1", Title: "Aprender Go", Details: "Estudiar conceptos básicos de Go."})
	tasks = append(tasks, Task{ID: "2", Title: "Construir una API", Details: "Crear una API REST en Go."})

	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Println("Servidor corriendo en http://localhost:8000 🚀")
	log.Fatal(http.ListenAndServe(":8000", router))
}
