package ports

import (
	"github.com/Nathan-S19/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
