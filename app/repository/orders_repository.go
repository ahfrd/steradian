package repository

import (
	"database/sql"
	"fmt"
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type OrdersRepositoryImpl struct {
	db *sql.DB
}

func NewOrdersRepositoryImpl(db *sql.DB) OrdersRepository {
	return &OrdersRepositoryImpl{
		db: db,
	}
}

func (cr *OrdersRepositoryImpl) CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) error {
	q := `insert into orders(car_id,order_date,pickup_date,dropoff_date,pickup_location,dropoff_location) values (?,?,?,?,?,?)`
	if _, err := cr.db.ExecContext(ctx, q, request.CarId, request.OrderDate, request.PickupDate, request.DropoffDate, request.PickupLocation, request.DropoffLocation); err != nil {
		return err
	}
	return nil
}

func (cr *OrdersRepositoryImpl) SelectListDataOrders(ctx *gin.Context) ([]responses.SelectListDataOrdersResponses, error) {
	var datas []responses.SelectListDataOrdersResponses

	q := fmt.Sprintf(`select order_id,car_id,order_date,pickup_date,dropoff_date,pickup_location,dropoff_location from orders`)
	result, err := cr.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var data responses.SelectListDataOrdersResponses
		if err := result.Scan(&data.OrdersId, &data.CarId, &data.OrderDate, &data.PickupDate, &data.DropoffDate, &data.PickupLocation, &data.DropoffLocation); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (cr *OrdersRepositoryImpl) SelectDataOrdersById(ctx *gin.Context, ordersId string) (*responses.SelectListDataOrdersResponses, error) {
	var data responses.SelectListDataOrdersResponses
	q := fmt.Sprintf(`select order_id,car_id,order_date,pickup_date,dropoff_date,pickup_location,dropoff_location from orders where order_id = '%s'`, ordersId)
	if err := cr.db.QueryRowContext(ctx, q).Scan(&data.OrdersId, &data.CarId, &data.OrderDate, &data.PickupDate, &data.DropoffDate, &data.PickupLocation, &data.DropoffLocation); err != nil {
		return nil, err
	}
	return &data, nil
}

func (cr *OrdersRepositoryImpl) UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) error {
	q := fmt.Sprintf(`update orders set car_id = '%s', order_date = '%s', pickup_date = '%s', dropoff_date = '%s', pickup_location = '%s', dropoff_location = '%s' where order_id ='%s'`,
		request.CarId, request.OrderDate, request.PickupDate, request.DropoffDate, request.PickupLocation, request.DropoffLocation, request.OrdersId)
	if _, err := cr.db.ExecContext(ctx, q); err != nil {
		return err
	}
	return nil
}

func (cr *OrdersRepositoryImpl) DeleteOrdersByOrdersId(ctx *gin.Context, ordersId string) error {
	q := fmt.Sprintf(`delete from orders where order_id = '%s'`, ordersId)
	fmt.Println(q)
	if _, err := cr.db.ExecContext(ctx, q); err != nil {
		return err
	}
	return nil
}
