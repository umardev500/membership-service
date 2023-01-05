package usecase

import (
	"errors"
	"fmt"
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

func (p *pgUsecase) chargePermata(orderId string, payment map[string]interface{}, resp *domain.BankPaymentResponse) (err error) {
	response, err := p.repo.BankPermataCharge(orderId, payment)
	if err != nil {
		return
	}
	statusCode, _ := strconv.Atoi(response.StatusCode)
	if !(statusCode >= 200 && statusCode < 300) {
		fmt.Println("Error status code")
		return errors.New(response.StatusMessage)
	}

	resp.PaymentType = response.PaymentType
	resp.OrderId = response.OrderID
	resp.Bank = "permata"
	resp.VaNumber = response.PermataVaNumber
	resp.GrossAmount = helper.RemovePriceDot(response.GrossAmount)

	return
}

func (p *pgUsecase) BankCharge(orderId string, payment map[string]interface{}) (resp *domain.BankPaymentResponse, err error) {
	resp = &domain.BankPaymentResponse{}

	if payment["payment"] != nil {
		payment = payment["payment"].(map[string]interface{})
	}

	if payment != nil {
		bankTransfer := payment["bank_transfer"].(map[string]interface{})
		bank := bankTransfer["bank"]

		if bank == "permata" {
			err = p.chargePermata(orderId, payment, resp)
		}
	}

	return
}
