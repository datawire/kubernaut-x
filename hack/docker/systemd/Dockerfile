FROM ubuntu:18.10

# This is set for Systemd to identify where it is running
#
# Reference: https://www.freedesktop.org/wiki/Software/systemd/ContainerInterface/ (see: Environment Variables)
ENV container=docker

# This is set for Systemd to cleanly shutdown
#
# Reference: https://www.freedesktop.org/wiki/Software/systemd/ContainerInterface/ (see: Advanced Integration)
STOPSIGNAL SIGRTMIN+3

RUN find /etc/systemd/system \
         /lib/systemd/system \
         -path '*.wants/*' \
         -not -name '*journald*' \
         -not -name '*systemd-tmpfiles*' \
         -not -name '*systemd-user-sessions*' \
         -exec rm \{} \;

RUN apt-get update && apt-get install -y \
        ca-certificates \
        dbus \
        openssl \
        sudo \
        systemd \
        tzdata \
    && rm -rf /var/lib/apt/lists/*

RUN useradd dev \
    && echo "dev ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/dev \
    && chmod 0440 /etc/sudoers.d/dev \
    && ln -sf /usr/bin/python3 /usr/bin/python

RUN systemctl set-default multi-user.target
RUN systemctl mask \
        auditd.service \
        console-getty.service \
        dev-hugepages.mount \
        docker-storage-setup.service \
        getty.target \
        lvm2-lvmetad.service \
        sys-fs-fuse-connections.mount \
        systemd-logind.service \
        systemd-remount-fs.service \
        systemd-udev-hwdb-update.service \
        systemd-udev-trigger.service \
        systemd-udevd.service \
        systemd-vconsole-setup.service;

# Workaround for docker/docker#27202, technique based on comments from docker/docker#9212
CMD ["/bin/bash", "-c", "exec /sbin/init --log-target=journal 3>&1"]
