package cmdlet

import (
	"bytes"
	"control2/util/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

// PutComputerName 修改计算机名
func PutComputerName(c *gin.Context) {
	var body map[string]any
	err := c.BindJSON(&body)
	if err != nil {
		c.String(http.StatusBadRequest, "参数格式错误")
		return
	}
	newName := c.DefaultQuery("newName", body["newName"].(string))

	cmd := exec.Command("wmic")
	cmdLine := `computersystem where name="$computername" rename ` + newName
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + os.ExpandEnv(cmdLine)}
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(encoding.ShouldAutoDecode(out))
	}
	if bytes.Equal(regexp.MustCompile(`ReturnValue\s*=\s*(\d+)`).FindSubmatch(out)[0], []byte("ReturnValue = 5")) {
		c.String(http.StatusInternalServerError, "计算机名修改失败, 可以考虑手动以管理员权限运行")
		return
	}
	if !bytes.Equal(regexp.MustCompile(`ReturnValue\s*=\s*(\d+)`).FindSubmatch(out)[0], []byte("ReturnValue = 0")) {
		c.String(http.StatusInternalServerError, "未知错误"+string(encoding.ShouldAutoDecode(out)))
		return
	}
	c.String(http.StatusOK, "计算机名修改成功")
}
