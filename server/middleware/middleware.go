package middleware

import(
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"go.mangodb.org/mongo-driver/bson"
	"go.mangodb.org/mongo-driver/bson/primitive"
	"go.mangodb.org/mongo-driver/mongo"
	"go.mangodb.org/mongo-driver/mongo/options"
	""
)
var collection *mongo.Collection

func init(){
	loadTheEnv()
	creatDbInstance()
}

func loadTheEnv(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading the .env file")
	}
}

func creatDbInstance(){
	connectionString := os.Getenv(key)("DB_URI")
	dbname := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err!=nill{
		log.Fatal(err)
	}

	err := client.Ping(context.TODO(),nil)
	if err!=nill{
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb!")

	client.Database(dbName).Collection(collName)
	fmt.Println("collection intance created")
}

func GetAllTasks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	payload := GetAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")
	w.Header().Set("Acess-Control-Allow-Headers", "Context-Type")
	var task models.ToDoList 
	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")
	w.Header().Set("Acess-Control-Allow-Headers", "Context-Type")

	params := mux.Vars(r)
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params['id'])
}

func UndoTask(w http.ResponseWriter,r *http.Request){

}

func DeleteTask(w http.ResponseWriter,r *http.Request){

}

func DeleteAllTask(w http.ResponseWriter,r *http.Request){

}

func GetAllTasks(){

}

func TaskComplete(){

}