package utils

import (
	"../model"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func IsUserActive(cookie string) bool {
	return model.CookieManagement.Cookies[cookie] != ""
}

func GetActiveUser(cookie string) (model.User, error) {
	userId := model.CookieManagement.Cookies[cookie]

	for _, user := range model.UserCredentialsManagement.Users {
		if user.Id != userId {
			continue
		}
		return user, nil
	}

	return model.User{}, errors.New("user not found")
}

func Authenticate(name string, password string) bool {
	correspondingUser, err := GetCorrespondingUser(name)

	if err != nil {
		return false
	}

	hashed, _ := HashPasswd(password, correspondingUser.Salt)
	return hashed == correspondingUser.Passwd
}

func GetCorrespondingUser(name string) (model.User, error) {
	var correspondingUser model.User
	found := false

	for _, user := range model.UserCredentialsManagement.Users {
		if user.Name != name {
			continue
		}

		correspondingUser = user
		found = true
		break
	}

	if !found {
		return model.User{}, errors.New("user not found")
	}

	return correspondingUser, nil
}

func HashPasswd(passwd string, salt string) (string, string) {
	if salt == "" {
		salt = RandomString(10)
	}

	passwdSalted := passwd + salt
	saltedByte := []byte(passwdSalted)
	hashedByte := md5.Sum(saltedByte)
	hashedPasswd := hex.EncodeToString(hashedByte[:])

	return hashedPasswd, salt
}

func Login(user model.User) string {
	cookieJar := &model.CookieManagement
	cookie := RandomString(20)
	cookieJar.Cookies[cookie] = user.Id
	return cookie
}
