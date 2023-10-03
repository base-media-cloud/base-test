package product

type Products []Product

type Product struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Weight float64 `json:"weight"`
}
