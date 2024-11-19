package dao

import "src/internal/db/models"

type OrderDAO interface {
	CreateOrder(order models.Order) (int, error)
	GetOrder(id int) (models.Order, error)
}
