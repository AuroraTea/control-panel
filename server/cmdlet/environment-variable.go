package cmdlet

import (
	"bytes"
	"control2/util/encoding"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
)

func ListSystemPathEnvironmentVariables(c *gin.Context) {
	cmd := exec.Command("chcp", "65001", "whoami", "/user")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(string(encoding.ShouldAutoDecode(stderr.Bytes())))
	}
	fmt.Println(string(encoding.ShouldAutoDecode(stdout.Bytes())))
	stdoutString, _ := simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	stdoutStringSlice := strings.Fields(stdoutString)
	fmt.Println(stdoutStringSlice[len(stdoutStringSlice)-1])

	cmd = exec.Command("reg", "query", `HKEY_USERS\`+stdoutStringSlice[len(stdoutStringSlice)-1]+`\Environment`, "/v", "Path")
	stdout.Reset()
	stderr.Reset()
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(string(encoding.ShouldAutoDecode(stderr.Bytes())))
	}
	stdoutString, _ = simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	oldUserPath := strings.Join(strings.Fields(stdoutString)[3:len(strings.Fields(stdoutString))], "")
	fmt.Println(oldUserPath)
}

func ListUserPathEnvironmentVariables(c *gin.Context) {
	cmd := exec.Command("chcp", "65001", "whoami", "/user")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		//fmt.Println(string(ShouldAutoDecode(stderr.Bytes())))
	}
	fmt.Println(string(encoding.ShouldAutoDecode(stdout.Bytes())))
	stdoutString, _ := simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	stdoutStringSlice := strings.Fields(stdoutString)
	fmt.Println(stdoutStringSlice[len(stdoutStringSlice)-1])

	cmd = exec.Command("reg", "query", `HKEY_USERS\`+stdoutStringSlice[len(stdoutStringSlice)-1]+`\Environment`, "/v", "Path")
	stdout.Reset()
	stderr.Reset()
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(string(encoding.ShouldAutoDecode(stderr.Bytes())))
	}
	stdoutString, _ = simplifiedchinese.GBK.NewDecoder().String(stdout.String())
	oldUserPath := strings.Join(strings.Fields(stdoutString)[3:len(strings.Fields(stdoutString))], "")
	fmt.Println(oldUserPath)
}

func ListSystemEnvironmentVariables(c *gin.Context) {

}
