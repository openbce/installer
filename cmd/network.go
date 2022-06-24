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
	"os"
	"os/exec"
)

var modules = []string{"overlay", "nbr_netfilter"}

var sysctls = map[string]string{
	"net.bridge.bridge-nf-call-iptables":  "1",
	"net.bridge.bridge-nf-call-ip6tables": "1",
	"net.ipv4.ip_forward":                 "1",
}

func createNetworkConf(_ *XflopsConfiguration) error {
	modFile, err := os.OpenFile("/etc/modules-load.d/xflops.conf", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer modFile.Close()
	if err != nil {
		return err
	}

	for _, m := range modules {
		if _, err = modFile.WriteString(m + "\n"); err != nil {
			return err
		}
	}

	sysFile, err := os.OpenFile("/etc/sysctl.d/xflops.conf", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer sysFile.Close()
	if err != nil {
		return err
	}

	for k, v := range sysctls {
		if _, err = sysFile.WriteString(fmt.Sprintf("%s = %s\n", k, v)); err != nil {
			return err
		}
	}

	return nil
}

func setupNetwork(conf *XflopsConfiguration) error {
	if err := createNetworkConf(conf); err != nil {
		return err
	}

	for _, m := range modules {
		cmd := exec.Command("modprobe", m)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	cmd := exec.Command("sysctl", "--system")
	if err := cmd.Run(); err != nil {
		return err
	}

	// cat <<EOF | tee /etc/modules-load.d/xflops.conf
	//  overlay
	//  br_netfilter
	// EOF

	// sudo modprobe overlay
	// sudo modprobe br_netfilter

	// # sysctl params required by setup, params persist across reboots
	// cat <<EOF | tee /etc/sysctl.d/k8s.conf
	// net.bridge.bridge-nf-call-iptables  = 1
	// net.bridge.bridge-nf-call-ip6tables = 1
	// net.ipv4.ip_forward                 = 1
	// EOF
	// sysctl --system
	//

	return nil
}
