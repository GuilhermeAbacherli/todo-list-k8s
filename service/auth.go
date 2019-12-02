package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GuilhermeAbacherli/todolistgo/dao"
	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register register a new user
func (dc *DatabaseConnection) Register(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)

	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Println("Error on decoding user to json")
		json.NewEncoder(writer).Encode(err.Error())
		return
	}

	var existingUser entity.User
	dao.SelectOneUser(dc.Client, bson.M{"username": user.Username}, &existingUser)
	if existingUser.Username != "" {
		log.Println("This username already exists")
		json.NewEncoder(writer).Encode("This username already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		log.Println("Error while hashing password: ", err)
		json.NewEncoder(writer).Encode("Error while hashing password")
	}
	user.Password = string(hash)

	insertedID := dao.InsertOneUser(dc.Client, user)
	json.NewEncoder(writer).Encode(insertedID)
	return
}

// Login authorize an existing user
func (dc *DatabaseConnection) Login(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	var userRequestingLogin entity.User
	err := json.NewDecoder(request.Body).Decode(&userRequestingLogin)
	if err != nil {
		log.Println("Error on decoding user to json: ", err)
		json.NewEncoder(writer).Encode("Error while decoding user's request info")
		return
	}

	var userFoundInDatabase entity.User
	dao.SelectOneUser(dc.Client, bson.M{"username": userRequestingLogin.Username}, &userFoundInDatabase)
	if userFoundInDatabase.Username == "" {
		log.Println("User not found")
		json.NewEncoder(writer).Encode("User not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFoundInDatabase.Password), []byte(userRequestingLogin.Password))
	if err != nil {
		log.Println("Invalid password")
		json.NewEncoder(writer).Encode("Invalid password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userFoundInDatabase.Username,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println("Error while generating token", err)
		json.NewEncoder(writer).Encode("Error while generating token")
		return
	}

	userFoundInDatabase.Token = tokenString
	userFoundInDatabase.Password = ""

	json.NewEncoder(writer).Encode(userFoundInDatabase)

}

// Profile returns information about an existing user
func (dc *DatabaseConnection) Profile(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)
	tokenString := request.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Missing token in request header")
		json.NewEncoder(writer).Encode("Missing token in request header")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	var user entity.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		// user.Email = claims["email"].(string) // Must be in the payload when login will generate the token
		json.NewEncoder(writer).Encode(user)
		return
	}

	log.Println("Error while validating the token", err.Error())
	json.NewEncoder(writer).Encode(err.Error())
}
