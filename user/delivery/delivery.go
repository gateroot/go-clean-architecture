package delivery

type UserDelivery interface {
	Start(addr string) error
}
