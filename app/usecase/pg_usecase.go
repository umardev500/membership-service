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

func (p *pgUsecase) chargePermata(orderId string, payment map[string]interface{}) {
	// _, err := p.repo.BankPermataCharge(orderId, payment)
	// if err != nil {
	// 	return
	// }
}

func (p *pgUsecase) BankCharge(orderId string, payment map[string]interface{}) (resp interface{}, err error) {
	if payment["payment"] != nil {
		payment = payment["payment"].(map[string]interface{})
	}

	if payment != nil {
		bankTransfer := payment["bank_transfer"].(map[string]interface{})
		bank := bankTransfer["bank"]

		if bank == "permata" {
			p.chargePermata(orderId, payment)
		}
	}

	return
}
