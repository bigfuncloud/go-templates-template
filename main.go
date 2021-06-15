package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/bigfuncloud/@BFC_APP_DOMAIN@/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var t *template.Template

func main() {
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("PG_DSN"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Add more models in the models/ folder and here for auto-migration.
	if err := db.AutoMigrate(&models.Friend{}); err != nil {
		log.Fatalf("could not migrate: %v", err)
	}

	// Templates can be added in the templates/ folder.
	t, err = template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		log.Fatalf("could not load templates: %v", err)
	}

	r := mux.NewRouter()

	// Add more routes here.
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/", AddFriendHandler).Methods("POST")

	// This route can serve static files like CSS, images, etc.
	// For more flexibility, you may wish to use Caddy or another web server front.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}
