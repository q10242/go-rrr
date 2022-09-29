package controllers

import (
	"encoding/json"
	"go-rrr/pkg/models"
	"go-rrr/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
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
	b, err := CreateRedirect.CreateRedirect()
	if nil != err {
		res, _ := json.Marshal(map[string]string{
			"message": "This is not an regular url.",
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	} else {
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
