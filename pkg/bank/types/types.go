package types

//Money представляет собой денежную сумму в минимальных единицах
type Money int64

// Currency код валюты
type Currency string

// код валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)
// PAN номер карты
type PAN string

//Card информация о платежной карте
type Card struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
}
//Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
}
 
type PaymentSource struct {
	Type  string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в драмах
}