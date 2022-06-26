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

package app

import (
	"io"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"

	"xflops.cn/installer/pkg/api"
	"xflops.cn/installer/pkg/cmd"
)

// NewXFLOPSCommand returns cobra.Command to run xflops-installer command
func NewXFLOPSCommand(in io.Reader, out, err io.Writer, conf *api.XflopsConfiguration) *cobra.Command {
	var rootfsPath string

	cmds := &cobra.Command{
		Use:           "xflops-installer",
		Short:         "xflops-installer: easily bootstrap a secure Kubernetes cluster",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if rootfsPath != "" {
				if err := kubeadmutil.Chroot(rootfsPath); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmds.ResetFlags()

	cmds.AddCommand(cmd.NewInitCMD(out, conf))
	cmds.AddCommand(cmd.NewJoinCMD(out, nil))

	options.AddKubeadmOtherFlags(cmds.PersistentFlags(), &rootfsPath)

	return cmds
}
