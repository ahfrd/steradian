package repository

import (
	"database/sql"
	"fmt"
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type carRepositoryImpl struct {
	db *sql.DB
}

func NewCarRepositoryImpl(db *sql.DB) CarRepository {
	return &carRepositoryImpl{
		db: db,
	}
}

func (cr *carRepositoryImpl) CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) (int, error) {
	q := `insert into car(car_name,day_rate,month_rate,image_car) values (?,?,?,?)`
	result, err := cr.db.ExecContext(ctx, q, request.CarName, request.DayRate, request.MonthRate, request.ImageCar)
	if err != nil {
		return 0, err
	}
	lastInsertId, _ := result.LastInsertId()
	return int(lastInsertId), nil
}

func (cr *carRepositoryImpl) SelectListDataCar(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error) {
	var datas []responses.SelectListDataCarResponses

	q := fmt.Sprintf(`select car_id,car_name,day_rate,month_rate,image_car from car`)
	result, err := cr.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var data responses.SelectListDataCarResponses
		if err := result.Scan(&data.CarId, &data.CarName, &data.DayRate, &data.MonthRate, &data.ImageChar); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (cr *carRepositoryImpl) SelectDataCarById(ctx *gin.Context, carId string) (*responses.SelectListDataCarResponses, error) {
	var data responses.SelectListDataCarResponses
	q := fmt.Sprintf(`select car_id,car_name,day_rate,month_rate,image_car from car where car_id = '%s'`, carId)
	if err := cr.db.QueryRowContext(ctx, q).Scan(&data.CarId, &data.CarName, &data.DayRate, &data.MonthRate, &data.ImageChar); err != nil {
		return nil, err
	}
	return &data, nil
}

func (cr *carRepositoryImpl) UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error {
	q := fmt.Sprintf(`update car set car_name = '%s', day_rate = '%d', month_rate = '%d', image_car = '%s' where car_id = '%s'`, request.CarName, request.DayRate, request.MonthRate, request.ImageCar, request.CarId)
	if _, err := cr.db.ExecContext(ctx, q); err != nil {
		return err
	}
	return nil
}

func (cr *carRepositoryImpl) DeleteCarByCarId(ctx *gin.Context, carId string) error {
	q := fmt.Sprintf(`delete from car where car_id = '%s'`, carId)
	if _, err := cr.db.ExecContext(ctx, q); err != nil {
		return err
	}
	return nil
}
