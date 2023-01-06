package usecase

import (
	"errors"
	"membership/domain"
	"membership/helper"
	"strconv"
)

type pgUsecase struct {
	repo domain.PGRepository
}

func NewPGUsecase(repo domain.PGRepository) domain.PGUsecase {
	return &pgUsecase{repo: repo}
}

func (p *pgUsecase) BankCharge(orderId string, payment map[string]interface{}) (resp *domain.BankPaymentResponse, err error) {
	resp = &domain.BankPaymentResponse{}

	response, err := p.repo.BankCharge(orderId, payment)
	if err != nil {
		return
	}
	statusCode, _ := strconv.Atoi(response.StatusCode)
	if !(statusCode >= 200 && statusCode < 300) {
		return nil, errors.New(response.StatusMessage)
	}

	resp.PaymentType = response.PaymentType
	resp.OrderId = response.OrderID
	resp.Bank = "permata"
	resp.VaNumber = response.PermataVaNumber
	resp.GrossAmount = helper.RemovePriceDot(response.GrossAmount)

	return
}
