package ports

//go:generate mockgen -source delivery.go -destination=../../mocks/ports_mocks/delivery.go -package=ports_mocks

type DeliverySvc interface {
	Price(grams float64) float64
	Type() string
}
