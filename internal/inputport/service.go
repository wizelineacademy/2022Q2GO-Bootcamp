package inputport

import (
	"github.com/jesusrevilla/capstone/internal/app"
	"github.com/jesusrevilla/capstone/internal/inputport/http"
)

// Services contains the ports services
type Services struct {
	Server *http.Server
}

// NewServices instatiates the services of input ports
func NewServices(appServices app.Services) Services {
	return Services{Server: http.NewServer(appServices)}
}
