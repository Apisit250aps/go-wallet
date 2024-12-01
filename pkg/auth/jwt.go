package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

func GenerateToken(userID string, secret string, expiryHours int) (string, error) {
    claims := JWTClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiryHours)).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, jwt.ErrInvalidKey
}