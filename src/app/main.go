package main

import (
	"log"
	"net/http"
	"realworld/user"
	"realworld/utils"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// url := "postgres://stef:stefpass@localhost:5437/realworld"

	// _, err := pgx.Connect(context.Background(), url)

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "Successfully connecrted")
	// }
	db, cleanup := utils.DbConnection()
	utils.SeedDbUgly(db)

	r := mux.NewRouter()

	userService := user.NewUserService(db)
	userController := user.NewUserController(userService)
	userController.RegisterHandlers(r)

	defer cleanup() // Ensure cleanup is called at the end
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
