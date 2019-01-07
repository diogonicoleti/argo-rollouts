FROM golang:1.10.3 as builder

RUN apt-get update && apt-get install -y \
    git \
    make \
    wget \
    gcc \
    zip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install dep
ENV DEP_VERSION=0.5.0
RUN wget https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -O /usr/local/bin/dep && \
    chmod +x /usr/local/bin/dep


# A dummy directory is created under $GOPATH/src/dummy so we are able to use dep
# to install all the packages of our dep lock file
COPY Gopkg.toml ${GOPATH}/src/dummy/Gopkg.toml
COPY Gopkg.lock ${GOPATH}/src/dummy/Gopkg.lock

RUN cd ${GOPATH}/src/dummy && \
    dep ensure -vendor-only && \
    mv vendor/* ${GOPATH}/src/ && \
    rmdir vendor

# Perform the build
WORKDIR /go/src/github.com/argoproj/rollout-controller
COPY . .
ARG MAKE_TARGET="controller"
RUN make ${MAKE_TARGET}

FROM debian:9.5-slim

COPY --from=builder /go/src/github.com/argoproj/rollout-controller/dist/rollouts-controller /bin/

RUN groupadd -g 999 rollout-controller && \
    useradd -r -u 999 -g rollout-controller rollout-controller && \
    mkdir -p /home/rollout-controller && \
    chown rollout-controller:rollout-controller /home/rollout-controller


USER rollout-controller

WORKDIR /home/rollout-controller

ENTRYPOINT [ "/bin/rollouts-controller" ]