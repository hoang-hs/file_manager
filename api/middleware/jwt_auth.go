package middleware

import (
	"file_manager/internal/enums"
	"file_manager/internal/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RequiredJwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := verifyToken(c)
		if err != nil {
			c.JSON(err.GetHttpCode(), err.GetMessage())
			c.Abort()
			return
		}
		c.Next()
	}
}

func verifyToken(c *gin.Context) enums.Error {
	cookieToken, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Printf("cookie token is empty, err:[%v]", err.Error())
			return enums.NewCustomHttpError(http.StatusUnauthorized, "cookie token empty")
		}
		log.Printf("get token fail, err:[%v]", err.Error())
		return enums.NewCustomSystemError("get token_fail")
	}

	tokenStr := cookieToken.Value
	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		log.Printf("parse private key error, err:[%v]", err.Error())
		return enums.NewCustomSystemError("parse private key error")
	}

	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("err signature invalid, err:[%v]", err.Error())
			return enums.NewCustomHttpError(http.StatusUnauthorized, "signature_invalid")
		}
		log.Printf("parse with claims error, err:[%v]", err.Error())
		return enums.NewCustomSystemError("parseWithClaims error")
	}

	if !tkn.Valid {
		log.Printf("token invalid")
		return enums.NewCustomHttpError(http.StatusUnauthorized, "token_invalid")
	}
	return nil
}
