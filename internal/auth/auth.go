package auth

import (
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/hberkayozdemir/hibemi-be/internal/user"
)

var secret = []byte("YsO8ecO8ayBoaWxtaSBzaXppIHNldml5b3I=")

const DB_URL = "mongodb+srv://hbo:Hbo.1998@hibemibe.zbozrlc.mongodb.net/?retryWrites=true&w=majority"

func VerifyToken(bearerToken string) (bool, error) {

	splitToken := strings.Split(bearerToken, "Bearer ")
	if len(splitToken) != 2 {
		return false, NotAuthorizedError
	}

	token := splitToken[1]

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return false, NotAuthorizedError
	}

	jwtClaims, isVerified := parsedToken.Claims.(jwt.MapClaims)

	if !isVerified && !parsedToken.Valid {
		return false, NotAuthorizedError
	}
	parsedClaims := GetTokenClaims(jwtClaims)

	userRepository := user.NewRepository(DB_URL)
	_, err = userRepository.GetUser(parsedClaims.Issuer)

	if err != nil {
		return false, NotAuthorizedError
	}

	return true, nil
}

func GetTokenClaims(jwtMapClaims jwt.MapClaims) Claims {

	jsonString, _ := json.Marshal(jwtMapClaims)

	parsedClaims := Claims{}

	json.Unmarshal(jsonString, &parsedClaims)
	return parsedClaims
}
