package royalmail

const Type = "royal_mail"

type RoyalMail struct {
	ppGram float64
}

func New(ppGram float64) *RoyalMail {
	return &RoyalMail{
		ppGram: ppGram,
	}
}

func (r *RoyalMail) Price(grams float64) float64 {
	return grams * r.ppGram
}

func (r *RoyalMail) Type() string {
	return Type
}
