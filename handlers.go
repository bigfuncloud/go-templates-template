package main

import (
	"fmt"
	"net/http"

	"github.com/bigfuncloud/@BFC_APP_DOMAIN@/models"
)

type HomeData struct {
	Friends []*models.Friend
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var d HomeData
	result := db.Find(&d.Friends)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("could not get friends: %v", result.Error), http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "index", d); err != nil {
		http.Error(w, fmt.Sprintf("could not execute template: %v", err), http.StatusInternalServerError)
	}
}

func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	result := db.Create(&models.Friend{Name: name, Email: email})

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("could not create friend: %v", result.Error), http.StatusInternalServerError)
		return
	}

	var d HomeData
	result = db.Find(&d.Friends)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf("could not get friends: %v", result.Error), http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "index", d); err != nil {
		http.Error(w, fmt.Sprintf("could not execute template: %v", err), http.StatusInternalServerError)
	}
}
