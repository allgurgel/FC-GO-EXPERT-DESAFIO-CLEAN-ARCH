package usecase

import "github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/entity"

type ListOrdersInputDTO struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type OrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute(input ListOrdersInputDTO) (*[]OrdersOutputDTO, error) {
	orders, err := l.OrderRepository.List(input.Page, input.Limit, input.Sort)
	if err != nil {
		return &[]OrdersOutputDTO{}, nil
	}

	var ordersOutput []OrdersOutputDTO
	ordersOutput = []OrdersOutputDTO{}

	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return &ordersOutput, nil
}
