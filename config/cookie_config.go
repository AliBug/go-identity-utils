package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// CookieConfig - contaiin cookie setting
type CookieConfig struct {
	secure             bool
	httpOnly           bool
	accessTokenMaxAge  int
	refreshTokenMaxAge int
	domain             string
}

// GetAccessTokenMaxAge -
func (c *CookieConfig) GetAccessTokenMaxAge() int {
	return c.accessTokenMaxAge
}

// GetRefreshTokenMaxAge -
func (c *CookieConfig) GetRefreshTokenMaxAge() int {
	return c.refreshTokenMaxAge
}

// GetDomain -
func (c *CookieConfig) GetDomain() string {
	return c.domain
}

// GetSecure -
func (c *CookieConfig) GetSecure() bool {
	return c.secure
}

// GetHTTPOnly -
func (c *CookieConfig) GetHTTPOnly() bool {
	return c.httpOnly
}

// ReadCookieConfig -
func ReadCookieConfig(cookieSect string, maxAgeSect string) *CookieConfig {
	secure := viper.GetBool(fmt.Sprintf("%s.accessToken", cookieSect))
	httpOnly := viper.GetBool(fmt.Sprintf("%s.httpOnly", cookieSect))

	accessTokenMaxAge := viper.GetInt(fmt.Sprintf("%s.accessToken", maxAgeSect))
	refreshTokenMaxAge := viper.GetInt(fmt.Sprintf("%s.refreshToken", maxAgeSect))
	domain := viper.GetString(fmt.Sprintf("%s.domain", cookieSect))
	return &CookieConfig{
		secure:             secure,
		httpOnly:           httpOnly,
		accessTokenMaxAge:  accessTokenMaxAge,
		refreshTokenMaxAge: refreshTokenMaxAge,
		domain:             domain,
	}
}
