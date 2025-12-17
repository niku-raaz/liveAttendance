package utils

import(
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct{
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

func VerifyJWT(tokenStr string, secret string)(*JwtClaims, error){
	token, err := jwt.Parse(tokenStr,  func(t *jwt.Token) (interface{}, error) {
		return []byte(secret) ,nil
	})

	if err!=nil || !token.Valid{
		return nil, errors.New("invalid token")
	}

	claimsMap, ok := token.Claims.(jwt.MapClaims)
    if !ok{
		return nil,errors.New("invalid claims")
	}

	userId, ok1 :=claimsMap["userId"].(string)
	role,   ok2 :=claimsMap["role"].(string)

	if !ok1 || !ok2 {
		return nil, errors.New("missing claims")
	}

	return &JwtClaims{
		UserID: userId,
		Role: role,
	}, nil 

}
