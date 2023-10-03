package yodal

const Type = "yodal"

type Yodal struct {
	ppGram float64
}

func New(ppGram float64) *Yodal {
	return &Yodal{
		ppGram: ppGram,
	}
}

func (y *Yodal) Price(grams float64) float64 {
	return grams * y.ppGram
}

func (y *Yodal) Type() string {
	return Type
}
