package rest

import (
	"cost-calculator/store"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	store store.Store
}

// NewHandler returns *Handler
func NewHandler(store store.Store) (*Handler, error) {
	return &Handler{
		store: store,
	}, nil
}

// Run runs the server.
func Run(store store.Store, address string) error {
	h, err := NewHandler(store)
	if err != nil {
		return err
	}

	http.HandleFunc("/category/insert", h.InsertCategory)
	http.HandleFunc("/category/get", h.GetCategory)
	http.HandleFunc("/category/update", h.UpdateCategory)
	http.HandleFunc("/category/delete", h.DeleteCategory)
	http.HandleFunc("/source/insert", h.InsertSource)
	http.HandleFunc("/source/get", h.GetSource)
	http.HandleFunc("/source/update", h.UpdateSource)
	http.HandleFunc("/source/delete", h.DeleteSource)
	http.HandleFunc("/incoming/insert", h.InsertIncome)
	http.HandleFunc("/incoming/get/one", h.GetIncome)
	http.HandleFunc("/incoming/get/", h.GetIncoming)
	http.HandleFunc("/incoming/update", h.UpdateIncome)
	http.HandleFunc("/incoming/delete", h.DeleteIncome)
	http.HandleFunc("/spend/insert", h.InsertSpend)
	http.HandleFunc("/spend/get/one", h.GetSpend)
	http.HandleFunc("/spend/get", h.GetSpends)
	http.HandleFunc("/spend/update", h.UpdateSpend)
	http.HandleFunc("/spend/delete", h.DeleteSpend)

	log.Println("REST server is running")
	return http.ListenAndServe(address, nil)
}

func response(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return
	}
}
