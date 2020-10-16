package user

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "strings"
)

func AuthorizeJWT() gin.HandlerFunc {
    return func(c *gin.Context) {

        const BearerSchema = "Bearer "

        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Provide JWT"})
            return
        }

        claims := jwt.MapClaims{}

        extractedToken := strings.Split(authHeader, "Bearer ")

        if len(extractedToken) == 2 {
            authHeader = strings.TrimSpace(extractedToken[1])
        } else {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Format of Authorization Token "})
            return
        }

        parsedToken, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
            return os.Getenv("JWT_SECRET"), nil
        })

        if err != nil {
            if err == jwt.ErrSignatureInvalid || err == jwt.ErrInvalidKeyType {
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
                return
            }
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad Request"})
            return
        }

        if !parsedToken.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
            return
        }

        c.Next()

    }

}
