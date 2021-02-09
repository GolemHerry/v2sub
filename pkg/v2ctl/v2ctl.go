package v2ctl

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func (v *v2Ray) Start() error {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, ":")

	var v2rayExecPath string

	if path, ok := v.flags["path"]; ok {
		v2rayExecPath = path
	} else {
		for _, v := range paths {
			if _, err := os.Open(v + "v2ray"); err == nil {
				v2rayExecPath = v
				break
			}
		}
	}
	if v2rayExecPath == "" {
		return errors.New("v2ray is not installed")
	}
	cmd := exec.Command("v2ray")
	cmd.CombinedOutput()

	return nil
}

func (v *v2Ray) Stop() error {
	return nil
}

func (v *v2Ray) Restart() error {
	return nil
}

func (v *v2Ray) AddFlag(key, value string) {
	v.flags[key] = value
}
