package ups

const Type = "ups"

type UPS struct {
	ppGram float64
}

func New(ppGram float64) *UPS {
	return &UPS{
		ppGram: ppGram,
	}
}

func (u *UPS) Price(grams float64) float64 {
	return grams * u.ppGram
}

func (u *UPS) Type() string {
	return Type
}
