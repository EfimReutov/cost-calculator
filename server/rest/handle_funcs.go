package rest

import (
	"cost-calculator/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) InsertCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	category := new(models.Category)
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.InsertCategory(category)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", category.ID))
}

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	category := new(models.Category)
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	category, err = h.store.GetCategory(category.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, category)
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	category := new(models.Category)
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}
	err = h.store.UpdateCategory(category)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", category.ID))
}

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	category := new(models.Category)
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteCategory(category.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", category.ID))
}

func (h *Handler) InsertSource(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	source := new(models.Source)
	err := json.NewDecoder(r.Body).Decode(source)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.InsertSource(source)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", source.ID))
}

func (h *Handler) GetSource(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	source := new(models.Source)
	err := json.NewDecoder(r.Body).Decode(source)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	source, err = h.store.GetSource(source.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, source)
}

func (h *Handler) UpdateSource(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	source := new(models.Source)
	err := json.NewDecoder(r.Body).Decode(source)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}
	err = h.store.UpdateSource(source)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", source.ID))
}

func (h *Handler) DeleteSource(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	source := new(models.Source)
	err := json.NewDecoder(r.Body).Decode(source)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteSpend(source.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", source.ID))
}

func (h *Handler) InsertIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	incoming.Date = time.Now().Local()

	err = h.store.InsertIncoming(incoming)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", incoming.ID))

}

func (h *Handler) GetIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	incoming, err = h.store.GetIncoming(incoming.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, incoming)
}

func (h *Handler) UpdateIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.UpdateIncoming(incoming)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", incoming.ID))
}

func (h *Handler) DeleteIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	incoming := new(models.Incoming)
	err := json.NewDecoder(r.Body).Decode(incoming)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteIncoming(incoming.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", incoming.ID))
}

func (h *Handler) InsertSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.InsertSpend(spend)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", spend.ID))
}

func (h *Handler) GetSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	spend, err = h.store.GetSpend(spend.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, spend)
}

func (h *Handler) UpdateSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}
	err = h.store.UpdateSpend(spend)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", spend.ID))
}

func (h *Handler) DeleteSpend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	spend := new(models.Spend)
	err := json.NewDecoder(r.Body).Decode(spend)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteSpend(spend.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", spend.ID))
}
