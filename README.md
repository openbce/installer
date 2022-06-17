# Installer

## Steps

### 1. Setup environment

`./bins` includes kubelet, kubectl, kubeadm and crictl binaries for the cluster, those binaries are used to install & manage
the cluster.

```
sudo cp ./bins/* /usr/bin
```

And then, install dependencies and setup ip forward by `release/setup.sh` (only support `apt/yum` right now)

```
sudo source release/setup.sh
```

Finally, apply kubelet & containerd configuration and start them

```
sudo cp release/kubelet.service
sudo systemctl start kubelet

sudo cp release/config.toml /etc/containerd/config.toml
sudo systemctl restart containerd
```

### 2. Execute kubeadm for master node

After seting up the environment, leverage `kubeadm` to install Kubernetes master as follow

```
sudo kubeadm init \
    --kubernetes-version=1.24.1 \
    --image-repository=docker.io/xflops \
    --pod-network-cidr=172.88.0.0/16 \
    --apiserver-advertise-address=10.210.0.113 \
    --service-cidr=172.96.0.0/12
```

### 3. Execute kubeadm for worker node

NOTES: For the worker node, it's also necessary to setup environment by setp-1

```
sudo kubeadm join 10.0.0.9:6443 --token xxxx \
        --discovery-token-ca-cert-hash xxx
```

### 4. Setup cluster network

```
kubectl apply -f release/calico.yaml
```

### 5. Install kubeflow/training-operator

```
kubectl apply -k "github.com/kubeflow/training-operator/manifests/overlays/standalone?ref=v1.4.0"
```

## FAQ

* If disk is full, update `/etc/containerd/config.toml` to assign a new directory for images & containers

    ```
    root="xxx"
    ```

