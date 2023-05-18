package cmdlet

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	//out, err := exec.Command("powershell", `chcp 65001; whoami /user`).CombinedOutput()
	//if err != nil {
	//	fmt.Println(string(out))
	//	return
	//}
	//outSlice := strings.Fields(string(out))
	//sid := outSlice[len(outSlice)-1]
	out, err := exec.Command("powershell", `reg query "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v path`).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return
	}

	oldUserPath := strings.Split(strings.Join(strings.Fields(string(out))[4:], ""), ";")

	fmt.Println(oldUserPath)

}
