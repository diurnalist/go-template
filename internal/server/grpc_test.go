package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/diurnalist/go-template/gen/go/acme/weather/v1"
)

func TestNewWeatherServer(t *testing.T) {
	server := NewWeatherServer()
	assert.NotNil(t, server)
	assert.IsType(t, &WeatherServer{}, server)
}

func TestWeatherServer_GetWeather(t *testing.T) {
	server := NewWeatherServer()
	ctx := context.Background()

	response, err := server.GetWeather(ctx, &v1.GetWeatherRequest{
		Latitude:  40.0,
		Longitude: 0.0,
	})
	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, float32(22.5), response.Temperature)
	assert.Equal(t, v1.Condition_CONDITION_SUNNY, response.Condition)
}

func TestNewGRPCServer(t *testing.T) {
	port := 8081
	grpcServer := NewGRPCServer(port)

	assert.NotNil(t, grpcServer)
	assert.Equal(t, port, grpcServer.port)
	assert.NotNil(t, grpcServer.server)
}
