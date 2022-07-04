package usecase

import (
	"fmt"
	"time"

	"github.com/Qaz3xsw/gophermart/internal/gophermart/domain/core"
	"github.com/Qaz3xsw/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		GetOrdersByUser(string) []core.OrderNumber
	}

	ListUserOrdersInputPort interface {
		Execute(user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error)
	}

	ListUserOrdersOutputDTO struct {
		UploadedAt time.Time `json:"uploaded_at"` // nolint:tagliatelle // ok
		Number     string    `json:"number"`
		Status     string    `json:"status"`
		Accrual    int       `json:"accrual"`
	}

	ListUserOrders struct {
		Repo ListUserOrdersRepository
	}
)

func NewListOrderNums(repo ListUserOrdersRepository) *ListUserOrders {
	return &ListUserOrders{
		Repo: repo,
	}
}

func (l *ListUserOrders) Execute(user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	orders := l.Repo.GetOrdersByUser(user.ID())

	fmt.Println(orders)

	return []ListUserOrdersOutputDTO{
		ListUserOrdersOutputDTO{},
	}, nil
}
