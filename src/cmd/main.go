package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rrune/AoCZoom/api"
	"github.com/rrune/AoCZoom/models"
	"github.com/rrune/AoCZoom/router"
	"github.com/rrune/AoCZoom/util"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.OpenFile("../data/zoom.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(2 | 3)
	log.Println("")

	var config models.Config
	ymlData, err := os.ReadFile("../data/config.yml")
	util.CheckPanic(err)
	err = yaml.Unmarshal(ymlData, &config)
	util.CheckPanic(err)

	a := api.New(config.ImagePath)

	log.Printf("Running on Port %s", config.Port)
	if config.Ssl {
		log.Fatal(http.ListenAndServeTLS(":"+config.Port, config.Certfile, config.Keyfile, router.NewRouter(a, config)))
	} else {
		log.Fatal(http.ListenAndServe(":"+config.Port, router.NewRouter(a, config)))
	}
}
