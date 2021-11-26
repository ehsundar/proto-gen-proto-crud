# proto-gen-proto-crud

This is a `protoc` plugin which creates auto-generated proto for a particular proto message. This auto-generated proto contains suitable requests and responses along with a gRPC service in order to serve CRUD operations on that particuar proto message.

here is a list of generated RPCs:

- GetItem
- GetAllItems
- CreateItem
- UpdateItem
- DeleteItem

to install the plugin:

```bash
go install github.com/ehsundar/protoc-gen-proto-crud/cmd/protoc-gen-proto-crud
```

to generate proto for the example message:

```bash
protoc --proto-crud_out=. --proto-crud_opt=paths=source_relative example/item.proto
```

to generate gRPC and protobuf codes for go:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative example/*.proto
```

# Interact with server

get one with id:

```bash
grpcurl -d '{"id":1}' -plaintext localhost:8000 example.ItemStorage.GetItem
```

get all:

```bash
grpcurl -plaintext localhost:8000 example.ItemStorage.GetAllItems
```

create one:

```bash
grpcurl -d '{"item": {"Name": "Ehsan"}}' -plaintext localhost:8000 example.ItemStorage.CreateItem
```

update one:

```bash
grpcurl -d '{"id": 2, "item": {"Name": "Mehdi"}}' -plaintext localhost:8000 example.ItemStorage.UpdateItem
```

delete one:

```bash
grpcurl -d '{"id": 1}' -plaintext localhost:8000 example.ItemStorage.DeleteItem
```



# TODO

- [ ] Embed template inside binary
- [ ] Generate interface implementation for postgres
