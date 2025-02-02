FROM golang:1.17

ENV https_proxy=http://10.0.0.1:1087
RUN go get github.com/tcnksm/ghr && \
    mkdir -p /go/src/k8s.io && \
    git clone https://github.com/kubernetes/code-generator /go/src/k8s.io/code-generator && \
    cd /go/src/k8s.io/code-generator && \
    git checkout 1eeed5f600b70f788fa97951e1e7b47ce212c242 && \
    go mod init && go mod vendor && \
    go build -o /go/bin/conversion-gen ./cmd/conversion-gen

FROM golang:1.17

ENV DOCKER_VERSION "17.03.0-ce"

# The following is based on https://github.com/kubernetes/release/blob/master/debian/Dockerfile

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y && \
    apt-get -yy -q install --no-install-recommends --no-install-suggests --fix-missing \
      dpkg-dev build-essential debhelper socat \
      ca-certificates curl && \
    apt-get upgrade -y && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -sSL -o "/tmp/docker-${DOCKER_VERSION}.tgz" "https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION}.tgz" && \
    tar -xz -C /tmp -f "/tmp/docker-${DOCKER_VERSION}.tgz" && \
    mv /tmp/docker/* /usr/bin

COPY --from=0 /go/bin/* /go/bin/

ADD . /go/src/github.com/nxsre/criproxy
WORKDIR /go/src/github.com/nxsre/criproxy

CMD ["./build-package.sh"]
