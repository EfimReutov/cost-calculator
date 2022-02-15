package rest

import (
	"cost-calculator/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) InsertCategory(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) InsertIncome(w http.ResponseWriter, r *http.Request) {
	income := new(models.Income)
	err := json.NewDecoder(r.Body).Decode(income)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	income.Date = time.Now().Local()

	err = h.store.InsertIncome(income)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful inserted, id: %d", income.ID))

}

func (h *Handler) GetIncome(w http.ResponseWriter, r *http.Request) {
	income := new(models.Income)
	err := json.NewDecoder(r.Body).Decode(income)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	income, err = h.store.GetIncome(income.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, income)
}

func (h *Handler) GetIncoming(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	q := r.URL.Query()
	page := q.Get("page")
	if page == "" {
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt == 0 {
		response(w, http.StatusBadRequest, "invalid parameter page")
		return
	}
	limit := q.Get("limit")
	if page == "" {
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt == 0 {
		response(w, http.StatusBadRequest, "invalid parameter limit")
		return
	}

	incoming, err := h.store.GetIncoming(pageInt, limitInt)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, incoming)
}

func (h *Handler) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusBadRequest, "invalid method")
		return
	}
	income := new(models.Income)
	err := json.NewDecoder(r.Body).Decode(income)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.UpdateIncome(income)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful updated, id: %d", income.ID))
}

func (h *Handler) DeleteIncome(w http.ResponseWriter, r *http.Request) {
	income := new(models.Income)
	err := json.NewDecoder(r.Body).Decode(income)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteIncome(income.ID)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, fmt.Sprintf("successful deleted, id: %d", income.ID))
}

func (h *Handler) InsertSpend(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetSpends(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page := q.Get("page")
	if page == "" {
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt == 0 {
		response(w, http.StatusBadRequest, "invalid parameter page")
		return
	}
	limit := q.Get("limit")
	if page == "" {
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt == 0 {
		response(w, http.StatusBadRequest, "invalid parameter limit")
		return
	}

	spends, err := h.store.GetSpends(pageInt, limitInt)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, spends)
}

func (h *Handler) UpdateSpend(w http.ResponseWriter, r *http.Request) {
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
