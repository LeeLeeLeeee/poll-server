package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/conf"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func CheckJwt() gin.HandlerFunc { /* JWT MiddleWare */
	return func(c *gin.Context) {
		tokenAuth, err := ExtractTokenMetadata(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized-meta")
			c.Abort()
			return
		}

		userId, err := FetchAuth(tokenAuth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized-redis")
			c.Abort()
			return
		}
		c.Set("userID", userId)
		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	bearToken, err := c.Cookie("access_token")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)

	if tokenString == "" {
		return nil, jwt.NewValidationError("not a token", jwt.ValidationErrorExpired)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		/*
			interface.(type)  => type assertion
			1. 첫 번째는 값, 두 번째는 type 존재 유무
			2. 두 개의 인자를 받을 땐 값이 없어도 panic이 발생하지 않음
			3. 값 인자만 전달 받을 경우 panic 발생
		*/
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	token, err := VerifyToken(c)

	if err != nil {
		if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorExpired != 0 {
			authD, err := RefreshToken(c)

			if err != nil {
				return nil, err
			}

			return authD, err
		} else {
			return nil, err
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)

		if !ok {
			return nil, err
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}

	return nil, err
}

func FetchAuth(authD *AccessDetails) (uint64, error) {
	client := conf.GetRedis().RedisClient
	userid, err := client.Get(authD.AccessUuid).Result()

	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)

	return userID, nil
}

func RefreshToken(c *gin.Context) (*AccessDetails, error) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		return nil, errors.New("refresh expired")
	}

	refreshKey := os.Getenv("JWT_REFRESH_SECRET")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(refreshKey), nil
	})

	if err != nil {
		return nil, errors.New("refresh token expired")
	}

	if !token.Valid {
		return nil, errors.New("refresh token error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, errors.New("refresh uuid error")
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			return nil, errors.New("userId error")
		}

		deleted, delErr := DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 {
			return nil, errors.New("unauthorized")
		}

		ts, createErr := createToken(userId)

		if createErr != nil {
			return nil, errors.New(createErr.Error())
		}

		saveErr := saveAuthKeyToRedis(userId, ts)

		if saveErr != nil {
			return nil, errors.New(saveErr.Error())
		}
		c.SetCookie("access_token", ts.AccessToken, int(time.Minute)*20, "/", "/localhost", false, false)
		c.SetCookie("refresh_token", ts.RefreshToken, int(time.Hour)*24*7, "/", "/localhost", false, true)

		return &AccessDetails{
			AccessUuid: ts.AccessUuid,
			UserId:     userId,
		}, nil
	} else {
		return nil, errors.New("refresh token error")
	}
}
