package handler

import (
	"github.com/esvarez/go-api/pkg/web"
	"log"
	"net/url"
	"strconv"
)

const (
	items               = "items"
	tpe                 = "type"
	itemsPerWorker      = "items_per_workers"
	defaultItems        = 5
	defaultItemsWorkers = 1
)

var (
	validTypes = map[string]bool{"odd": true, "even": true}
)

func getValidType(params url.Values) (string, *web.AppError) {
	if _, ok := params[tpe]; !ok {
		log.Println("no type")
		return "", &web.BadRequestError

	}
	if val, ok := validTypes[params[tpe][0]]; !ok {
		log.Println("invalid type", val)
		return "", &web.BadRequestError
	}
	return params[tpe][0], nil
}

func getValidItems(params url.Values) (int, *web.AppError) {
	return getIntOrDefault(params, items, defaultItems)
}

func getValidItemsWorkers(params url.Values) (int, *web.AppError) {
	return getIntOrDefault(params, itemsPerWorker, defaultItemsWorkers)
}

func getIntOrDefault(params url.Values, key string, defaultValue int) (int, *web.AppError) {
	if val, ok := params[key]; ok {
		itms, _ := strconv.Atoi(val[0])
		if itms < 1 {
			return 0, &web.BadRequestError
		}
		return itms, nil
	}
	return defaultValue, nil
}
