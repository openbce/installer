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

type XflopsConfiguration struct {
	WorkDir           string `yaml:"work-dir"`
	HomeDir           string `yaml:"home-dir"`
	Repository        string `yaml:"image-repository"`
	PodCIDR           string `yaml:"pod-network-cidr"`
	APIAddress        string `yaml:"apiserver-advertise-address"`
	ServiceCIDR       string `yaml:"service-cidr"`
	KubernetesVersion string `yaml:"kubernetes-version"`
}
