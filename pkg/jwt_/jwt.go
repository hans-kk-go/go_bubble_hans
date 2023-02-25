package jwt_

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId int64 `json:"id"`
	jwt.StandardClaims
}

func JwtGettoken(userid int64) string {
	mySigningKey := []byte("AllYourBase")
	c := Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60, //生效时间
			ExpiresAt: time.Now().Unix() + 5*50*50,
			Issuer:    "hans",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
		return "token_get err"
	} else {
		fmt.Println(signedString)
		return signedString
	}

}

func ParseToken(signedString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(signedString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	return token, claims, err

}
