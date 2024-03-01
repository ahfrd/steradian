package requests

type CreateCarRequests struct {
	CarName   string `json:"carName"`
	DayRate   int    `json:"dayRate"`
	MonthRate int    `json:"monthRate"`
	ImageCar  string `json:"imageCar"`
}

type UpdateCarByCarIdRequest struct {
	CarName   string `json:"carName"`
	DayRate   int    `json:"dayRate"`
	MonthRate int    `json:"monthRate"`
	ImageCar  string `json:"imageCar"`
	CarId     string `json:"carId"`
}

type GetDetailCarRequest struct {
	CarId string `json:"carId"`
}
type DeleteCarRequest struct {
	CarId string `json:"carId"`
}
