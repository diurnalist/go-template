package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	v1 "github.com/diurnalist/go-template/gen/go/acme/weather/v1"
)

// WeatherServer implements the WeatherServiceServer interface.
type WeatherServer struct {
	v1.UnimplementedWeatherServiceServer
}

// NewWeatherServer creates a new WeatherServer instance.
func NewWeatherServer() *WeatherServer {
	return &WeatherServer{
		UnimplementedWeatherServiceServer: v1.UnimplementedWeatherServiceServer{},
	}
}

// GetWeather implements the GetWeather RPC method.
func (s *WeatherServer) GetWeather(_ context.Context, req *v1.GetWeatherRequest) (*v1.GetWeatherResponse, error) {
	// Simple mock implementation - in a real service, this would call a weather API
	// or database to get actual weather data based on the coordinates.
	const mockTemperature = 22.5 // Mock temperature in Celsius
	condition := v1.Condition_CONDITION_SUNNY

	// Simple logic to determine weather condition based on coordinates.
	if req.GetLatitude() > 50 || req.GetLongitude() < -10 {
		condition = v1.Condition_CONDITION_RAINY
	}

	return &v1.GetWeatherResponse{
		Temperature: mockTemperature,
		Condition:   condition,
	}, nil
}

// GRPCServer represents a gRPC server instance.
type GRPCServer struct {
	server *grpc.Server
	port   int
}

// NewGRPCServer creates a new gRPC server instance.
func NewGRPCServer(port int) *GRPCServer {
	server := grpc.NewServer()

	// Register the weather service.
	weatherServer := NewWeatherServer()
	v1.RegisterWeatherServiceServer(server, weatherServer)

	// Enable reflection for debugging.
	reflection.Register(server)

	return &GRPCServer{
		server: server,
		port:   port,
	}
}

// Start starts the gRPC server.
func (s *GRPCServer) Start(ctx context.Context) error {
	lisCfg := net.ListenConfig{
		Control:   nil,
		KeepAlive: -1,
		KeepAliveConfig: net.KeepAliveConfig{
			Enable:   false,
			Idle:     0,
			Interval: 0,
			Count:    0,
		},
	}
	lis, err := lisCfg.Listen(ctx, "tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed	to listen: %w", err)
	}

	slog.InfoContext(ctx, "gRPC server starting", "port", s.port)
	return s.server.Serve(lis)
}

// Stop gracefully stops the gRPC server.
func (s *GRPCServer) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}
