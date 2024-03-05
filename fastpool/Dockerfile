FROM ubuntu:22.04

COPY script/sources.list /etc/apt/sources.list

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get update && apt-get install -y cmake autoconf libtool git python3 gcc-12 g++-12 build-essential ninja-build curl libfmt-dev psmisc fish gdb pkg-config wget
RUN update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-12 60 --slave /usr/bin/g++ g++ /usr/bin/g++-12

# install grpc
WORKDIR /dep
RUN git clone --recurse-submodules -b v1.61.0 --depth 1 --shallow-submodules https://github.com/grpc/grpc
COPY script/build_grpc.sh ./grpc
RUN ./grpc/build_grpc.sh

COPY script/config.fish /root/.config/fish/config.fish
RUN chsh -s $(which fish)

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN mkdir -p /workspace

WORKDIR /workspace