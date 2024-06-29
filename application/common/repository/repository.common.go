package repository

import (
	"antrein/dd-dashboard-analytic/application/common/resource"
	"antrein/dd-dashboard-analytic/model/config"
)

type CommonRepository struct {
}

func NewCommonRepository(cfg *config.Config, rsc *resource.CommonResource) (*CommonRepository, error) {

	commonRepo := CommonRepository{}
	return &commonRepo, nil
}
