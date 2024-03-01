package requests

type CreateOrdersRequests struct {
	CarId           string `json:"carId"`
	OrderDate       string `json:"orderDate"`
	PickupDate      string `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
}

type UpdateOrdersByOrdersIdRequest struct {
	CarId           string `json:"carId"`
	OrderDate       string `json:"orderDate"`
	PickupDate      string `json:"pickupDate"`
	DropoffDate     string `json:"dropoffDate"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
	OrdersId        string `json:"OrdersId"`
}

type GetDetailOrdersRequest struct {
	OrdersId string `json:"OrdersId"`
}
type DeleteOrdersRequest struct {
	OrdersId string `json:"OrdersId"`
}
