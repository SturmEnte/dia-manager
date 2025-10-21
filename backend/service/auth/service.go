package auth

import (
	"context"
	"errors"
	"time"

	"dia-manager-backend/config"
	"dia-manager-backend/utils"

	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(username string, password string) (string, error) {

    var id string

    hashedPassword, err := utils.HashPassword(password)

    if err != nil {
        println(err)
        return "", errors.New("failed to hash password")
    }
    
    err = config.DB.QueryRow(context.Background(), `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`, username, hashedPassword).Scan(&id)

    if err != nil {
        println(err.Error())
        return "", errors.New("failed to insert the new user into the database")
    }

    return id, nil
}

func CreateToken(cfg *config.Config, id string, username string) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  id,
        "username": username,
        "exp":     time.Now().Add(time.Duration(cfg.TokenLifetime) * time.Minute).Unix(),
        "iat":      time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    secretKey := []byte(cfg.TokenSecret)
    tokenString, err := token.SignedString(secretKey)
    
    if err != nil {
        println(err.Error())
        return "", errors.New("failed to create JWT token")
    }

    return tokenString, nil
}

func DisableToken(cfg *config.Config, tokenString string) (error) {
    
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(cfg.TokenSecret), nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        return errors.New("invalid token")
    }

    // Extract claims and read "exp"
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return errors.New("invalid token claims")
    }

    expVal, ok := claims["exp"]
    if !ok {
        return errors.New("token missing exp claim")
    }

    // MapClaims typically stores numeric values as float64
    var expTime time.Time
    switch v := expVal.(type) {
    case float64:
        expTime = time.Unix(int64(v), 0)
    case int64:
        expTime = time.Unix(v, 0)
    default:
        return errors.New("unexpected exp claim type")
    }
    
    remaining := time.Until(expTime)
    
    if remaining <= 0 {
        return errors.New("token already expired")
    }

    // Store invalid token so that the auth middleware knows this token is not valid even though its not expired
    _, err = config.DB.Exec(context.Background(), `INSERT INTO invalid_tokens (token, expires) VALUES ($1, $2)`, tokenString, expTime)
    
    if err != nil {
        return err
    }

    return nil
}