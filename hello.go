package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Message string `json:"message"`
}

func main() {
	httpServer()
}

func httpServer() {
	this, err := loadConfigFile("config/this.json")
	if err != nil {
		fmt.Println(err)
	}
	that, err := loadConfigFile("config/that.json")
	if err != nil {
		fmt.Println(err)
	}
	multix := http.NewServeMux()
	multix.HandleFunc("/this", handler(*this))
	multix.HandleFunc("/that", handler(*that))

	fmt.Println("Go Web App Started on Port 3000")
	err = http.ListenAndServe(":3000", multix)
	fmt.Println(err)
}

func handler(cfg Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, cfg.Message)
	})

}

func loadConfigFile(filepath string) (*Config, error) {
	var config Config
	fi, err := os.Open(filepath)
	//defer fi.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(fi)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
