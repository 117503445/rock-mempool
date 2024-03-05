/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

#include <iostream>
#include <memory>
#include <string>
#include <grpcpp/grpcpp.h>
#include "rock.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using rock::Mempool;
using rock::AddTxsRequest;
using rock::AddTxsResponse;

rock::Transaction genNTX()
{
  auto tx = rock::Transaction();
  tx.set_to("0x000000000000000000000000000000000000100c");
  tx.set_extradata("01d0439600000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005416c696365000000000000000000000000000000000000000000000000000000");
  tx.set_version("1");
  return tx;
}

class MempoolClient
{
public:
  MempoolClient(std::shared_ptr<Channel> channel)
      : stub_(Mempool::NewStub(channel)) {}

  // Assembles the client's payload, sends it and presents the response back
  // from the server.
  std::string AddTxs(AddTxsRequest request)
  {
    // Data we are sending to the server.
    // AddTxsRequest request;

    // Container for the data we expect from the server.
    AddTxsResponse reply;

    // Context for the client. It could be used to convey extra information to
    // the server and/or tweak certain RPC behaviors.
    ClientContext context;

    // The actual RPC.
    Status status = stub_->AddTxs(&context, request, &reply);

    // Act upon its status.
    if (status.ok())
    {
      // return reply.status();
      return "";
    }
    else
    {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }

private:
  std::unique_ptr<Mempool::Stub> stub_;
};

int main(int argc, char **argv)
{
  // Instantiate the client. It requires a channel, out of which the actual RPCs
  // are created. This channel models a connection to an endpoint specified by
  // the argument "--target=" which is the only expected argument.
  std::string target_str = "localhost:50051";
  // We indicate that the channel isn't authenticated (use of
  // InsecureChannelCredentials()).
  MempoolClient Mempool(
      grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));
  std::string user("world");

  auto request = AddTxsRequest();
  for (int i = 0; i < 100000; i++)
  {
    request.mutable_txs()->Add(genNTX());
  }
  for (int i = 0; i < 50; i++)
  {
    // 计时
    auto now = std::chrono::system_clock::now();
    std::string reply = Mempool.AddTxs(request); // The actual RPC call!
    std::cout << "Mempool received: " << reply << std::endl;
    auto end = std::chrono::system_clock::now();
    std::chrono::duration<double> elapsed_seconds = end - now;
    std::cout << "elapsed time: " << elapsed_seconds.count() << "s\n";
  }

  return 0;
}
