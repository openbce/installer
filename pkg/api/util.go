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

package api

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfiguration() *XflopsConfiguration {
	conf := &XflopsConfiguration{}

	if y, err := os.ReadFile("xflops.yaml"); err != nil {
		return nil
	} else {
		yaml.Unmarshal(y, conf)
	}

	if conf.HomeDir == "" {
		if wd, err := os.Getwd(); err != nil {
			return nil
		} else {
			conf.HomeDir = wd
		}
	}

	return conf
}
