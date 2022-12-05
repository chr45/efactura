package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	h "git.linuxit.ro/chr45/e-factura/handlers"
	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

//GOOS=windows GOARCH=386 go build -o efactura32.exe .
//GOOS=windows GOARCH=amd64 go build -o efactura32.exe .

func main() {
	log.SetPrefix(color.Ize(color.Cyan, "contractare: "))
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	}
	log.SetPrefix(color.Ize(color.Cyan, "e-factura: "))
	godotenv.Load(".env")

	ip := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	http.HandleFunc("/", h.Upload)
	http.HandleFunc("/um", h.UnitatiMasura)
	http.HandleFunc("/convert", h.Convert)

	cfg := fmt.Sprintf("%s:%s", ip, port)
	log.Printf("Starting server http://%s\n", cfg)
	err := http.ListenAndServe(cfg, nil)
	if err != nil {
		log.Fatalf("Error on starting server %s: %v ", cfg, err)
	}
}
