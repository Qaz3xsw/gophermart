package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Qaz3xsw/gophermart/internal/gophermart/domain/usecase"
	"github.com/Qaz3xsw/gophermart/internal/gophermart/handler/middleware"
	"github.com/Qaz3xsw/gophermart/internal/sharedkernel"
)

// PostRegisterOrder POST /api/user/orders — загрузка пользователем номера заказа для расчёта.
// 200 — номер заказа уже был загружен этим пользователем;
// 202 — новый номер заказа принят в обработку;
// 400 — неверный формат запроса;
// 401 — пользователь не аутентифицирован;
// 409 — номер заказа уже был загружен другим пользователем;
// 422 — неверный формат номера заказа;
// 500 — внутренняя ошибка сервера.
func PostRegisterOrder(registerOrderUsecase usecase.RegisterOrderPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		orderNumber, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		err = registerOrderUsecase.Execute(orderNumber, user)
		if err != nil {
			log.Println(err)
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// GetOrders GET /api/user/orders — получение списка загруженных пользователем номеров заказов,
// статусов их обработки и информации о начислениях
// 200 — успешная обработка запроса.
// 204 — нет данных для ответа.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func GetOrders(listOrdersUsecase usecase.ListUserOrdersInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		_, err := listOrdersUsecase.Execute(user)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// GetBalance GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя
// 200 — успешная обработка запроса.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func GetBalance(showBalanceUsecase usecase.ShowUserBalanceInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		_, err := showBalanceUsecase.Execute(user)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// PostWithdraw POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта
// в счёт оплаты нового заказа
// 200 — успешная обработка запроса;
// 401 — пользователь не авторизован;
// 402 — на счету недостаточно средств;
// 422 — неверный номер заказа;
// 500 — внутренняя ошибка сервера.
func PostWithdraw(withdrawFundsUsecase usecase.WithdrawFundsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		dto := usecase.WithdrawUserFundsInputDTO{} // nolint:exhaustivestruct // ok,  exhaustive // ok.

		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal(bytes, &dto)
		if err != nil {
			log.Println(err)
		}

		err = withdrawFundsUsecase.Execute(user, dto)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// GetWithdrawals GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта
// 200 — успешная обработка запроса;
// 204 — нет ни одного списания.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func GetWithdrawals(listWithdrawalsUsecase usecase.ListUserWithdrawalsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		_, err := listWithdrawalsUsecase.Execute(user)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
