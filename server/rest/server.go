package rest

import (
	"cost-calculator/config"
	"cost-calculator/store"
	"encoding/json"
	"fmt"
	"github.com/EfimReutov/mail_sender"
	"log"
	"net/http"
)

type Handler struct {
	store  store.StorePostgres
	sender *mail_sender.Sender
}

// NewHandler returns *Handler
func NewHandler(store store.StorePostgres, cfg *config.Configuration) (*Handler, error) {
	return &Handler{
		store: store,
		sender: mail_sender.NewSender(
			mail_sender.Configuration{
				SMTPServer:   cfg.SMTPServer,
				SMTPPort:     cfg.SMTPPort,
				MailUser:     cfg.MailUser,
				MailPassword: cfg.MailPassword,
			},
		),
	}, nil
}

// Run runs the server.
func Run(store store.StorePostgres, cfg *config.Configuration) error {
	h, err := NewHandler(store, cfg)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("/otp/get", methods(h.OTPSend, http.MethodGet))
	mux.Handle("/auth", methods(h.authorization, http.MethodGet))

	mux.Handle("/category/insert", access(methods(h.InsertCategory, http.MethodPost)))
	mux.Handle("/category/get", access(methods(h.GetCategory, http.MethodGet)))
	mux.Handle("/category/update", access(methods(h.UpdateCategory, http.MethodPut)))
	mux.Handle("/category/delete", access(methods(h.DeleteCategory, http.MethodDelete)))

	mux.Handle("/source/insert", access(methods(h.InsertSource, http.MethodPost)))
	mux.Handle("/source/get", access(methods(h.GetSource, http.MethodGet)))
	mux.Handle("/source/update", access(methods(h.UpdateSource, http.MethodPut)))
	mux.Handle("/source/delete", access(methods(h.DeleteSource, http.MethodDelete)))

	mux.Handle("/income/insert", access(methods(h.InsertIncome, http.MethodPost)))
	mux.Handle("/income/get", access(methods(h.GetIncome, http.MethodGet)))
	mux.Handle("/income/update", access(methods(h.UpdateIncome, http.MethodPut)))
	mux.Handle("/income/delete", access(methods(h.DeleteIncome, http.MethodDelete)))
	mux.Handle("/incoming/get", access(methods(h.GetIncoming, http.MethodGet)))

	mux.Handle("/spend/insert", access(methods(h.InsertSpend, http.MethodPost)))
	mux.Handle("/spend/get", access(methods(h.GetSpend, http.MethodGet)))
	mux.Handle("/spend/update", access(methods(h.UpdateSpend, http.MethodPut)))
	mux.Handle("/spend/delete", access(methods(h.DeleteSpend, http.MethodDelete)))

	handler := timing(mux)

	log.Println("REST server is running")
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.ServiceHost, cfg.ServicePort), handler)
}

func response(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return
	}
}
