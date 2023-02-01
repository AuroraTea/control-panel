package cmdlet

import (
	"bytes"
	"control2/util/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type SetIPv4Request struct {
	NetAdapter string `json:"netAdapter" binding:"required"`
	IP         string `json:"ip" binding:"required,ipv4"`
	Mask       string `json:"mask" binding:"required,ipv4"`
	Gateway    string `json:"gateway" binding:"required,ipv4"`
}

func SetIPv4(c *gin.Context) {
	var body SetIPv4Request
	err := c.BindJSON(&body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	cmd := exec.Command("netsh",
		"interface", "ip", "set", "address", body.NetAdapter, "static", body.IP, body.Mask, body.Gateway)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		panic(encoding.ShouldAutoDecode(stdout.Bytes()))
	}
	c.String(http.StatusOK, body.NetAdapter+"'s 'IPv4 configuration changed successfully")
}
