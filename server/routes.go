package server

import (
	"mailto_link_generator/handler"
	"mailto_link_generator/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (m *MailtoGeneratorServer) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/mailto", m.loadMailtoRoutes)

	m.router = router
}

func (m *MailtoGeneratorServer) loadMailtoRoutes(router chi.Router) {
	mailToHandler := &handler.Mail{
		Service: &services.FirestoreMailtoGeneratorService{
			Database: m.firestoreDb,
		},
	}

	router.Get("/", mailToHandler.GetMailToForm)
	router.Post("/", mailToHandler.CreateMailtoLink)
	router.Get("/{id}", mailToHandler.GetMailtoLink)
}
