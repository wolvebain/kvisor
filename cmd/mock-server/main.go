package main

import (
	"context"
	"log/slog"
	"net"
	"os/signal"
	"syscall"

	castaipb "github.com/castai/kvisor/api/v1/runtime"
	"github.com/castai/kvisor/pkg/logging"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log := logging.New(&logging.Config{
		Level: slog.LevelDebug,
	})

	// nolint:gosec
	lis, err := net.Listen("tcp", "0.0.0.0:8443")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer lis.Close()

	srv := grpc.NewServer()
	castaipb.RegisterRuntimeSecurityAgentAPIServer(srv, NewMockServer(log))

	go func() {
		<-ctx.Done()
		log.Info("shutting down grpc ingestor server")
		srv.GracefulStop()
	}()

	println("listening at 0.0.0.0:8443")
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

var _ castaipb.RuntimeSecurityAgentAPIServer = (*MockServer)(nil)

func NewMockServer(log *logging.Logger) *MockServer {
	return &MockServer{
		log: log.WithField("component", "mock_castai"),
	}
}

type MockServer struct {
	log *logging.Logger
}

func (m *MockServer) KubeBenchReportIngest(ctx context.Context, in *castaipb.KubeBenchReport) (*castaipb.KubeBenchReportIngestResponse, error) {
	m.log.Debugf("KubeBenchReportIngest: %v", in)
	return &castaipb.KubeBenchReportIngestResponse{}, nil
}

func (m *MockServer) KubeLinterReportIngest(ctx context.Context, in *castaipb.KubeLinterReport) (*castaipb.KubeLinterReportIngestResponse, error) {
	m.log.Debugf("KubeLinterReportIngest: %v", in)
	return &castaipb.KubeLinterReportIngestResponse{}, nil
}

func (m *MockServer) GetSyncState(ctx context.Context, in *castaipb.GetSyncStateRequest) (*castaipb.GetSyncStateResponse, error) {
	m.log.Debugf("GetSyncState: %v", in.ImageIds)
	return &castaipb.GetSyncStateResponse{}, nil
}

func (m *MockServer) ImageMetadataIngest(ctx context.Context, in *castaipb.ImageMetadata) (*castaipb.ImageMetadataIngestResponse, error) {
	m.log.Debugf("ImageMetadataIngest: %v", in)
	return &castaipb.ImageMetadataIngestResponse{}, nil
}

func (m *MockServer) UpdateSyncState(ctx context.Context, in *castaipb.UpdateSyncStateRequest) (*castaipb.UpdateSyncStateResponse, error) {
	m.log.Debugf("UpdateSyncState: %v", in)
	return &castaipb.UpdateSyncStateResponse{}, nil
}

func (m *MockServer) GetConfiguration(ctx context.Context, in *castaipb.GetConfigurationRequest) (*castaipb.GetConfigurationResponse, error) {
	m.log.Debugf("GetConfiguration: %v", in)
	return &castaipb.GetConfigurationResponse{}, nil
}

func (m *MockServer) EventsWriteStream(server castaipb.RuntimeSecurityAgentAPI_EventsWriteStreamServer) error {
	for {
		event, err := server.Recv()
		if err != nil {
			return err
		}
		m.log.Debugf("event: %v", event)
	}
}

func (m *MockServer) LogsWriteStream(server castaipb.RuntimeSecurityAgentAPI_LogsWriteStreamServer) error {
	for {
		event, err := server.Recv()
		if err != nil {
			return err
		}
		m.log.Debugf("log: %v", event)
	}
}

func (m *MockServer) ContainerStatsWriteStream(server castaipb.RuntimeSecurityAgentAPI_ContainerStatsWriteStreamServer) error {
	for {
		event, err := server.Recv()
		if err != nil {
			return err
		}
		for _, v := range event.Items {
			m.log.Debugf("container_stats, ns=%s, pod=%s, cont=%s, items=%d", v.Namespace, v.PodName, v.ContainerName, len(v.Stats))
		}
	}
}

func (m *MockServer) KubernetesDeltaIngest(server castaipb.RuntimeSecurityAgentAPI_KubernetesDeltaIngestServer) error {
	for {
		event, err := server.Recv()
		if err != nil {
			m.log.Warnf("delta recv: %v", err)
			break
		}
		m.log.Debugf("delta_item: %v", event)
		if err := server.Send(&castaipb.KubernetesDeltaIngestResponse{}); err != nil {
			m.log.Warnf("delta ack send: %v", err)
			break
		}
	}
	return nil
}
