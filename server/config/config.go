package config

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type GetConfigResponse struct {
	IsExist bool   `json:"isExist"`
	Config  string `json:"config"`
}

func GetConfig(c *gin.Context) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		c.JSON(http.StatusOK, &GetConfigResponse{
			IsExist: false,
		})
		return
	}
	c.JSON(http.StatusOK, &GetConfigResponse{
		IsExist: true,
		Config:  string(file),
	})
}

type SetConfigRequest struct {
	Config string `json:"config"`
}

func SetConfig(c *gin.Context) {
	var body SetConfigRequest
	err := c.BindJSON(&body)
	err = ioutil.WriteFile("config.json", []byte(body.Config), 0644)
	if err != nil {
		panic("Failed to write to config")
	}
	c.String(http.StatusOK, "")
}
