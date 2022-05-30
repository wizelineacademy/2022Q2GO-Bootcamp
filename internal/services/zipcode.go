package services

import (
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
)

func CreateCSV(empData model.ZipCodeInfo) error {
	err := repository.CreateCSV(empData)
	return err
}
