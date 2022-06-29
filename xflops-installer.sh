#!/bin/bash

ansible-playbook -i invetory/hosts playbook/prepare.yaml
ansible-playbook -i invetory/hosts playbook/master.yaml
ansible-playbook -i invetory/hosts playbook/worker.yaml