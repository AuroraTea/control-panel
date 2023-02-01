package cmdlet

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func SetIPv4(c *gin.Context) {
	name := c.DefaultQuery("name", c.Param("networkInterface"))

	var body map[string]any
	err := c.BindJSON(&body)
	if err != nil {
		c.String(http.StatusBadRequest, "参数格式错误")
		return
	}
	ip := c.DefaultQuery("ip", body["ip"].(string))
	mask := c.DefaultQuery("mask", body["mask"].(string))
	gateway := c.DefaultQuery("gateway", body["gateway"].(string))
	if name == "" || gateway == "10.210.." || mask == "" || ip == "10.210..254" {
		c.String(http.StatusBadRequest, "错误: 参数不足")
		return
	}

	cmd := exec.Command("netsh", "interface", "ip", "set", "address", name, "static", ip, mask, gateway)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		panic(stdout.String())
	}
	c.String(http.StatusOK, "网卡"+name+"的ipv4配置修改成功")
}
