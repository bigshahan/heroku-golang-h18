package main

import (
	"os"
	"time"

	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup Routes
	router := httprouter.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.POST("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		// simulate API calls, etc...
		time.Sleep(300 * time.Millisecond)
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
