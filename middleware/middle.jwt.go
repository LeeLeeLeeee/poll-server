package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		tokenAuth, err := ExtractTokenMetadata(c.Request)
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

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

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

func CheckTokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return nil, err
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
