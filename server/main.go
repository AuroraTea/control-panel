package main

import (
	"control2/cmdlet"
	"control2/config"
	"control2/static"
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(gin.Recovery())
	r.Use(CORS())
	if err != nil {
		panic(err)
	}

	// Static Web Files in '/' Path, Optional
	r.Use(static.Web())
	r.NoRoute(static.VueHistoryRouter)

	api := r.Group("api")

	api.GET("/config", config.GetConfig)
	api.PUT("/config", config.PutConfig)

	// TODO: rename API
	api.GET("/network-interfaces", cmdlet.GetNetAdapter)
	api.PUT("/network-interface/:networkInterface/ipv4", cmdlet.PutNetworkInterfaceIP)
	api.PUT("/computer-name", cmdlet.PutComputerName)

	go func() {
		err := exec.Command(`cmd`, `/c`, `start`, `http://localhost:5222/`).Run()
		if err != nil {
			panic(err)
		}
	}()

	_ = r.Run("127.0.0.1:5222")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, Access-Token, Content-Type, X-Requested-With, XMLHttpRequest")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.String(http.StatusOK, "Options")
		}
		c.Next()
	}
}
