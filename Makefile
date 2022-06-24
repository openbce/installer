K8S_VER=v1.24.1
ETCD_VER=3.5.3-0
DNS_VER=v1.8.6
PAUSE_VER=3.7

xflops-installer:
	go build -o xflops-installer cmd/main.go

docker-build: kube-apiserver kube-scheduler kube-controller-manager kube-proxy pause etcd coredns

kube-apiserver:
	docker build sbin -t xflops/kube-apiserver:${K8S_VER} -f docker/Dockerfile.kube-apiserver

kube-scheduler:
	docker build sbin -t xflops/kube-scheduler:${K8S_VER} -f docker/Dockerfile.kube-scheduler

kube-controller-manager:
	docker build sbin -t xflops/kube-controller-manager:${K8S_VER} -f docker/Dockerfile.kube-controller-manager

kube-proxy:
	docker build sbin -t xflops/kube-proxy:${K8S_VER} -f docker/Dockerfile.kube-proxy

kubeadm:
	docker build sbin -t xflops/kubeadm:${K8S_VER} -f docker/Dockerfile.kubeadm

kubelet:
	docker build sbin -t xflops/kubelet:${K8S_VER} -f docker/Dockerfile.kubelet

etcd:
	docker build sbin -t xflops/etcd:${ETCD_VER} -f docker/Dockerfile.etcd

coredns:
	docker build sbin -t xflops/coredns:${DNS_VER} -f docker/Dockerfile.coredns

pause:
	docker build sbin -t xflops/pause:${PAUSE_VER} -f docker/Dockerfile.pause
