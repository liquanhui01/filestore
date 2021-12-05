// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/liquanhui01/filestore/internal/apiserver/store"
	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

// type Claims struct {
// 	Flag string
// 	jwt.StandardClaims
// }

// // GenerateToken generate jwt.
// func GenerateToken(userID uint) (string, error) {
// 	ex := viper.GetInt("jwt.expiredtime")
// 	flag := GenerateUserFlag(userID)
// 	expiredTime := time.Now().Add(time.Duration(ex) * time.Hour)

// 	claims := &Claims{
// 		Flag: flag,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expiredTime.Unix(), // expire time
// 			IssuedAt:  time.Now().Unix(),  // issue time
// 			Id:        flag,
// 			Issuer:    viper.GetString("jwt.issuer"),  // issuer
// 			NotBefore: time.Now().Unix(),              // effective time
// 			Subject:   viper.GetString("jwt.subject"), // jwt subject
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.salt")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// // GenerateUserFlag genrate only user flag.
// func GenerateUserFlag(userID uint) string {
// 	md5 := md5.New()
// 	md5.Write([]byte(strconv.Itoa(int(userID))))
// 	return hex.EncodeToString(md5.Sum([]byte("")))
// }

// // ParseToken: 解析token
// func ParseToken(token string) (*Claims, error) {
// 	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (
// 		interface{}, error) {
// 		return []byte(viper.GetString("jwt.salt")), nil
// 	})
// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
// 			return claims, nil
// 		}
// 	}
// 	return nil, err
// }

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func NewAuth() *jwt.GinJWTMiddleware {
	ginJWTMiddleware := &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(viper.GetString("jwt.salt")),
		Timeout:     time.Duration(viper.GetInt("jwt.timeout")) * time.Hour,
		MaxRefresh:  time.Duration(viper.GetInt("jwt.maxrefresh")) * time.Hour,
		PayloadFunc: payloadFunc(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		RefreshResponse: refreshResponse(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		LoginResponse:   loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
	}
	jwtMiddleware, _ := jwt.New(ginJWTMiddleware)

	return jwtMiddleware
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var login UserInfo
		var err error

		if c.Request.Header.Get("Authorization") != "" {
			login, err = parseWithHeader(c)
		} else {
			login, err = parseWithBody(c)
		}
		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		user, err := store.Client().Users().FindUserByName(c, login.Username)
		if err != nil {
			fmt.Printf("get user information failed: %s", err.Error())

			return "", jwt.ErrFailedAuthentication
		}

		if err = core.ValidatePassword(login.Password, user.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(string); ok {
			fmt.Printf("user `%s` is authenticated.", v)
			return true
		}

		return false
	}
}

func parseWithHeader(c *gin.Context) (UserInfo, error) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		fmt.Printf("get basic string from Authorization header failed")

		return UserInfo{}, jwt.ErrFailedAuthentication
	}

	payload, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		fmt.Printf("decode basic string: %s", err.Error())

		return UserInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		fmt.Println("parse payload failed")

		return UserInfo{}, jwt.ErrFailedAuthentication
	}

	return UserInfo{
		Username: pair[0],
		Password: pair[1],
	}, nil
}

func parseWithBody(c *gin.Context) (UserInfo, error) {
	var login UserInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		fmt.Printf("parse login parameters: %s", err.Error())

		return login, jwt.ErrFailedAuthentication
	}

	return login, nil
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"issuer": viper.GetString("jwt.issuer"),
			// "audience":
		}
		if u, ok := data.(*rp.User); ok {
			claims[jwt.IdentityKey] = u.Username
			claims["sub"] = u.Username
		}

		return claims
	}
}

func loginResponse() func(c *gin.Context, code int, message string, expire time.Time) {
	return func(c *gin.Context, code int, message string, expire time.Time) {
		fmt.Println("进入到登陆回复步骤")
		c.JSON(http.StatusOK, gin.H{
			"token":  message,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func refreshResponse() func(c *gin.Context, code int, message string, expire time.Time) {
	return func(c *gin.Context, code int, message string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  message,
			"expire": expire.Format(time.RFC3339),
		})
	}
}
