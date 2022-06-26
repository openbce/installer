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
	"html/template"
	"os"
	"path/filepath"

	"xflops.cn/installer/pkg/api"
)

var templates = map[string]string{
	"templates/containerd.toml": "/etc/containerd/config.toml",
	"templates/kubelet.service": "/etc/systemd/system/kubelet.service",
	"templates/xflops_envs.csh": "env/xflops.csh",
	"templates/xflops_envs.sh":  "env/xflops.sh",
}

func generateFiles(conf *api.XflopsConfiguration) error {
	for k, v := range templates {
		t1, err := template.ParseFiles(k)
		if err != nil {
			return err
		}

		if err = func() error {
			if err := os.MkdirAll(filepath.Dir(v), 0644); err != nil {
				return fmt.Errorf("failed to create parent directory of %s", v)
			}
			f, err := os.OpenFile(v, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			defer f.Close()
			if err != nil {
				return err
			}

			t1.Execute(f, &conf)

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}