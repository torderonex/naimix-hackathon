package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"os"
	"time"
)

var (
	salt       = os.Getenv("SALT")
	tokenTTL   = 120 * time.Hour
	signingKey = os.Getenv("SIGNING_KEY")
)

type AuthService struct {
	repo repo.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repo.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	ownerId, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		slog.Error(err.Error())
		return ownerId, err
	}

	return ownerId, err
}

func (s *AuthService) GenerateToken(ctx context.Context, login, password string) (string, error) {
	user, err := s.repo.GetUserByCredentials(ctx, login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GetUserById(ctx context.Context, id int) (entities.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *AuthService) ParseToken(ctx context.Context, accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
