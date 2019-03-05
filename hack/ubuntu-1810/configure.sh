#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
set -o verbose

export DEBIAN_FRONTEND=noninteractive
export DEBIAN_DOCKER_VERSION="18.06.1~ce~3-0~ubuntu"
export DEBIAN_KUBERNETES_VERSION=""

# ======================================
# Purge Snapd
# ======================================
# "snapd" is the daemon component of the Snap application format. It is not
# needed for Kubernaut. If you're curious see the below URL for more info.
#
# Snapd Project: https://github.com/snapcore/snapd
apt-get -y autoremove --purge snapd

# ======================================
# System Upgrade
# ======================================
apt-get update
apt-get \
    -y \
    -o Dpkg::Options::="--force-confdef" \
    -o Dpkg::Options::="--force-confold" \
    dist-upgrade

apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    git \
    software-properties-common

rm -rf /var/lib/apt/lists/*

# ======================================
# Docker Installation
# ======================================

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

apt-get update
apt-get -y install docker-ce=${DEBIAN_DOCKER_VERSION}

rm -rf /var/lib/apt/lists/*

systemctl stop docker
systemctl enable docker

# ======================================
# Kubernetes Installation
# ======================================

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF

apt-get update
apt-get -y install \
	kubelet=${DEBIAN_KUBERNETES_VERSION} \
	kubeadm=${DEBIAN_KUBERNETES_VERSION} \
	kubectl=${DEBIAN_KUBERNETES_VERSION}

rm -rf /var/lib/apt/lists/*

systemctl stop kubelet
systemctl disable kubelet

# ======================================
# Finalization
# ======================================