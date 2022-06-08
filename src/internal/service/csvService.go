package service
import (
	"log"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/entity"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/repository"
)
type CsvService interface{
	ExtractCsvRecords() ([]entity.CsvRecord, error)
	SaveCsvRecords(records []entity.CatFact) (error)
}
type csvService struct{
	repo repository.CsvRepository
}
func NewCsvService(repo repository.CsvRepository) CsvService{
	return &csvService{repo}
} 
func (s *csvService) ExtractCsvRecords() ([]entity.CsvRecord, error){
	records, err := s.repo.ReadRecords()
	if err != nil {
		log.Printf("SVC-ERROR repository %v", err)
		return nil, err
	}
	return records, nil
}
func (s *csvService) SaveCsvRecords(records []entity.CatFact) (error){
	err := s.repo.WriteRecords(records)
	if err != nil {
		log.Printf("SVC-ERROR repository %v", err)
		return err
	}
	return nil
}