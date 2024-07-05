package orders

import (
    "gorm.io/gorm"
)

type OrderRepository struct {
    DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{DB: db}
}

func (or *OrderRepository) saveOrder(order *Order) error {
    return or.DB.Create(&order).Error
}

func (or *OrderRepository) getOrderById(id string) (*Order, error) {
    var order Order
    if err := or.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
    return &order, nil
}

func (or *OrderRepository) getAllOrders() ([]Order, error) {
    var orders []Order
    if err := or.DB.Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (or *OrderRepository) updateOrder(id string, updatedOrder *Order) error {
    if err := or.DB.Model(&Order{}).Where("id = ?", id).Updates(updatedOrder).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) deleteOrder(id string) error {
    if err := or.DB.Where("id = ?", id).Delete(&Order{}).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) searchOrdersByUserID(userID string) ([]Order, error) {
    var orders []Order
    if err := or.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (or *OrderRepository) searchOrdersByStatus(status string) ([]Order, error) {
    var orders []Order
    if err := or.DB.Where("status = ?", status).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}
