# Installer

## Steps

### Update `hosts` list accordingly

```
[masters]
10.0.0.10

[workers]
10.0.0.11
```

### Execute installer command on bootstrap nodes

```
sudo cbe-installer
```

## Configurations

```
xflops-version: 1.0
kubernetes-version: 1.24.1
work_dir: /scrap/users/klausm
image-repository: docker.io/xflops
pod-network-cidr: 172.88.0.0/16
apiserver-advertise-address: 10.210.0.125
service-cidr: 172.96.0.0/12
```