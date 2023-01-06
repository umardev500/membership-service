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

	var bank, vaNumber string
	permataVaNumber := response.PermataVaNumber
	if permataVaNumber != "" {
		bank = "permata"
		vaNumber = response.PermataVaNumber
	}
	if permataVaNumber == "" {
		bank = response.VaNumbers[0].Bank
		vaNumber = response.VaNumbers[0].VaNumber
	}

	resp.PaymentType = response.PaymentType
	resp.OrderId = response.OrderID
	resp.Bank = bank
	resp.VaNumber = vaNumber
	resp.GrossAmount = helper.RemovePriceDot(response.GrossAmount)

	return
}
