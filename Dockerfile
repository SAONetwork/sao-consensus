FROM ubuntu:20.04

ENV GO_VERSION=1.19.2
ENV IGNITE_VERSION=0.25.1
ENV NODE_VERSION=18.x

ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

RUN apt-get update -y
RUN apt update && apt install -y build-essential clang curl gcc jq wget zsh net-tools git

# Install Go
#RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C $LOCAL -xzf -
RUN curl -L https://studygolang.com/dl/golang/go${GO_VERSION}.linux-amd64.tar.gz | tar -C $LOCAL -xzf -

# Install Node
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

RUN mkdir -p /sao-consensus
ADD . /sao-consensus
RUN cd /sao-consensus && make

WORKDIR /sao-consensus
VOLUME /var/lib/sao-consensus
EXPOSE 1317 3000 4500 5000 26657

COPY --from=builder $GOPATH/bin/saod /usr/local/bin/

ENTRYPOINT ["/bin/bash"]
