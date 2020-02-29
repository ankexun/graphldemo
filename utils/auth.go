package utils

import (
	"encoding/json"
	"graphqldemo/models"
	"log"
	"net/http"
)

// login handler
func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	if _, err := user.Authenticate(); err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusUnauthorized)
		result := `{
	"code": 401,
	"msg": "未授权的访问",
}`
		response.Write([]byte(result))
		return
	}

	token, err := GenerateToken(user.Username, user.Password)
	if err != nil {
		log.Println(err)
	}

	response.Header().Set("content-type", "application/json")
	response.Header().Set("token", token)
	response.Write([]byte(`{ "token": "` + token + `" }`))
}
