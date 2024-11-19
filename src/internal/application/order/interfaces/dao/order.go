package dao

import "github.com/D1sordxr/aviasales/src/internal/db/models"

type OrderDAO interface {
	CreateOrder(order models.Order) (int, error)
	GetOrder(id int) (models.Order, error)
}
