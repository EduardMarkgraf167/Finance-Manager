package handler

import (
	"net/http"
	"../utils"
	"path/filepath"
)


func HomeHandler(writer http.ResponseWriter, request *http.Request){
	if request.URL.Path != "/home" && request.URL.Path != "/" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	cookie, err := request.Cookie("cookie")

	if err != nil || cookie == nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	if !utils.IsUserActive(cookie.Value) {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	writer.WriteHeader(200)
	htmlPath := filepath.FromSlash("../pages/homePage.html")
	http.ServeFile(writer, request, htmlPath)
}
