package main

import (
	"log"
	"net/http"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// port := os.Getenv("PORT")
	port := "5001"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// router := gin.New()
	// router.Use(gin.Logger())
	// // router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")
	// router.Static("/build/web", "build/web")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })

	fileServer := http.FileServer(http.Dir("./build/web")) // New code
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// router.Run(":" + port)
}
