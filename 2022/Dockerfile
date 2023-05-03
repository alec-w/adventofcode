FROM debian:11.6-slim

# Set default command
CMD ["sh"]

# Create a user that will be used when inisde the container rather than root
RUN useradd dev --create-home && \
    echo "dev:pass" | chpasswd && \
    adduser dev sudo

# bash-completion, gcc, make and vim are useful development tools
# curl is needed for installation of Go
# git is needed to run commands such as go mod tidy during development
# wget and unzip are needed for installation of proto tooling
RUN apt-get update && apt-get -y install \
    bash-completion \
    curl \
    gcc \
    git \
    make \
    vim \
    wget \
    unzip

# Install Go
ARG GO_VERSION=1.20.3
# Modify PATH to make go immediately available after installation
ENV PATH=$PATH:/usr/local/go/bin
RUN curl -sSfL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -C /usr/local -xzf -  && \
    go env -w GOBIN=/usr/local/go/bin GOROOT=/usr/local/go && \
    echo "export GOROOT=/usr/local/go" > /etc/profile.d/go.sh  && \
    echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh

# Install Go linter
ARG GOLANGCI_VERSION=1.52.2
RUN ["/bin/bash", "-c", "set -o pipefail && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOROOT)/bin v${GOLANGCI_VERSION}"]

# Instal goimport-reviser
ARG GOIMPORTS_REVISER_VERSION=3.3.1
RUN go install -v github.com/incu6us/goimports-reviser/v3@v${GOIMPORTS_REVISER_VERSION}

# Install mockery
ARG MOCKERY_VERSION=2.16.0
RUN ["/bin/bash", "-c", "set -o pipefail && wget -qO- https://github.com/vektra/mockery/releases/download/v${MOCKERY_VERSION}/mockery_${MOCKERY_VERSION}_Linux_x86_64.tar.gz | tar -xzv -C /usr/local/bin mockery"]

ENTRYPOINT ["bash", "-lec"]
