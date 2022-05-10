package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luischitala/2022Q2GO-Bootcamp/controller"
	infrastructure "github.com/luischitala/2022Q2GO-Bootcamp/infrastructure/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	const port string = ":5050"
	//Router
	r := mux.NewRouter()
	router := infrastructure.NewMuxRouter(r)

	//Controller
	hc := controller.NewHomeController()

	//Routes
	router.GET("/", hc.Home)
	router.SERVE(port)
}

// func BindRoutes(s server.Server, r *mux.Router) {

// 	//To bypass the middleware

// 	//Create a new hub
// 	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
// 	r.HandleFunc("/characters", handlers.ListCharacterHandler(s)).Methods(http.MethodGet)
// 	r.HandleFunc("/readCsv", handlers.ReadCsvHandler(s)).Methods(http.MethodGet)
// }

// PORT := os.Getenv("PORT")
// JWT_SECRET := os.Getenv("JWT_SECRET")
// DATABASE_URL := os.Getenv("DATABASE_URL")

// s, err := server.NewServer(context.Background(), &server.Config{
// 	JWTSecret:   JWT_SECRET,
// 	Port:        PORT,
// 	DatabaseUrl: DATABASE_URL,
// })
// if err != nil {
// 	log.Fatal(err)
// }
// s.Start(BindRoutes)
