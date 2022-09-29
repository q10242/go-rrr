package controllers

import (
	"encoding/json"

	"github.com/gorilla/mux"

	"net/http"

	"go-rrr/pkg/models"
	"go-rrr/pkg/utils"
)

func GetRedirectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ShortUrl := vars["ShortUrl"]
	redirectDetails, _ := models.GetRedirectByShortUrl(ShortUrl)
	res, _ := json.Marshal(redirectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateRedirect(w http.ResponseWriter, r *http.Request) {
	CreateRedirect := &models.Redirect{}
	utils.ParseBody(r, CreateRedirect)
	b := CreateRedirect.CreateRedirect()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
