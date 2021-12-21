package handler

import (
	"cost-calculator/store"
)

type Handler struct {
	store store.Store
}

func NewHandler(store store.Store) (*Handler, error) {
	return &Handler{
		store: store,
	}, nil
}
