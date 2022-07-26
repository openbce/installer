BCE_VER=v0.1.0
K8S_VER=v1.24.1
ETCD_VER=3.5.3-0
DNS_VER=v1.8.6
PAUSE_VER=3.7
IRONIC_VER=20.2.0
IRONIC_INSPECTOR_VER=10.12.0

REL_PKGS="invetory sbin playbook bce.yaml bce-installer"
REL_NAME=${BCE_VER}-linux-`uname -m`

docker-build: kube-apiserver kube-scheduler kube-controller-manager kube-proxy pause etcd coredns

kube-apiserver:
	docker build sbin -t openbce/kube-apiserver:${K8S_VER} -f docker/Dockerfile.kube-apiserver

kube-scheduler:
	docker build sbin -t openbce/kube-scheduler:${K8S_VER} -f docker/Dockerfile.kube-scheduler

kube-controller-manager:
	docker build sbin -t openbce/kube-controller-manager:${K8S_VER} -f docker/Dockerfile.kube-controller-manager

kube-proxy:
	docker build sbin -t openbce/kube-proxy:${K8S_VER} -f docker/Dockerfile.kube-proxy

kubeadm:
	docker build sbin -t openbce/kubeadm:${K8S_VER} -f docker/Dockerfile.kubeadm

kubelet:
	docker build sbin -t openbce/kubelet:${K8S_VER} -f docker/Dockerfile.kubelet

etcd:
	docker build sbin -t openbce/etcd:${ETCD_VER} -f docker/Dockerfile.etcd

coredns:
	docker build sbin -t openbce/coredns:${DNS_VER} -f docker/Dockerfile.coredns

pause:
	docker build sbin -t openbce/pause:${PAUSE_VER} -f docker/Dockerfile.pause

ironic:
	docker build . -t openbce/ironic:${IRONIC_VER} -f docker/Dockerfile.ironic
	docker build . -t openbce/ironic-inspector:${IRONIC_INSPECTOR_VER} -f docker/Dockerfile.ironic-inspector
	docker build . -t openbce/ironic-init:${BCE_VER} -f docker/Dockerfile.ironic-init

debug-tools:
	docker build . -t openbce/debug:${BCE_VER} -f docker/Dockerfile.debug

release:
	mkdir -p build/bce-${BCE_VER}
	for number in ${REL_PKGS} ; do \
        cp -R $$number build/bce-${BCE_VER}/ ; \
    done
	tar cvf build/bce-${REL_NAME}.tar -C build bce-${BCE_VER}
	gzip build/bce-${REL_NAME}.tar

clean:
	rm -rf build
