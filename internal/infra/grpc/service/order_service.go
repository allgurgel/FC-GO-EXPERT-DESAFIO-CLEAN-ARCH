package service

import (
	"context"

	"github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/infra/grpc/pb"
	"github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrderRequest) (*pb.OrderList, error) {
	dto := usecase.ListOrdersInputDTO{
		Page:  int(in.Page),
		Limit: int(in.Limit),
		Sort:  in.Sort,
	}

	output, err := s.ListOrdersUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	var ordersModel []*pb.Order
	ordersModel = []*pb.Order{}

	for _, order := range *output {
		ordersModel = append(ordersModel, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.OrderList{Orders: ordersModel}, nil
}
