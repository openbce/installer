K8S_VER=v1.24.1
ETCD_VER=3.5.3-0
DNS_VER=v1.8.6
PAUSE_VER=3.7


all: kube-apiserver kube-scheduler kube-controller-manager kube-proxy pause etcd coredns


kube-apiserver:
	docker build bins -t xflops/kube-apiserver:${K8S_VER} -f docker/Dockerfile.kube-apiserver

kube-scheduler:
	docker build bins -t xflops/kube-scheduler:${K8S_VER} -f docker/Dockerfile.kube-scheduler

kube-controller-manager:
	docker build bins -t xflops/kube-controller-manager:${K8S_VER} -f docker/Dockerfile.kube-controller-manager

kube-proxy:
	docker build bins -t xflops/kube-proxy:${K8S_VER} -f docker/Dockerfile.kube-proxy

etcd:
	docker build bins -t xflops/etcd:${ETCD_VER} -f docker/Dockerfile.etcd

coredns:
	docker build bins -t xflops/coredns:${DNS_VER} -f docker/Dockerfile.coredns

pause:
	docker build bins -t xflops/pause:${PAUSE_VER} -f docker/Dockerfile.pause

