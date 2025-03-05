package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogeServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResonse, error) {
	input := req.GetLogEntry()

	//write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)

	if err != nil {
		res := &logs.LogResonse{Result: "failed"}
		return res, err
	}
	//return response
	res := &logs.LogResonse{Result: "logged"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))

	if err != nil {
		fmt.Println("failing to listen here")
		log.Fatalf("Failed to listen to gRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogeServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC server started on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen to gRPC: %v", err)
	}
}
