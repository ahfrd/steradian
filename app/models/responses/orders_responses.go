package responses

type SelectListDataOrdersResponses struct {
	CarId           string `json:"carId"`
	OrderDate       string `json:"orderDate"`
	PickupDate      string `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
	OrdersId        string `json:"ordersId"`
}
