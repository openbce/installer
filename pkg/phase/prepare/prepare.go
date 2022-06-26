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

	"xflops.cn/installer/pkg/api"
)

func Execute(conf *api.XflopsConfiguration) error {
	// Create work dir if not exist
	if err := os.MkdirAll(conf.WorkDir, 0644); err != nil {
		return fmt.Errorf("failed to create work dir at <%s>: %v\n", conf.WorkDir, err)
	}

	if err := setupNetwork(conf); err != nil {
		return fmt.Errorf("failed to setup network: %v\n", err)
	}

	if err := generateFiles(conf); err != nil {
		return fmt.Errorf("failed to generate configuration files: %v\n", err)
	}

	if err := reloadSystemd(); err != nil {
		return fmt.Errorf("failed to reload systemd service: %v\n", err)
	}

	if err := setEnv(conf); err != nil {
		return fmt.Errorf("failed to set envs: %v", err)
	}

	return nil
}
