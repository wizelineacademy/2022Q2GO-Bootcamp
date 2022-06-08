package entity

//Type to store csv records from CSV file
type CsvRecord struct {
	Key int64
	Value  string
}

//Type to be used in CatFact Api
type CatFact struct {
	Fact string `json:"fact"`
	Length int64 `json:"length"`
}