
# Klever Challenge

 API with Golang using gRPC with stream pipes that exposes an upvote service endpoints:

- [x] The API must guarantee the typing of user inputs.

- [x] If an input is expected as a string, it can only be received as a string.

- [x] The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct

- [ ] The API should contain unit test of methods it uses

- [ ] Deliver the whole solution running in some free cloud service



## Acknowledgements

 - [gRPC Doc](https://grpc.io/)
 - [Mongo Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
 - [Learning Go: MongoDB CRUD with gRPC](https://itnext.io/learning-go-mongodb-crud-with-grpc-98e425aeaae6)
 - [Eric Lau](https://www.youtube.com/watch?v=f6KG5eqOPFw)

## Run Locally

Assuming that you've already mongodb installed and a gRPC client server for method calls.

- [Evans](https://evans.syfm.me/about)
- [Postman gRPC Client](https://blog.postman.com/postman-now-supports-grpc/)
- [MongoDb Installation Guide](https://www.mongodb.com/docs/guides/server/install/)

Also you'll need `protoc` installed in your OS.

- [Protoc Installation Guide](https://grpc.io/docs/protoc-installation/)

Clone the project

```bash
  git clone https://github.com/gabriel-henriq/klever-challenge.git
```

Go to the project directory

```bash
  cd klever-challenge
```

Install Go Mod

```bash
  go mod init github.com/gabriel-henriq/klever-challenge/api
```

Gen Proto files

```bash
  protoc --go-grpc_out=. --proto_path= proto/upvote.proto
```

```bash
  protoc --go_out=. --proto_path= proto/upvote.proto
```

Install Dependencies

```bash
  go mod tidy
```

Run Server Service

```bash
  go run ./server/main.go
```

