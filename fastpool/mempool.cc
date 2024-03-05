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
#include <vector>
#include <sstream>
#include <ctime>
#include <fstream>
#include <iostream>
#include <chrono>
#include <ctime>
#include <random>
#include <string>
#include <sstream>
#include <thread>

#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>

#include "rock.grpc.pb.h"

#include <filesystem>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

using rock::AddTxsRequest;
using rock::AddTxsResponse;
using rock::Mempool;

class MempoolClient
{
public:
  MempoolClient(std::string target)
  {
    auto channel = grpc::CreateChannel(target, grpc::InsecureChannelCredentials());
    stub_ = Mempool::NewStub(channel);
  }

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
      std::cout << "[mempool::ClientAddTxs] " << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }

  void SendBatch(rock::SendBatchRequest request)
  {
    ClientContext context;
    rock::AddTxsResponse reply;
    stub_->SendBatch(&context, request, &reply);
  }

private:
  std::unique_ptr<Mempool::Stub> stub_;
};

// 生成随机字符串
std::string generateRandomString(int length)
{
  const std::string charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  std::string result;
  result.resize(length);

  for (int i = 0; i < length; i++)
  {
    result[i] = charset[rand() % charset.length()];
  }

  return result;
}

std::string gen_batch_id()
{
  time_t now_c = time(0);
  tm *ltm = localtime(&now_c);

  auto now = std::chrono::system_clock::now();
  auto ms = std::chrono::time_point_cast<std::chrono::milliseconds>(now);
  auto epoch = ms.time_since_epoch();
  auto value = std::chrono::duration_cast<std::chrono::milliseconds>(epoch);

  // 格式化时间字符串
  std::stringstream ss;

  auto mon = std::to_string(1 + ltm->tm_mon);
  if (mon.length() == 1)
  {
    mon = "0" + mon;
  }

  auto day = std::to_string(ltm->tm_mday);
  if (day.length() == 1)
  {
    day = "0" + day;
  }

  ss << 1900 + ltm->tm_year << "_" << mon << day << "_"
     << ltm->tm_hour << ltm->tm_min << ltm->tm_sec << "_"
     << value.count() % 1000 << "_" // 毫秒时间戳
     << generateRandomString(4);

  std::string timeString = ss.str();

  return timeString;
}

void saveBatch(const std::string &dir_data, const rock::Batch &batch)
{
  std::string file_path = dir_data + "/" + batch.id() + ".binpb";
  std::ofstream output(file_path, std::ios::out | std::ios::trunc | std::ios::binary);
  if (!batch.SerializeToOstream(&output))
  {
    std::cerr << "[mempool::saveBatch] Failed to write batch." << std::endl;
  }
  else
  {
    std::cout << "[mempool::saveBatch] Save batch to " << file_path << std::endl;
  }
}

// Logic and data behind the server's behavior.
class MempoolServiceImpl final : public Mempool::Service
{
private:
  std::string dir_data;
  std::vector<MempoolClient> clients;

public:
  MempoolServiceImpl(std::string dir_data, std::vector<std::string> others)
  {
    this->dir_data = dir_data;

    clients = std::vector<MempoolClient>();
    for (auto &other : others)
    {
      clients.push_back(MempoolClient(other));
    }
  }

  Status AddTxs(ServerContext *context, const AddTxsRequest *request,
                AddTxsResponse *reply) override
  {
    std::string prefix("Hello ");
    // reply->set_message(prefix + request->name());
    std::cout << "[mempool::AddTxs] Received " << request->txs_size() << " transactions" << std::endl;

    rock::Batch batch;
    batch.set_id(gen_batch_id());
    batch.mutable_txs()->CopyFrom(request->txs());

    auto dir_data = this->dir_data;

    std::thread t_save(saveBatch, dir_data, batch);
    t_save.detach();

    auto batch_request = rock::SendBatchRequest();
    batch_request.mutable_batch()->CopyFrom(batch);

    for (int i = 0; i < clients.size(); i++)
    {
      std::thread t_send(&MempoolClient::SendBatch, &clients[i], batch_request);
      t_send.detach();
    }

    return Status::OK;
  }

  Status SendBatch(ServerContext *context, const rock::SendBatchRequest *request,
                   rock::AddTxsResponse *reply) override
  {
    std::cout << "[mempool::SendBatch] Received " << request->batch().txs_size() << " transactions" << std::endl;

    std::thread t_save(saveBatch, this->dir_data, request->batch());
    t_save.detach();

    return Status::OK;
  }
};

void RunServer(std::string host, std::vector<std::string> others, std::string dir_data)
{
  // dir_data 不存在则创建
  if (!std::filesystem::exists(dir_data))
  {
    std::filesystem::create_directories(dir_data);
  }

  MempoolServiceImpl service = MempoolServiceImpl(dir_data, others);

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;
  // Listen on the given address without any authentication mechanism.
  builder.AddListeningPort(host, grpc::InsecureServerCredentials());

  builder.SetMaxMessageSize(-1);

  // Register "service" as the instance through which we'll communicate with
  // clients. In this case it corresponds to an *synchronous* service.
  builder.RegisterService(&service);
  // Finally assemble the server.
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "[mempool::RunServer] Server listening on " << host << std::endl;

  // Wait for the server to shutdown. Note that some other thread must be
  // responsible for shutting down the server for this call to ever return.
  server->Wait();
}

// splitString 将字符串按指定分隔符分割成多个子字符串
std::vector<std::string> splitString(const std::string &str, char delimiter)
{
  std::vector<std::string> tokens;
  std::string token;
  std::istringstream tokenStream(str);

  while (std::getline(tokenStream, token, delimiter))
  {
    tokens.push_back(token);
  }

  return tokens;
}

int main(int argc, char **argv)
{
  char *host_cchar = std::getenv("HOST");
  std::string host;
  if (host_cchar == NULL)
  {
    host = "localhost:50051";
  }
  else
  {
    host = host_cchar;
  }

  char *others_cchar = std::getenv("OTHERS");
  std::string others;
  if (others_cchar == NULL)
  {
    others = "";
  }
  else
  {
    others = others_cchar;
    // e.g "localhost:50052,localhost:50053,localhost:50054"
  }

  char *dir_data_cchar = std::getenv("DIR_DATA");
  std::string dir_data;
  if (dir_data_cchar == NULL)
  {
    // dir_data = "./data/" + std::to_string(getpid());
    dir_data = "/tmp/extpool/" + std::to_string(getpid());
  }
  else
  {
    dir_data = dir_data_cchar;
  }

  std::cout << "[mempool::main] host = " << host << ", others = " << others << ", dir_data = " << dir_data << std::endl;

  auto others_vector = splitString(others, ',');

  std::cout << "[mempool::main] others_vector = ";
  for (size_t i = 0; i < others_vector.size(); i++)
  {
    if (i < others_vector.size() - 1)
    {
      std::cout << others_vector[i] << ", ";
    }
    else
    {
      std::cout << others_vector[i];
    }
  }
  std::cout << std::endl;

  RunServer(host, others_vector, dir_data);
  return 0;
}
