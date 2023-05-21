package auth

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"restapi-golang/model/domain"
)

func Login(writer http.ResponseWriter, http *http.Request) {

}

func Register(writer http.ResponseWriter, request *http.Request) {
	var userInput domain.Users
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&userInput); err != nil {
		fmt.Println("gagal input")
	}
	defer request.Body.Close()

	//hash menggunakan bcyrpy
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert into database
}

func Logout(writer http.ResponseWriter, request *http.Request) {

}
