package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var secret = "Tra138381813@MNCKBNSRH"

func GenerateToken(email string, id string) (string, error) {
	payload := jwt.MapClaims{
		"email": email,
		"id":    id,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signed, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ValidateToken(tokenString string) (map[string]interface{}, error) {
	errResp := fmt.Errorf("need signin")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	var payload = map[string]interface{}{}
	claims := token.Claims.(jwt.MapClaims)
	payload["email"] = claims["email"]
	payload["id"] = claims["id"]

	// exp := fmt.Sprintf("%v", claims["exp"])

	// now := time.Now()
	// expTime, _ := time.Parse(time.RFC3339, exp)

	// fmt.Println("Now: ", now)
	// fmt.Println("Exp: ", expTime)

	// if !now.Before(expTime) {
	// 	return nil, fmt.Errorf("expired")
	// }

	return payload, nil
}
