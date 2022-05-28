package services

import (
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
)

func CreateCSV(empData map[string]interface{}) {
	repository.CreateCSV(empData)
	return
}
