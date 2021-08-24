#!/bin/bash

# To add CNAME in /etc/hosts
cluster_ip=$(echo `minikube ip`)
echo "$cluster_ip config-service" | sudo tee -a /etc/hosts

# To deploy all resources
helmfile sync
