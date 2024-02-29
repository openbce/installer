#!/usr/bin/bash

kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/nmstate.io_nmstates.yaml
kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/namespace.yaml
kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/service_account.yaml
kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/role.yaml
kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/role_binding.yaml
kubectl apply -f https://github.com/nmstate/kubernetes-nmstate/releases/download/v0.80.1/operator.yaml


cat <<EOF | kubectl create -f -
apiVersion: nmstate.io/v1
kind: NMState
metadata:
  name: nmstate
EOF