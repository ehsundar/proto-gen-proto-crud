package main

import (
	"context"
	"database/sql"
	"net"

	"github.com/ehsundar/protoc-gen-proto-crud/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	db := getDBOrPanic()
	storage := example.NewPostgresItemStorage(db)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	example.RegisterItemStorageServer(grpcServer, storage)

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func getDBOrPanic() *sql.DB {
	dataSource := "postgres://postgres:mypassword@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	if err = db.PingContext(context.Background()); err != nil {
		panic(err)
	}
	return db
}
