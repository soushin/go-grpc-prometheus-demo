package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/nsoushi/go-grpc-prometheus-demo/protobuf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

type grpcServer struct{}

func main() {

	// Create the main listener.
	s, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT")))
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(s)

	// Match connections in order:
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// gRPC server
	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	pb.RegisterEchoServiceServer(grpcS, newGrpcServer())

	// prometheus metrics server
	grpc_prometheus.Register(grpcS)
	httpS := &http.Server{
		Handler: promhttp.Handler(),
	}

	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	m.Serve()
}

func newGrpcServer() *grpcServer {
	return &grpcServer{}
}

func (s *grpcServer) EchoService(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	return &pb.Message{Message: fmt.Sprintf("echo %s", msg.Message)}, nil
}
