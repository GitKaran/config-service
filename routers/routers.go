package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/GitKaran-devops-test/database"
	"github.com/hellofreshdevtests/GitKaran-devops-test/handlers"
	"log"
	"net/http"
)

type application struct {
	Router *mux.Router
}

func NewApp() *application {
	router := mux.NewRouter()

	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	router.HandleFunc("/configs/{name}", api.GetConfig).Methods("GET")
	router.HandleFunc("/configs", api.GetAllConfigs).Methods("GET")
	router.HandleFunc("/configs", api.CreateConfig).Methods("POST")
	router.HandleFunc("/configs/{name}", api.UpdateConfig).Methods("PUT", "PATCH")
	router.HandleFunc("/configs/{name}", api.DeleteConfig).Methods("DELETE")
	router.HandleFunc("/search", api.SearchConfigByKey).Methods("GET")

	return &application{Router: router}
}

func (a *application) Start(port string) {
	fmt.Printf("Starting server on the port %v", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), a.Router))
}
