package jwt

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, username string, secretKey string) (string, error) {
	// Implementation for generating JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"user_name": username,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // Token expires in 15 minute
	})
	key := []byte(secretKey)
	tokenstring, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func GenerateRefreshToken(userID int, secretKey string) (string, error) {
	bytes := make([]byte, 32)
	_, err :=rand.Read(bytes)
	if err != nil{
		return "", errors.New("failed to generate refresh token")
	}
    return hex.EncodeToString(bytes), nil
}

func ValidateToken(tokenstr, secretKey string, withClaimValidation bool) (int, string, error) {
	if secretKey == ""{
	   return 0, "", errors.New("Secret key is empty")
	}
    
	
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenstr, &claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
               return nil, errors.New("unexpected signing method")
            }  
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
               return nil, errors.New("unexpected signing method")
            }  
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
    return 0, "", errors.New("invalid user_id claim")
    }
	username, ok := claims["user_name"].(string)
	if !ok {
    return 0, "", errors.New("invalid user_id claim")
    }

	return int(userID), username, nil
}
