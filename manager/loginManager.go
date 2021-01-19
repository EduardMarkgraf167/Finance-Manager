package manager

import (
	"net/http"
	"../utils"
)

func LoginManager(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/loginManager" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Method != "POST" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	userName := request.FormValue("uname")
	passwd := request.FormValue("psw")

	if !utils.Authenticate(userName, passwd) {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	user, err := utils.GetCorrespondingUser(userName)

	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	cookie := utils.Login(user)

	http.SetCookie(writer, &http.Cookie{
		Name:  "cookie",
		Value: cookie,
		Path:  "/",
	})
	http.Redirect(writer, request, "/home", 200)
}
