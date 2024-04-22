package ports

import (
	"github.com/Nathan-S19/microservices/order/internal/application/core/domain"
)

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
