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
sudo bce-installer
```

## Configurations

```
bce_version: 0.1
k8s_version: v1.33.0
cri_version: v1.33.0
cilium_version: v1.17.5
certmanager_version: v1.18.2

work_dir: /opt/openbce
binary_dir: sbin
config_dir: config

pod_cidr: 172.88.0.0/16
service_cidr: 172.96.0.0/12
```
