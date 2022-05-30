package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/controller"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func mustJson(t *testing.T, v interface{}) string {
	t.Helper()
	out, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestCreateCsv(t *testing.T) {

	tests := []struct {
		name    string
		url     string
		param   string
		want    string
		wantErr bool
	}{
		{
			"valid",
			"/",
			"01530",
			"{\"post code\":\"01530\",\"country\":\"Mexico\",\"country abbreviation\":\"MX\",\"places\":[{\"latitude\":\"19.3771\",\"longitude\":\"-99.2272\",\"place name\":\"Corpus Christi\",\"state\":\"Distrito Federal\",\"state abbreviation\":\"DIF\"}]}",
			false,
		},
		{
			"invalid",
			"/",
			"015fd",
			"{\"code\":404,\"error\":\"NOT_FOUND_ANY_ITEM_WITH_THIS_ID\",\"description\":\"not found any item with this id\"}",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/zipcode/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.param)

			fmt.Println("Name test: ", tt.name)

			// Assertions
			if !tt.wantErr {
				if assert.NoError(t, controller.CreateCsv(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)

					s := strings.TrimSpace(fmt.Sprintf(tt.want))
					s1 := strings.TrimSpace(fmt.Sprintf(rec.Body.String()))

					assert.Equal(t, s, s1)
				}
			} else {
				assert.Error(t, controller.CreateCsv(c))
			}
		})
	}
}
