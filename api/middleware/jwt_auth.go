package middleware

import (
	log2 "file_manager/internal/common/log"
	"file_manager/internal/enums"
	"file_manager/internal/helpers"
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

func verifyToken(c *gin.Context) enums.Error {
	cookieToken, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log2.Errorf("cookie token is empty, err:[%v]", err)
			return enums.NewCustomHttpError(http.StatusUnauthorized, "cookie token empty")
		}
		log2.Error("get token fail, err:[%v]", err)
		return enums.NewCustomSystemError("get token fail")
	}

	tokenStr := cookieToken.Value
	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		log2.Errorf("parse private key error, err:[%v]", err)
		return enums.NewCustomSystemError("parse private key error")
	}

	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log2.Errorf("err signature invalid, err:[%v]", err)
			return enums.NewCustomHttpError(http.StatusUnauthorized, "signature_invalid")
		}
		log2.Errorf("parse with claims error, err:[%v]", err)
		return enums.NewCustomSystemError("parseWithClaims error")
	}

	if !tkn.Valid {
		log2.Error("token invalid")
		return enums.NewCustomHttpError(http.StatusUnauthorized, "token_invalid")
	}
	return nil
}
