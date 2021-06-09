package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// TokenConfig -
type TokenConfig interface {
	GetIssuer() string
	GetAccessTokenSecret() []byte
	GetRefreshTokenSecret() []byte
	GetAccessExpirationSeconds() time.Duration
	GetRefreshExpirationSeconds() time.Duration
}

// TokenConfigBody - implement TokenConfig interface
type TokenConfigBody struct {
	accessExpirationSeconds  time.Duration
	refreshExpirationSeconds time.Duration
	accessTokenSecret        []byte
	refreshTokenSecret       []byte
	issuer                   string
}

// GetIssuer -
func (t *TokenConfigBody) GetIssuer() string {
	return t.issuer
}

// GetAccessTokenSecret -
func (t *TokenConfigBody) GetAccessTokenSecret() []byte {
	return t.accessTokenSecret
}

// GetRefreshTokenSecret -
func (t *TokenConfigBody) GetRefreshTokenSecret() []byte {
	return t.refreshTokenSecret
}

// GetAccessExpirationSeconds -
func (t *TokenConfigBody) GetAccessExpirationSeconds() time.Duration {
	return t.accessExpirationSeconds
}

// GetRefreshExpirationSeconds -
func (t *TokenConfigBody) GetRefreshExpirationSeconds() time.Duration {
	return t.refreshExpirationSeconds
}

// ReadTokenConfig - read Token Config
func ReadTokenConfig(tokenSect string, maxAgeSect string) TokenConfig {
	accessTokenSecret := []byte(viper.GetString(fmt.Sprintf("%s.accessSecret", tokenSect)))
	refreshTokenSecret := []byte(viper.GetString(fmt.Sprintf("%s.refreshSecret", tokenSect)))
	accessExpirationSeconds := time.Duration(viper.GetInt(fmt.Sprintf("%s.accessToken", maxAgeSect)))
	refreshExpirationSeconds := time.Duration(viper.GetInt(fmt.Sprintf("%s.refreshToken", maxAgeSect)))
	issuer := viper.GetString(fmt.Sprintf("%s.issuer", tokenSect))

	if accessExpirationSeconds <= 0 || refreshExpirationSeconds <= 0 || accessExpirationSeconds >= refreshExpirationSeconds {
		log.Fatalf("Some Max age are error! accessExpirationSeconds: %v | refreshExpirationSeconds: %v | accessTokenMaxAge>=refreshTokenMaxAge: %v", accessExpirationSeconds, refreshExpirationSeconds, accessExpirationSeconds >= refreshExpirationSeconds)
	}

	return &TokenConfigBody{
		accessTokenSecret:        accessTokenSecret,
		refreshTokenSecret:       refreshTokenSecret,
		accessExpirationSeconds:  time.Second * accessExpirationSeconds,
		refreshExpirationSeconds: time.Second * refreshExpirationSeconds,
		issuer:                   issuer,
	}
}
