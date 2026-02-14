package service

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repository"
	"errors"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
	CartRepo  *repository.CartRepository
}

type AddressInput struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
}

func NewOrderService(orderRepo *repository.OrderRepository, cartRepo *repository.CartRepository) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo,
		CartRepo:  cartRepo,
	}
}

func (s *OrderService) CreateOrder(userID uint) error {

	cartItems, err := s.CartRepo.GetByUser(userID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return errors.New("cart is empty")
	}

	var total float64
	var orderItems []models.OrderItem

	for _, cart := range cartItems {

		price := cart.Product.Price
		qty := cart.Quantity

		total += price * float64(qty)

		orderItems = append(orderItems, models.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  qty,
			Price:     price,
		})
	}

	order := models.Order{
		UserID:      userID,
		TotalAmount: total,
		Status:      "pending",
	}

	err = s.OrderRepo.CreateOrder(&order, orderItems)
	if err != nil {
		return err
	}

	return s.OrderRepo.ClearCart(userID)
}

func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.OrderRepo.GetUserOrders(userID)
}

func (s *OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}

func (s *OrderService) CreateOrderWithAddress(userID uint, addr AddressInput) error {

	cartItems, err := s.CartRepo.GetByUser(userID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return errors.New("cart is empty")
	}

	var total float64
	var orderItems []models.OrderItem

	for _, cart := range cartItems {

		total += cart.Product.Price * float64(cart.Quantity)

		orderItems = append(orderItems, models.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
		})
	}

	order := models.Order{
		UserID:      userID,
		TotalAmount: total,
		Status:      "pending",

		FullName: addr.FullName,
		Phone:    addr.Phone,
		Address:  addr.Address,
		City:     addr.City,
		State:    addr.State,
		Pincode:  addr.Pincode,
	}

	err = s.OrderRepo.CreateOrder(&order, orderItems)
	if err != nil {
		return err
	}

	return s.OrderRepo.ClearCart(userID)
}