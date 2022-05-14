package server
import "net/http"

type CsvDataLines struct {
	Column1 string
	Column2 string
 }
 var datas []CsvDataLines = []CsvDataLines{}
func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
