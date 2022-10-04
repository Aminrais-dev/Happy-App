package midtrans

type DataInterface interface {
	WeebHookUpdate(orderid, status string) (string, error)
}

type UsecaseInterface interface {
	WeebHookTransaction(orderid, status string) (string, error)
}
