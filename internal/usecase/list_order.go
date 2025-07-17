package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCaseInterface interface {
	Execute(input ListOrderInputDTO) ([]ListOrderOutputDTO, error)
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}
func (l *ListOrderUseCase) Execute(input ListOrderInputDTO) ([]ListOrderOutputDTO, error) {
	order, err := l.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	output := make([]ListOrderOutputDTO, 0, len(order))
	for _, o := range order {
		output = append(output, ListOrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.Price + o.Tax,
		})
	}

	return output, nil
}
