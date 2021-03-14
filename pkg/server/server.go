package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/amrali/FOODOFFERS/pkg/db"
	"github.com/amrali/FOODOFFERS/pkg/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	name := ""
	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}
	if err := tmpl.ExecuteTemplate(w, name, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func GetActiveRestaurantOffers(w http.ResponseWriter, r *http.Request) {
	database := db.ConnectDB()
	defer database.Close()
	var validRestaurantOffers []models.RestaurantOffers
	database.Where("valid = ?", "Y").Find(&validRestaurantOffers)
	json.NewEncoder(w).Encode(validRestaurantOffers)
}
func GetActiveSiteOffers(w http.ResponseWriter, r *http.Request) {
	database := db.ConnectDB()
	defer database.Close()
	validSiteOffers := models.SiteOffers{}
	database.Where("valid = ?", "Y").First(&validSiteOffers)
	json.NewEncoder(w).Encode(validSiteOffers)
}
