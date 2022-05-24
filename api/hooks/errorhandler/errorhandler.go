package errorhandler

import (
	"fmt"
	"net/http"
)

var (
	ErrUnsupportedMediaType        = New(http.StatusUnsupportedMediaType, http.StatusUnsupportedMediaType, "UN_SUPPORT_MEDIA_TYPE", "API support only application/json type")
	ErrSomeFieldAreNotValid        = New(http.StatusBadRequest, http.StatusBadRequest, "SOME_FIELDS_ARE_NOT_VALID", "some fields are not valid in JSON format")
	ErrNotFound                    = New(http.StatusNotFound, http.StatusNotFound, "NOT_FOUND", "not found this resource")
	ErrUnauthorized                = New(http.StatusUnauthorized, http.StatusUnauthorized, "UN_AUTHORIZED", "you don't have access to this resource")
	ErrMethodNotAllowed            = New(http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "method not allowed, please see the API document for more information")
	ErrNotValidItemId              = New(http.StatusBadRequest, http.StatusBadRequest, "NOT_VALID_ITEM_ID", "not valid item id")
	ErrNotFoundAnyItemWithThisId   = New(http.StatusNotFound, http.StatusNotFound, "NOT_FOUND_ANY_ITEM_WITH_THIS_ID", "not found any item with this id")
	ErrInternalServerError         = New(http.StatusInternalServerError, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error")
	ErrCanNotAccessToTheseResource = New(http.StatusForbidden, http.StatusForbidden, "CAN_NOT_ACCESS_TO_THESE_RESOURCES", "you can't access to these resources")
)

type Error struct {
	HttpCode    int    `json:"-"`
	Code        int    `json:"code" bson:"code"`
	Message     string `json:"error" bson:"message"`
	Description string `json:"description" bson:"description"`
}

func New(httpCode int, code int, message string, description string) *Error {
	return &Error{httpCode, code, message, description}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error Code is %d - %s - %s", e.Code, e.Message, e.Description)
}

// func NewErrorHandler(err error, c echo.Context) {
// 	speError := New(http.StatusInternalServerError,
// 		http.StatusInternalServerError,
// 		http.StatusText(http.StatusInternalServerError),
// 		http.StatusText(http.StatusInternalServerError),
// 	)
// 	if he, ok := err.(*Error); ok {
// 		speError = he
// 	}

// 	if !c.Response().Committed() {
// 		c.JSON(speError.HttpCode, speError)
// 	}
// }
