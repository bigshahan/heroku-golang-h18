package main

import (
	"os"

	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup Routes
	router := httprouter.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("{ \"code\":3.14159265359, \"description\":\"HI!\" }"))
	})

	router.POST("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("{ \"code\":3.14159265359, \"description\":\"HI!\" }"))
	})

	// Setup Server
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	// start server
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)
}
