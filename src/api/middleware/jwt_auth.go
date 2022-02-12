package middleware

import (
	"file_manager/src/common/log"
	"file_manager/src/core/errors"
	"file_manager/src/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func verifyToken(c *gin.Context) errors.Error {
	cookieToken, err := c.Request.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Errorf("cookie token is empty, err:[%v]", err)
			return errors.NewCustomHttpError(http.StatusUnauthorized, "cookie token empty")
		} else {
			log.Errorf("get token fail, err:[%v]", err)
			return errors.NewCustomSystemError("get token fail")
		}
	}

	tokenStr := cookieToken.Value
	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		log.Errorf("parse private key error, err:[%v]", err)
		return errors.NewCustomSystemError("parse private key error")
	}

	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Errorf("err signature invalid, err:[%v]", err)
			return errors.NewCustomHttpError(http.StatusUnauthorized, "signature_invalid")
		}
		log.Errorf("parse with claims error, err:[%v]", err)
		return errors.NewCustomSystemError("parseWithClaims error")
	}

	if !tkn.Valid {
		log.Error("token invalid")
		return errors.NewCustomHttpError(http.StatusUnauthorized, "token_invalid")
	}
	return nil
}
