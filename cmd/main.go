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

package main

import (
	"fmt"

	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"

	"xflops.cn/installer/cmd/app"
	"xflops.cn/installer/pkg/api"
	"xflops.cn/installer/pkg/phase/prepare"
)

func main() {
	//	Read xflops.conf
	conf := api.LoadConfiguration()

	if err := prepare.Execute(conf); err != nil {
		fmt.Printf("Failed to setup environment: %v\n", err)
		return
	}

	kubeadmutil.CheckErr(app.Run(conf))
}
