package main

import (
	"log"
	"net/http"
)

type ControlPlane struct {
	Templates struct {
		GetToken Template
		Token    Template
	}
	db *DbService
}

const BytesPerToken = 30

func (ctrl *ControlPlane) Home(w http.ResponseWriter, r *http.Request) {
	ctrl.Templates.GetToken.Execute(w, r, nil)
}

func (ctrl *ControlPlane) Token(w http.ResponseWriter, r *http.Request) {

	token, err := String(BytesPerToken)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = ctrl.db.InsertApiToken(Hash(token))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var data struct {
		Token string
	}
	data.Token = token
	ctrl.Templates.Token.Execute(w, r, data)
}
