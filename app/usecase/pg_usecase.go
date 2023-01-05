package usecase

import (
	"membership/domain"
)

type pgUsecase struct {
	repo domain.PGRepository
}

func NewPGUsecase(repo domain.PGRepository) domain.PGUsecase {
	return &pgUsecase{repo: repo}
}

func (p *pgUsecase) chargePermata(orderId string, payment map[string]interface{}, resp *domain.BankPaymentResponse) (err error) {
	// _, err := p.repo.BankPermataCharge(orderId, payment)
	// if err != nil {
	// 	return
	// }
	resp.PaymentType = "bank_transfer"
	resp.OrderId = "123123123213"
	resp.Bank = "bni"
	resp.VaNumber = "12321421214"
	resp.GrossAmount = 100

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
