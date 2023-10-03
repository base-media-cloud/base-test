package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/base-media-cloud/base-test/internal/domain/services"
	"github.com/base-media-cloud/base-test/internal/mocks/ports_mocks"
)

func TestSvc_DeliveryPrice(t *testing.T) {
	ctl := gomock.NewController(t)

	delSvc := ports_mocks.NewMockDeliverySvc(ctl)
	delSvc.EXPECT().Price(gomock.Any()).Return(1.23).Times(1)

	svc := services.New(delSvc)
	result := svc.DeliveryPrice(1.23)

	assert.Equal(t, 1.23, result)
}

func TestSvc_DeliveryType(t *testing.T) {
	ctl := gomock.NewController(t)

	delSvc := ports_mocks.NewMockDeliverySvc(ctl)
	delSvc.EXPECT().Type().Return("mock-type").Times(1)

	svc := services.New(delSvc)
	result := svc.DeliveryType()

	assert.Equal(t, "mock-type", result)
}
