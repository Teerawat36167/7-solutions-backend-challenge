package api

import (
	counter "github.com/Teerawat36167/PieFireDire/internal/util"
	"github.com/gorilla/mux"
)

type Handler struct {
	counter *counter.MeatCounter
}

func NewHandler() *Handler {
	return &Handler{
		counter: counter.NewMeatCounter(),
	}
}

func SetupRouter(handler *Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/beef/summary", handler.HandleBeefSummary).Methods("GET")
	return r
}
