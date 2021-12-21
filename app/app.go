package app

import (
	"cost-calculator/config"
	"cost-calculator/handler"
	"cost-calculator/store"
	"fmt"
	"net/http"
)

func Run() error {
	cfg, err := config.LoadCfg()
	if err != nil {
		return err
	}

	pg, err := store.NewDB(cfg)
	if err != nil {
		return err
	}
	defer pg.Close()

	h, err := handler.NewHandler(pg)
	if err != nil {
		return err
	}

	http.HandleFunc("/incoming/insert", h.InsertIncoming)
	http.HandleFunc("/incoming/get", h.GetIncoming)
	http.HandleFunc("/incoming/update", h.UpdateIncoming)
	http.HandleFunc("/incoming/delete", h.DeleteIncoming)
	http.HandleFunc("/spend/insert", h.InsertSpends)
	http.HandleFunc("/spend/get", h.GetSpend)
	http.HandleFunc("/spend/update", h.UpdateSpend)
	http.HandleFunc("/spend/delete", h.DeleteSpend)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.ServiceHost, cfg.ServicePort), nil)
}
