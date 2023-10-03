package product

type Res struct {
	Data []ResProd `json:"data"`
}

type ResProd struct {
	Name          string `json:"name"`
	DeliveryPrice string `json:"delivery_price"`
	ProductPrice  string `json:"product_price"`
	TotalPrice    string `json:"total_price"`
}

func NewRes(prods []ResProd) Res {
	return Res{
		Data: prods,
	}
}
