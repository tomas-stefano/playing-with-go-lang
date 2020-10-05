package main

// How to do versioning?
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type Service struct {
	gorm.Model
	Name string
}

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/services").
		HandlerFunc(createService)

	router.
		Methods("GET").
		Path("/services/{id}").
		HandlerFunc(getService)
}

func createService(w http.ResponseWriter, r *http.Request) {
	var service Service
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the name")
	}
	if err := json.Unmarshal(reqBody, &service); err != nil {
		panic(err)
	}
	db.Create(&service)

	json.NewEncoder(w).Encode(service)
}

func getService(w http.ResponseWriter, r *http.Request) {
	var service Service
	serviceID := mux.Vars(r)["id"]
	db.First(&service, serviceID)

	json.NewEncoder(w).Encode(service)
}

func main() {
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	//	dbConnection := viper.Get("DATABASE_URL")

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password=password dbname=service_api_local host=service-api-db sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Service{})

	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":8888", router))
}
