package payments

import (
    "gorm.io/gorm"
)

type PaymentRepository struct {
    DB *gorm.DB
}

func (repo *PaymentRepository) GetAllPayments() ([]Payment, error) {
    var payments []Payment
    err := repo.DB.Find(&payments).Error
    return payments, err
}

func (repo *PaymentRepository) CreatePayment(payment *Payment) error {
    return repo.DB.Create(payment).Error
}

func (repo *PaymentRepository) GetPaymentByID(id string) (*Payment, error) {
    var payment Payment
    err := repo.DB.First(&payment, "id = ?", id).Error
    return &payment, err
}

func (repo *PaymentRepository) UpdatePayment(payment *Payment) error {
    return repo.DB.Save(payment).Error
}

func (repo *PaymentRepository) DeletePayment(id string) error {
    return repo.DB.Delete(&Payment{}, "id = ?", id).Error
}

func (repo *PaymentRepository) SearchPaymentsByUserID(userID string) ([]Payment, error) {
    var payments []Payment
    err := repo.DB.Where("user_id = ?", userID).Find(&payments).Error
    return payments, err
}

func (repo *PaymentRepository) SearchPaymentsByOrderID(orderID string) ([]Payment, error) {
    var payments []Payment
    err := repo.DB.Where("order_id = ?", orderID).Find(&payments).Error
    return payments, err
}

func (repo *PaymentRepository) SearchPaymentsByStatus(status string) ([]Payment, error) {
    var payments []Payment
    err := repo.DB.Where("status = ?", status).Find(&payments).Error
    return payments, err
}
