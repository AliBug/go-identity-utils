package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// TokenConfig -
type TokenConfig struct {
	accessExpirationSeconds  time.Duration
	refreshExpirationSeconds time.Duration
	accessTokenSecret        []byte
	refreshTokenSecret       []byte
	issuer                   string
}

// GetIssuer -
func (t *TokenConfig) GetIssuer() string {
	return t.issuer
}

// GetAccessTokenSecret -
func (t *TokenConfig) GetAccessTokenSecret() []byte {
	return t.accessTokenSecret
}

// GetRefreshTokenSecret -
func (t *TokenConfig) GetRefreshTokenSecret() []byte {
	return t.refreshTokenSecret
}

// GetAccessExpirationSeconds -
func (t *TokenConfig) GetAccessExpirationSeconds() time.Duration {
	return t.accessExpirationSeconds
}

// GetRefreshExpirationSeconds -
func (t *TokenConfig) GetRefreshExpirationSeconds() time.Duration {
	return t.refreshExpirationSeconds
}

// ReadTokenConfig - read Token Config
func ReadTokenConfig(tokenSect string, maxAgeSect string) *TokenConfig {
	accessTokenSecret := []byte(viper.GetString(fmt.Sprintf("%s.accessSecret", tokenSect)))
	refreshTokenSecret := []byte(viper.GetString(fmt.Sprintf("%s.refreshSecret", tokenSect)))
	accessExpirationSeconds := time.Duration(viper.GetInt(fmt.Sprintf("%s.accessToken", maxAgeSect)))
	refreshExpirationSeconds := time.Duration(viper.GetInt(fmt.Sprintf("%s.refreshToken", maxAgeSect)))
	issuer := viper.GetString(fmt.Sprintf("%s.issuer", tokenSect))
	return &TokenConfig{
		accessTokenSecret:        accessTokenSecret,
		refreshTokenSecret:       refreshTokenSecret,
		accessExpirationSeconds:  accessExpirationSeconds,
		refreshExpirationSeconds: refreshExpirationSeconds,
		issuer:                   issuer,
	}
}
