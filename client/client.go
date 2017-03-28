package main

import (
	"context"
	"fmt"
	"github.com/azer/logger"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/nsoushi/go-grpc-prometheus-demo/protobuf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net/http"
	"time"
	"os"
)

var (
	conn *grpc.ClientConn
	log  *logger.Logger
)

func main() {
	//gRPC connection
	var err error
	conn, err = grpc.Dial(
		fmt.Sprintf("%s:%s", os.Getenv("GRPC_SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT")),
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
	if err != nil {
		log.Error("Connection error: %v", err)
	}
	defer conn.Close()

	// handle http
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/", indexHandler)

	// serve http
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GRPC_CLIENT_PORT")), nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	client := pb.NewEchoServiceClient(conn)

	message := &pb.Message{
		Message: r.URL.Query().Get("m"),
	}

	response, err := client.EchoService(context.Background(), message)
	if err != nil {
		log.Error("Error:%v", err)
	}

	fmt.Fprintf(w, "Response: %v", response.Message)
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "gRPC client")
	return
}
