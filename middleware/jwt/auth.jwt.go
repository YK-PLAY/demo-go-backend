package jwt

import (
	"fmt"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtMiddleware struct {
}

func New(mw *JwtMiddleware) (*JwtMiddleware, error) {
	return mw, nil
}

func (mw *JwtMiddleware) MiddlewareFunction() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.middlewareImpl(c)
	}
}

func (mw *JwtMiddleware) middlewareImpl(c *gin.Context) {
	matched, err := regexp.MatchString("^\\/api\\/v\\d?\\/auth?", c.Request.URL.Path)
	if err != nil {
		fmt.Printf("Regex error: %s\n", err.Error())
	}

	if matched {
		fmt.Printf("Match non-auth url: %s\n", c.Request.URL.Path)
		c.Next()
	} else {
		tokenString := c.Request.Header.Get("JWT-TOKEN")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("hello"), nil
		})

		if err != nil {
			fmt.Printf("Parse token error: %s\n", err.Error())
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Printf("Hello %s\n", claims["userId"])

		c.Next()
	}
}
