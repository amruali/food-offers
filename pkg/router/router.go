package router

import (
	"net/http"

	"github.com/amrali/FOODOFFERS/pkg/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HttpRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))
	router.HandleFunc("/", server.Index).Methods("GET")
	router.HandleFunc("/api/valid_restaurant_offers", server.GetActiveRestaurantOffers).Methods("GET")
	router.HandleFunc("/api/valid_site_offers", server.GetActiveSiteOffers).Methods("GET")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "HEAD", "OPTIONS", "PUT"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Access-Control-Allow-Origin"},
		OptionsPassthrough: true,
	})
	handler := corsOpts.Handler(router)
	http.ListenAndServe(":8080", handler)
}
