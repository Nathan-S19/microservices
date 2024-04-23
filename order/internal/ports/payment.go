package ports

import (
	"context"
	"github.com/Nathan-S19/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
