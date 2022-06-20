package hook

import (
	"fmt"
	"os"
)

func Getcwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFilePath := fmt.Sprintf("%s/data/data.csv", pwd)

	return dataFilePath
}
