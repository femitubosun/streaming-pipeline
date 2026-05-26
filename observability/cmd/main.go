package main

import (
	"context"
	"log"
	"log/slog"
	"maps"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/femitubosun/streaming-pipeline/proto/observability"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedObservabilityServiceServer
	mu             sync.RWMutex
	counts         map[string]int64
	totalProcessed int64
	lastUpdatedMs  int64
}

func (s *server) RecordTransactionMetrics(ctx context.Context, req *pb.TransactionMetricsRequest) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for status, count := range req.StatusCounts {
		s.counts[status] += count
		s.totalProcessed += count
	}
	s.lastUpdatedMs = time.Now().UnixMilli()

	return &pb.Empty{}, nil
}

func (s *server) GetTransactionMetrics(ctx context.Context, req *pb.Empty) (*pb.TransactionMetricsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	statusCounts := make(map[string]int64)

	maps.Copy(statusCounts, s.counts)

	return &pb.TransactionMetricsResponse{
		StatusCounts:   statusCounts,
		TotalProcessed: s.totalProcessed,
		LastUpdatedMs:  s.lastUpdatedMs,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterObservabilityServiceServer(s, &server{
		counts: make(map[string]int64),
	})

	go func() {
		slog.Info("observability server starting", "port", 50051)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	s.GracefulStop()
	slog.Info("observability server stopped")
}
