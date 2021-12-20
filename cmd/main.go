package main

import (
	"cost-calculator/models"
	"cost-calculator/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	pg *store.Postgres
}

func jsonResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return
	}
}

func main() {
	cfg := store.NewConfig("postgres", "efim", "25121997", "costdb", 5432, false)
	pg, err := store.NewDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer pg.Close()

	h := Handler{pg: pg}

	http.HandleFunc("/incoming/insert", h.InsertIncoming)
	http.HandleFunc("/incoming/get", h.GetIncoming)
	http.HandleFunc("/incoming/update", h.UpdateIncoming)
	http.HandleFunc("/incoming/delete", h.DeleteIncoming)
	http.HandleFunc("/spend/insert", h.InsertSpends)
	http.HandleFunc("/spend/get", h.GetSpend)
	http.HandleFunc("/spend/update", h.UpdateSpend)
	http.HandleFunc("/spend/delete", h.DeleteSpend)
	log.Fatal(http.ListenAndServe("0.0.0.0:8085", nil))
}

func (h *Handler) InsertIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.pg.InsertIncoming(incoming)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", incoming.ID))
}

func (h *Handler) GetIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	incoming, err = h.pg.GetIncoming(incoming.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, incoming)
}

func (h *Handler) UpdateIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.pg.UpdateIncoming(incoming)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", incoming.ID))
}

func (h *Handler) DeleteIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.pg.DeleteIncoming(incoming.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", incoming.ID))
}

func (h *Handler) InsertSpends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.pg.InsertSpends(spend)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", spend.ID))
}

func (h *Handler) GetSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	spend, err = h.pg.GetSpend(spend.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, spend)
}

func (h *Handler) UpdateSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}
	err = h.pg.UpdateSpend(spend)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", spend.ID))
}

func (h *Handler) DeleteSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		jsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.pg.DeleteSpend(spend.ID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", spend.ID))
}
