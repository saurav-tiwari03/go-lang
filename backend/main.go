package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"context"
	"time"
)

type Todo struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Completed   bool   `json:"completed" bson:"completed"`
}

var (
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
)

func init() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// MongoDB connection
	uri := os.Getenv("DB_URI")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("MongoDB not reachable: %v", err)
	}

	db = client.Database("todo_db")
	log.Println("Connected to MongoDB")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	r := mux.NewRouter()
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	log.Printf("Server is running on port %s", port)
	http.ListenAndServe(":"+port, r)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := db.Collection("todos")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var todos []Todo
	if err := cursor.All(ctx, &todos); err != nil {
		http.Error(w, "Failed to decode todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest);