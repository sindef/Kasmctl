FROM ubuntu:focal-20220531

RUN mkdir -p /tmp/build

#Add our dependencies and build essentials
RUN apt-get update && apt-get install -y \
    build-essential \
    libffi-dev \
    make \
    wget
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

#Install Golang
RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

COPY . /tmp/build

##Complete tests
WORKDIR /tmp/build
RUN export PATH=$PATH:/usr/local/go/bin && make tests && make build

##Check if the binary is in /usr/local/bin/kasmctl
RUN if [ -f /usr/local/bin/kasmctl ]; then echo "kasmctl is in /usr/local/bin/kasmctl"; else exit 1; fi

RUN /usr/local/bin/kasmctl version
