package web

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type WebListOrderHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type WebListOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebListOrderHandler(OrderRepository entity.OrderRepositoryInterface) *WebListOrderHandler {
	return &WebListOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebListOrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	listOrder := usecase.NewListOrderUseCase(h.OrderRepository)
	output, err := listOrder.Execute(usecase.ListOrderInputDTO{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
