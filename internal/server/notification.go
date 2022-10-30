package server

import (
	"context"

	"github.com/Kolyan4ik99/log-audit/internal/model"
	"github.com/Kolyan4ik99/log-audit/pkg/rpc"
	"go.mongodb.org/mongo-driver/mongo"
)

type GRPCServer struct {
	monDB *mongo.Database
}

func NewGRPCNotificationServer(monDB *mongo.Database) *GRPCServer {
	return &GRPCServer{monDB: monDB}
}

func (g *GRPCServer) Send(ctx context.Context, r *rpc.Request) (*rpc.Response, error) {
	logDoc := model.LogDoc{
		Message: r.Message,
		Owner:   r.Owner,
	}

	_, err := g.monDB.Collection("logs").InsertOne(ctx, logDoc)
	if err != nil {
		return nil, err
	}

	return &rpc.Response{}, nil
}
