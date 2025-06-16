package handler

import (
	"errors"
	"fmt"
	"html/template"
	"mailto_link_generator/services"
	"mailto_link_generator/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Mail struct {
	Service services.MailtoGeneratorService
}

func (m *Mail) GetMailToForm(w http.ResponseWriter, r *http.Request) {
	//TODO update to pull template from cloud file storage
	tmpl := template.Must(template.ParseFiles("C:/development/go_mailto_link_generator/templates/mailto.html"))
	fmt.Println("getting the mailto for the home page...")
	tmpl.Execute(w, nil)
}

func (m *Mail) CreateMailtoLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var err error
	var to string
	var subject string
	var body string

	tmpl := template.Must(template.ParseFiles("C:/development/go_mailto_link_generator/templates/mailto.html"))
	fmt.Println("getting the mailto for the home page...")

	fmt.Println("creating short url...")
	to = r.FormValue("to")
	subject = r.FormValue("subject")
	body = r.FormValue("body")

	if to == "" || subject == "" || body == "" {
		tmpl.Execute(w, map[string]string{"Error": "invalid form entries", "Code": "400"})
	}

	mailToLink, err := m.Service.CreateMailtoLink(r.Context(), to, subject, body)
	//determine what type of error and change code and return according error message
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": err.Error(), "Code": "500"})
	}

	tmpl.Execute(w, map[string]string{"MailToURL": mailToLink, "Generated": "true"})
}

func (m *Mail) GetMailtoLink(w http.ResponseWriter, r *http.Request) {
	var code = 200
	var err error

	shortLink := chi.URLParam(r, "id")
	if shortLink == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no shortLink passed to request").Error())
		return
	}

	//TODO get a check somewhere in here eventually to see if the url is in the cache

	mailToLink, err := m.Service.GetMailtoLink(r.Context(), shortLink)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}
	if mailToLink == "" {
		code = 404
		utils.RespondWithError(w, code, errors.New("cannot find mailto link").Error())
		return
	}

	fmt.Println(mailToLink)
	http.Redirect(w, r, mailToLink, 302)
	//should the backend redirect or should the front end do the redirect?
	utils.RespondWithJSON(w, code, mailToLink)
}
