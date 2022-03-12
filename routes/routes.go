package routes

import (
	"log"
	"net/http"

	controllers "api/Controllers"
	"api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/personalidades", controllers.GetAllPersonalities).Methods("Get")
	router.HandleFunc("/personalidades/{id}", controllers.GetPersonalityById).Methods("Get")
	router.HandleFunc("/personalidades", controllers.CreatePersonality).Methods("Post")
	router.HandleFunc("/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc("/personalidades/{id}", controllers.EditPersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))

}
