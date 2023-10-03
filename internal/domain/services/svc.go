package services

import "github.com/base-media-cloud/base-test/internal/domain/ports"

type Svc struct {
	delSvc ports.DeliverySvc
}

func New(delSvc ports.DeliverySvc) *Svc {
	return &Svc{
		delSvc: delSvc,
	}
}

func (s *Svc) DeliveryPrice(grams float64) float64 {
	return s.delSvc.Price(grams)
}

func (s *Svc) DeliveryType() string {
	return s.delSvc.Type()
}
