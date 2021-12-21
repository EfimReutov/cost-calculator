package handler

import (
	"cost-calculator/api"
	"cost-calculator/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) InsertIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.InsertIncoming(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", incoming.ID))
}

func (h *Handler) GetIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	incoming, err = h.store.GetIncoming(incoming.ID)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, incoming)
}

func (h *Handler) UpdateIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.UpdateIncoming(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", incoming.ID))
}

func (h *Handler) DeleteIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteIncoming(incoming.ID)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", incoming.ID))
}

func (h *Handler) InsertSpends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.InsertSpends(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", spend.ID))
}

func (h *Handler) GetSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	spend, err = h.store.GetSpend(spend.ID)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, spend)
}

func (h *Handler) UpdateSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}
	err = h.store.UpdateSpend(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", spend.ID))
}

func (h *Handler) DeleteSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		api.JsonResponse(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		api.JsonResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteSpend(spend.ID)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.JsonResponse(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", spend.ID))
}
