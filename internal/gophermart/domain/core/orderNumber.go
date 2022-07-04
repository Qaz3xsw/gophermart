package core

import "github.com/Qaz3xsw/gophermart/internal/sharedkernel"

// OrderNumber is now represent users registered order
type OrderNumber struct {
	id      string
	user    string
	status  sharedkernel.Status
	number  int
	accrual int
}

func NewOrderNumber(number, accrual int, userID string, status sharedkernel.Status) OrderNumber {
	return OrderNumber{
		id:      sharedkernel.NewUUID(),
		user:    userID,
		number:  number,
		status:  status,
		accrual: 0,
	}
}
