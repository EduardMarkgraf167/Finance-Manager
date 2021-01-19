package main

import (
	"./handler"
	"./manager"
	"./model"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	flags := model.Flag
	flags.InitFlags()
	certPath := filepath.FromSlash("./cert")
	certFile := filepath.Join(certPath, "cert.pem")
	keyFile := filepath.Join(certPath, "key.pem")

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/home", handler.HomeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/loginManager", manager.LoginManager)

	log.Fatalln(http.ListenAndServeTLS(":"+flags.Port, certFile, keyFile, nil))
}