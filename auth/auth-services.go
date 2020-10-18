package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type IJwtAuthService interface {
	GenerateJwtToken(username string) string
	ValidateJwtToken(token string) (*jwt.Token, error)
}

type tokenBody struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtService struct {
	secretKey       string
	tokenExpiration uint64
	logger          *log.Logger
}

func (service *JwtService) GenerateJwtToken(username string) string {
	claims := &tokenBody{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(service.tokenExpiration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		service.logger.WithField("error", err).Error("token encoding failed")
	}
	return t
}

func (service *JwtService) ValidateJwtToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}

func NewJwtService(secret string, tokenExpiration uint64, logger *log.Logger) *JwtService {
	return &JwtService{secretKey: secret, tokenExpiration: tokenExpiration, logger: logger}
}
