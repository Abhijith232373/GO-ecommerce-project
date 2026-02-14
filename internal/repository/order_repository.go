package repository

import (
	"e-commerce/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB)*OrderRepository{
	return &OrderRepository{DB: db}
}

func (r *OrderRepository)CreatOrder(order *models.Order,items []models.OrderItem)error{
	return r.DB.Transaction(func(tx *gorm.DB)error{
		
		if err:=tx.Create(order).Error;err!=nil{
			return err
		}

		for i:=range items {
			items[i].OrderID=order.ID
		}
		if err:=tx.Create(&items).Error;err!=nil{
			return err
		}
		return nil
	})
}

func (r *OrderRepository)GetUserOrders(userID uint)([]models.Order,error){
	var orders []models.Order

	err:=r.DB.
	Preload("OrderItems").
	Preload("OrderItems.Product").
	Where("user_id=?",userID).
	Order("id DESC").
	Find(&orders).Error

	return orders,err
}

func (r *OrderRepository)GetOrderByID(orderID uint)(*models.Order,error){
	var order models.Order

	err:=r.DB.
	Preload("OrderItems").
	Preload("OrderItems.Product").
	First(&order,orderID).Error

	return &order,err

}

func (r *OrderRepository)ClearCart(userID uint)error{
	return r.DB.
	Where("user_id=?",userID).
	Delete(&models.Cart{}).Error
}


func (r *OrderRepository)CreateOrder(order *models.Order, items []models.OrderItem) error {

	tx := r.DB.Begin()

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range items {
		items[i].OrderID = order.ID
	}

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}