package cmdlet

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

// GetNetAdapter 列出所有网卡的名字(可供其他shell命令使用的名字)
func GetNetAdapter(c *gin.Context) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var names []string
	for _, i := range interfaces {
		names = append(names, i.Name)
	}
	c.JSON(http.StatusOK, names)
}
