#!/bin/bash

function wait_for_ingress () {
    while [[ -z $(kubectl get service $1 -n $2 -o jsonpath="{.status.loadBalancer.ingress}" 2>/dev/null) ]]; do
        echo "INFO: waiting for $1/$2 to get ingress"
        sleep 3
	done
}

wait_for_ingress ipxe-svc ironic-system

export IPXE_HOST_IP=$(kubectl get svc ipxe-svc -n ironic-system -o json | jq -r ".spec.clusterIP")
export IPXE_PUBLIC_IP=$(kubectl get svc ipxe-svc -n ironic-system -o json | jq -r ".status.loadBalancer.ingress[0].ip")
export IPXE_HTTP_PORT=8080

export IRONIC_HOST_IP=$(hostname -i)

export IRONIC_HOME=/opt/ironic

SCRIPT_HOME=${IRONIC_HOME}/scripts

function render_j2_config() {
	python3 -c 'import os; import sys; import jinja2; sys.stdout.write(jinja2.Template(sys.stdin.read()).render(env=os.environ))' < $1
}

function init_conf_debug() {
	echo "---------------------------------------------------------------------------------------------------"
	echo "------------------------------------------ DEBUG MESSAGE ------------------------------------------"
	echo "---------------------------------------------------------------------------------------------------"
	cat /etc/ironic/ironic.conf
	echo "---------------------------------------------------------------------------------------------------"
	cat /etc/ironic/ironic-inspector.conf
}

function setup_mariadb() {
	# Grant all premission to ironic user.
	render_j2_config ${SCRIPT_HOME}/ironic.sql.j2 | mysql -h mariadb-svc.mariadb-system.svc.cluster.local -u root -p${MARIADB_ROOT_PASSWORD}

	# Create tables for ironic.
	ironic-dbsync 
}

function setup_images_dirs() {
	mkdir -p ${IRONIC_HOME}/images/tftpboot/master_images
	mkdir -p ${IRONIC_HOME}/images/httpboot/master_images
	mkdir -p ${IRONIC_HOME}/images/cache
	mkdir -p ${IRONIC_HOME}/images/tmp
}

# Generate ironic configuration
render_j2_config ${SCRIPT_HOME}/ironic.conf.j2 > /etc/ironic/ironic.conf 
render_j2_config ${SCRIPT_HOME}/ironic-inspector.conf.j2 > /etc/ironic/ironic-inspector.conf 

# Sync up database
setup_mariadb

# Create dirs for boot images
setup_images_dirs

# Print debug message
init_conf_debug

