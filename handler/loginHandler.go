package handler

import (
	"net/http"
	"path/filepath"
)

func LoginHandler(writer http.ResponseWriter, request *http.Request){

	if request.URL.Path != "/login" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	htmlPath := filepath.FromSlash("./pages/loginPage.html")
	http.ServeFile(writer, request, htmlPath)
}
