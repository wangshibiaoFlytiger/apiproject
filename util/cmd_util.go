package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

/**
执行系统命令
*/
func ExecCmd(name string, params ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, params...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err, stderr.String())
	}

	fmt.Println("执行系统命令完成", name, params)
	return stdout.String(), stderr.String(), err
}
