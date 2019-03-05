FROM ubuntu:18.10 as foundation

ARG UID=1000
ARG TERRAFORM_VERSION="0.11.11"

RUN apt-get update && apt-get -y install \
    ansible \
    make \
    sudo \
    unzip \
    virtualenv

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && mv terraform /bin/terraform \
    && /bin/terraform version

# Set the Docker image users UID to the build-arg values. Since we build this as part of using the development
# toolchain the values should always be correct for any given developers machine.
RUN useradd dev --uid ${UID} -U \
    && echo "dev ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/dev \
    && chmod 0440 /etc/sudoers.d/dev \
    && ln -sf /usr/bin/python3 /usr/bin/python

USER dev
