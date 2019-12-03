package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GuilhermeAbacherli/todolistgo/dao"
	"github.com/GuilhermeAbacherli/todolistgo/entity"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// AuthMiddleware checks if the token is valid before proceed. It returns 401 status if the client is not valid
func AuthMiddleware(next http.HandlerFunc) http.Handler {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return tokenValidation(jwtMiddleware, next)
}

func tokenValidation(middleware *jwtmiddleware.JWTMiddleware, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		tokenString := request.Header.Get("Authorization")
		if tokenString == "" {
			log.Println("Required authorization token not found")
			json.NewEncoder(writer).Encode("Required authorization token not found.")
			return
		}
		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte("secret"), nil
		})

		var user entity.User
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["exp"] != nil {
				user.Exp = claims["exp"].(string)
			}
		}

		expiration, err := time.Parse(time.RFC3339, user.Exp)
		if err != nil {
			log.Println("Error while parsing time from user exp: ", err)
			json.NewEncoder(writer).Encode("Error while validating the token.")
		}

		if time.Now().UTC().After(expiration) {
			log.Println("The token has expired")
			json.NewEncoder(writer).Encode("The token has expired.")
			return
		}

		// Let secure process the request. If it returns an error,
		// that indicates the request should not continue.
		err = middleware.CheckJWT(writer, request)

		// If there was an error, do not continue.
		if err != nil {
			log.Println("Error while validating the token: ", err)
			json.NewEncoder(writer).Encode("Error while validating the token.")
			return
		}

		handler.ServeHTTP(writer, request)
	})
}

// Register register a new user
func (dc *DatabaseConnection) Register(writer http.ResponseWriter, request *http.Request) {
	logRequest(request)

	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Println("Error on decoding user to json: ", err)
		json.NewEncoder(writer).Encode("Error on decoding user to json.")
		return
	}

	var existingUser entity.User
	dao.SelectOneUser(dc.Client, bson.M{"username": user.Username}, &existingUser)
	if existingUser.Username != "" {
		log.Println("This username already exists")
		json.NewEncoder(writer).Encode("This username already exists.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		log.Println("Error while hashing password: ", err)
		json.NewEncoder(writer).Encode("Error while hashing password.")
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
		json.NewEncoder(writer).Encode("Error while decoding user's request info.")
		return
	}

	var userFoundInDatabase entity.User
	dao.SelectOneUser(dc.Client, bson.M{"username": userRequestingLogin.Username}, &userFoundInDatabase)
	if userFoundInDatabase.Username == "" {
		log.Println("User not found")
		json.NewEncoder(writer).Encode("User not found.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFoundInDatabase.Password), []byte(userRequestingLogin.Password))
	if err != nil {
		log.Println("Invalid password")
		json.NewEncoder(writer).Encode("Invalid password.")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userFoundInDatabase.Username,
		"iat":      time.Now().Format(time.RFC3339),
		"exp":      time.Now().Add(time.Second * 25).Format(time.RFC3339),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println("Error while generating token: ", err)
		json.NewEncoder(writer).Encode("Error while generating token.")
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
		log.Println("Required authorization token not found")
		json.NewEncoder(writer).Encode("Required authorization token not found.")
		return
	}
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	var user entity.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] != nil {
			user.Username = claims["username"].(string)
		}
		if claims["iat"] != nil {
			user.Iat = claims["iat"].(string)
		}
		if claims["exp"] != nil {
			user.Exp = claims["exp"].(string)
		}
		// user.Email = claims["email"].(string) // Must be in the payload when login generate the token
		json.NewEncoder(writer).Encode(user)
		return
	}

	log.Println("Error while validating the token: ", err)
	json.NewEncoder(writer).Encode("Error while validating the token.")
}
