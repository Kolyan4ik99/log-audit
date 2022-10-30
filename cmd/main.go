package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Kolyan4ik99/log-audit/config"
	"github.com/Kolyan4ik99/log-audit/internal/db/mongodb"
	"github.com/Kolyan4ik99/log-audit/internal/server"
	"github.com/Kolyan4ik99/log-audit/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.NewConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := mongodb.NewClient(ctx, cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	s := grpc.NewServer()
	serv := server.NewGRPCNotificationServer(conn.Database("log"))
	rpc.RegisterNotificationServer(s, serv)

	listen, err := net.Listen(cfg.Server.NetProtocol, cfg.Server.GetAddressPort())
	if err != nil {
		return
	}
	fmt.Println(cfg.Server.GetAddressPort())

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
