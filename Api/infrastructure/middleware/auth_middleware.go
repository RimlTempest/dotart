package middleware

import (
	"fmt"
	"time"

	"api/constants"
	"api/domain/model"
	"api/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// SetAuth authentication
func SetAuth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           constants.REALM,
		Key:             []byte(constants.SESSION_SECRET),
		Timeout:         time.Hour * 24, //24時間後
		MaxRefresh:      time.Hour,
		IdentityKey:     constants.IDENTITY_KEY,
		PayloadFunc:     payload,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		LoginResponse:   loginResponse,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	return authMiddleware, err
}

// ペイロードのクレーム設定 IdentityKeyでUserIdをマッピング
func payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			constants.IDENTITY_KEY: v.UserId,
		}
	}
	return jwt.MapClaims{}
}

// identitityHandler クレームからログインIDを取得
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var user model.User
	// UserIdを元にcolomnの取得
	constants.GetDB().Where("user_id = ?", claims[constants.IDENTITY_KEY]).First(&user)

	return user
}

// authenticator ログイン認証をするコールバック
func authenticator(c *gin.Context) (interface{}, error) {
	var loginVals model.User
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	var result model.User
	// UserIdを元にcolomnの取得
	constants.GetDB().Where("user_id = ?", loginVals.UserId).First(&result)

	fmt.Println("AUTHENTICATOR INPUT : " + loginVals.Password)

	fmt.Println("AUTHENTICATOR DB : " + result.Password)

	// パスワードハッシュの確認
	if !utils.CheckHashOfPassword(result.Password, loginVals.Password) {
		return nil, jwt.ErrFailedAuthentication
	}

	return &result, nil
}

// authorizator Tokenのユーザ情報から認証を行う
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(model.User); ok {
		var result model.User
		constants.GetDB().Where("user_id = ?", v.UserId).First(&result)
		//fmt.Println("AUTHORIZATOR INPUT : " + v.UserId)
		//fmt.Println("AUTHORIZATOR DB : " + result.UserId)
		if v.UserId == result.UserId {
			return true
		}
	}

	return false
}

// unauthorized returns the messagge of unauthorization
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// loginResponse builds the response of success full login
func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, gin.H{
		constants.EXPIRE: expire,
		constants.TOKEN:  token,
	})
}
