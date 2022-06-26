/*
Copyright 2022 The XFLOPS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prepare

import (
	"fmt"
	"os"
	"os/exec"
	"xflops.cn/installer/pkg/api"
)

var systemds = []string{"containerd", "kubelet"}

func reloadSystemd() error {
	for _, s := range systemds {
		cmd := exec.Command("systemctl", "restart", s)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to run restart <%s>: %v", s, err)
		}
	}

	return nil
}

func setEnv(conf *api.XflopsConfiguration) error {
	path := os.Getenv("PATH")
	path = fmt.Sprintf("%s/bin:%s/sbin:%s", conf.HomeDir, conf.HomeDir, path)
	if err := os.Setenv("PATH", path); err != nil {
		return fmt.Errorf("failed to update PATH env: %v", err)
	}

	if err := os.Setenv("XFLOPS_HOME", conf.HomeDir); err != nil {
		return fmt.Errorf("failed to set XFLOPS_HOME env: %v", err)
	}

	if err := os.Setenv("XFLOPS_WORK_DIR", conf.WorkDir); err != nil {
		return fmt.Errorf("failed to set XFLOPS_WORK_DIR env: %v", err)
	}

	return nil
}
