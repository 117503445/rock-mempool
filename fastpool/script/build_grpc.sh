#!/usr/bin/env bash

set -ex

export MY_INSTALL_DIR=$HOME/.local
mkdir -p $MY_INSTALL_DIR

cd /dep/grpc
mkdir -p cmake/build
pushd cmake/build
# cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR -DgRPC_SSL_PROVIDER=package -DOPENSSL_INCLUDE_DIR="/root/.hunter/_Base/dc49367/2548d81/913ac48/Install/include/openssl" -DOPENSSL_SSL_LIBRARY="/root/.hunter/_Base/dc49367/2548d81/913ac48/Install/lib/libssl.a" -DOPENSSL_CRYPTO_LIBRARY="/root/.hunter/_Base/dc49367/2548d81/913ac48/Install/lib/libcrypto.a" ../..

# cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR -DOPENSSL_ROOT_DIR=/root/.hunter/_Base/dc49367/2548d81/913ac48/Install/include/openssl -DgRPC_SSL_PROVIDER=package ../..

# cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR -DgRPC_SSL_PROVIDER=package ../..


# cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR ../..

# package: 使用系统的openssl, module: 使用grpc自带的openssl
# -DgRPC_SSL_PROVIDER=package 

# -DBORINGSSL_PREFIX=grpc

cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR -DgRPC_SSL_PROVIDER=module  ../..

if [ `free -g | grep Mem | awk '{print $2}'` -gt 64 ]; then
  make -j
else
  make -j4
fi
make install
popd

# /root/.hunter/_Base/dc49367/2548d81/913ac48/Install/lib/libcrypto.a