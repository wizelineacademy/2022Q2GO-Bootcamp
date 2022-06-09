package repository

import "errors"

var (
	ErrFileCSV        = errors.New("csv file error")
	ErrParseCharacter = errors.New("parse character error")
	ErrDataPermission = errors.New("csv file permission error")
)
