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

func (p *pgUsecase) BankCharge(bank, orderId string, data interface{}) (resp interface{}, err error) {
	if bank == "permata" {
		resp, err = p.repo.BankPermataCharge(orderId, data)
		if err != nil {
			return
		}
	}

	return
}
