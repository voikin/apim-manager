package controller

import (
	"fmt"

	harprofilerpb "github.com/voikin/apim-proto/gen/go/apim_har_profiler/v1"
	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	openapiexporterpb "github.com/voikin/apim-proto/gen/go/apim_openapi_exporter/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConfigURLs struct {
	HARProfiler     string `mapstructure:"har_profiler"`
	ProfileStore    string `mapstructure:"profile_store"`
	OpenAPIExproter string `mapstructure:"openapi_exporter"`
}

type Controller struct {
	managerpb.UnimplementedManagerServiceServer

	profileStoreClient    profilestorepb.ProfileStoreServiceClient
	harProfilerClient     harprofilerpb.HARProfilerServiceClient
	openapiExporterClient openapiexporterpb.OpenAPIExporterServiceClient
}

func New(cfg *ConfigURLs) (*Controller, error) {
	profileStoreConn, err := grpc.NewClient(cfg.ProfileStore, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: profile store: %w", err)
	}

	profileStoreClient := profilestorepb.NewProfileStoreServiceClient(profileStoreConn)

	harProfilerConn, err := grpc.NewClient(cfg.HARProfiler, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: har profiler: %w", err)
	}

	harProfilerClient := harprofilerpb.NewHARProfilerServiceClient(harProfilerConn)

	openapiExporterConn, err := grpc.NewClient(cfg.OpenAPIExproter, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: openapi exporter: %w", err)
	}

	openapiExporterClient := openapiexporterpb.NewOpenAPIExporterServiceClient(openapiExporterConn)

	return &Controller{
		profileStoreClient:    profileStoreClient,
		harProfilerClient:     harProfilerClient,
		openapiExporterClient: openapiExporterClient,
	}, nil
}
