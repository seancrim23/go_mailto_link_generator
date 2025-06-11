package handler

import (
	"mailto_link_generator/services"
	"net/http"
)

type Mail struct {
	Service services.MailtoGeneratorService
}

func (m *Mail) GetMailToForm(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *Mail) CreateMailtoLink(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *Mail) GetMailtoLink(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}
