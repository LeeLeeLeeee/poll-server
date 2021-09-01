package auth

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/conf"
	"github.com/leeleeleeee/web-app/lib"
	"github.com/leeleeleeee/web-app/model"
	"github.com/twinj/uuid"
)

type LoginInfo struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

func LogIn(c *gin.Context) {
	var u *LoginInfo

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json")
		return
	}

	var ok bool
	var userId uint64

	if ok, userId = checkUserInfo(u); !ok {
		c.JSON(http.StatusUnauthorized, "Login is faild")
		return
	}

	token, err := createToken(userId)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := saveAuthKeyToRedis(userId, token)

	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	c.SetCookie("access_token", token.AccessToken, int(time.Minute)*20, "/", "/localhost", false, false)
	c.SetCookie("refresh_token", token.RefreshToken, int(time.Hour)*24*7, "/", "/localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"userid": userId,
	})

}

func checkUserInfo(u *LoginInfo) (ok bool, userId uint64) {
	qs := model.UserQuerySet{}
	var encryptErr error
	u.Password, encryptErr = lib.EncryptSha256(u.Password)

	if encryptErr != nil {
		return false, 0
	}

	uid, err := qs.CheckLogin(u.UserId, u.Password)

	if err != nil {
		return false, 0
	}

	return true, uid
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
	rtClaims["refresh_uuid"] = td.RefreshUuid
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
