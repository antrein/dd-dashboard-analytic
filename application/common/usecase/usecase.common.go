package usecase

import (
	"antrein/dd-dashboard-analytic/application/common/repository"
	"antrein/dd-dashboard-analytic/model/config"
)

type CommonUsecase struct {
}

func NewCommonUsecase(cfg *config.Config, repo *repository.CommonRepository) (*CommonUsecase, error) {

	commonUC := CommonUsecase{}
	return &commonUC, nil
}
