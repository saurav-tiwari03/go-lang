package middleware

import (
	"os"
	"net/http"
	"fmt"
	"context"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request)  {
	
}

func CreateTask()  {
	
}

func TaskComplete()  {
	
}
func UndoTask()  {
	
}
func DeleteTask()  {
	
}
func DeleteAllTask()  {
	
}