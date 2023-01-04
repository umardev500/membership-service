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

func (p *pgUsecase) BankCharge() (resp interface{}, err error) {
	resp, err = p.repo.BankPermataCharge()
	if err != nil {
		return
	}

	return
}
