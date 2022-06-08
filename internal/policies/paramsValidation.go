package policies

func IsValidParams(readingType string) bool {
	return readingType == "odd" || readingType == "even"
}
