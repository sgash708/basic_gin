FROM ubuntu:latest

RUN apt update \
    && apt install -y wget curl unzip make

RUN add-apt-repository ppa:longsleep/golang-backports \
    && apt install -y golang

ENV TZ=Asia/Tokyo

ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /home/go