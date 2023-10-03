package ports

type Services interface {
	DeliveryPrice(grams float64) float64
	DeliveryType() string
}
