package main

// How to do versioning?
import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type Service struct {
	gorm.model
	Name string
}

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/services").
		HandlerFunc(servicesCreate)
}

func servicesCreate(w http.ResponseWriter, r *http.Request) {
}

func main() {
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	dbConnection := viper.GET("DATABASE_URL")

	db, err = gorm.Open("postgres", dbConnection)

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Service{})

	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)
	log.Fatal(http.listenAndServe(":8888", router))
}
