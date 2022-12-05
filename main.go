package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	h "git.linuxit.ro/chr45/e-factura/handlers"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//GOOS=windows GOARCH=386 go build -o efactura32.exe .
//GOOS=windows GOARCH=amd64 go build -o efactura32.exe .

func main() {
	godotenv.Load(".env")

	ip := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	gin.SetMode(gin.ReleaseMode)
	web := gin.Default()

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}
	web.SetFuncMap(funcMap)
	web.LoadHTMLGlob("templates/*.html")
	web.GET("/", h.Upload)
	web.POST("/convert", h.Convert)
	web.GET("/um", h.UnitatiMasura)
	log.Println(color.Ize(color.Green, fmt.Sprintf("Server starting at http://%s:%s", ip, port)))
	web.Run(ip + ":" + port)
}
