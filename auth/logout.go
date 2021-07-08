package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/conf"
)

func LogOut(c *gin.Context) {
	au, err := ExtractTokenMetadata(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, delErr := DeleteAuth(au.AccessUuid)

	if delErr != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully loggedout")
}

func DeleteAuth(userId string) (int64, error) {
	client := conf.GetRedis().RedisClient
	deleted, err := client.Del(userId).Result()

	if err != nil {
		return 0, err
	}

	return deleted, nil
}
