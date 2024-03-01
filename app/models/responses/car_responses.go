package responses

type SelectListDataCarResponses struct {
	CarId     string `json:"carId"`
	CarName   string `json:"carName"`
	DayRate   int    `json:"dayRate"`
	MonthRate int    `json:"monthRate"`
	ImageChar string `json:"imageChar"`
}
