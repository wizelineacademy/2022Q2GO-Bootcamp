package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/entity"
	"encoding/json"
)
func main(){
	fmt.Println("hello external Api Client")
	response, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	var catFact entity.CatFact
	json.Unmarshal(responseData, &catFact)
	
	fmt.Println(catFact.Fact)
}