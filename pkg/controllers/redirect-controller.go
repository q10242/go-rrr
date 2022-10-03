package controllers

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateRedirect(w http.ResponseWriter, r *http.Request) {
	CreateRedirect := &models.Redirect{}
	utils.ParseBody(r, CreateRedirect)
	fmt.Println(CreateRedirect.Probability)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	b, err := CreateRedirect.CreateRedirect()
	if nil != err {
		res, _ := json.Marshal(map[string]string{
			"message": "This is not an regular url.",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		res, _ := json.Marshal(b)

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
