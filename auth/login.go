package auth

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/conf"
	"github.com/twinj/uuid"
)

type loginInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = loginInfo{
	Username: "username",
	Password: "password",
}

func LogIn(c *gin.Context) {
	var u loginInfo

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json")
		return
	}

	if !checkUserInfo(u) {
		c.JSON(http.StatusUnauthorized, "Login is faild")
		return
	}

	token, err := createToken(u.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := saveAuthKeyToRedis(u.ID, token)

	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	resToken := map[string]string{
		"access_token": token.AccessToken,
		"refesh_token": token.RefreshToken,
	}

	log.Println(token)

	c.JSON(http.StatusOK, resToken)

}

func checkUserInfo(u loginInfo) bool {
	if user.Username != u.Username || user.Password != u.Password {
		return false
	}

	return true
}

func createToken(userid uint64) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	secretKey := os.Getenv("JWT_ACCESS_SECRET")
	refreshKey := os.Getenv("JWT_REFRESH_SECRET")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshToken
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshKey))

	if err != nil {
		return nil, err
	}

	return td, nil

}

func saveAuthKeyToRedis(userid uint64, td *TokenDetail) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	client := conf.GetRedis().RedisClient

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
