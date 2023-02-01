package cmdlet

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type SetIPv4Request struct {
	NetAdapter string `json:"netAdapter"`
	IP         string `json:"ip"`
	Mask       string `json:"mask"`
	Gateway    string `json:"gateway"`
}

func SetIPv4(c *gin.Context) {
	var body SetIPv4Request
	err := c.BindJSON(&body)
	if err != nil {
		c.String(http.StatusBadRequest, "参数格式错误")
		return
	}

	cmd := exec.Command("netsh",
		"interface", "ip", "set", "address", body.NetAdapter, "static", body.IP, body.Mask, body.Gateway)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		panic(stdout.String())
	}
	c.String(http.StatusOK, "网卡"+body.NetAdapter+"的ipv4配置修改成功")
}
