package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// CookieConfig interface
type CookieConfig interface {
	GetAccessTokenMaxAge() int
	GetRefreshTokenMaxAge() int
	GetDomain() string
	GetSecure() bool
	GetHTTPOnly() bool
}

// CookieConfigBody - implement CookieConfig interface
type CookieConfigBody struct {
	secure             bool
	httpOnly           bool
	accessTokenMaxAge  int
	refreshTokenMaxAge int
	domain             string
}

// GetAccessTokenMaxAge -
func (c *CookieConfigBody) GetAccessTokenMaxAge() int {
	return c.accessTokenMaxAge
}

// GetRefreshTokenMaxAge -
func (c *CookieConfigBody) GetRefreshTokenMaxAge() int {
	return c.refreshTokenMaxAge
}

// GetDomain -
func (c *CookieConfigBody) GetDomain() string {
	return c.domain
}

// GetSecure -
func (c *CookieConfigBody) GetSecure() bool {
	return c.secure
}

// GetHTTPOnly -
func (c *CookieConfigBody) GetHTTPOnly() bool {
	return c.httpOnly
}

// ReadCookieConfig -
func ReadCookieConfig(cookieSect string, maxAgeSect string) CookieConfig {

	secure := viper.GetBool(fmt.Sprintf("%s.secure", cookieSect))
	httpOnly := viper.GetBool(fmt.Sprintf("%s.httpOnly", cookieSect))

	accessTokenMaxAge := viper.GetInt(fmt.Sprintf("%s.accessToken", maxAgeSect))
	refreshTokenMaxAge := viper.GetInt(fmt.Sprintf("%s.refreshToken", maxAgeSect))
	domain := viper.GetString(fmt.Sprintf("%s.domain", cookieSect))
	return &CookieConfigBody{
		secure:             secure,
		httpOnly:           httpOnly,
		accessTokenMaxAge:  accessTokenMaxAge,
		refreshTokenMaxAge: refreshTokenMaxAge,
		domain:             domain,
	}
}
