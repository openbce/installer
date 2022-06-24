#!/bin/bash

# install dependencies

XFLOPS_DEPS="ebtables socat conntrack containerd"

if type apt >/dev/null 2>&1; then 
  apt install -y ${XFLOPS_DEPS}
fi

if type yum >/dev/null 2>&1; then 
  yum install -y ${XFLOPS_DEPS}
fi

# Set up network environment

cat <<EOF | tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# Apply sysctl params without reboot
sysctl --system