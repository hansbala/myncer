package core

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/protobuf/encoding/prototext"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

const (
	// Referenced from the root of the server/ folder.
	cDevConfigPath  = "config.dev.textpb"
	cProdBucketName = "myncer-config"
	cProdObjectPath = "config.prod.textpb"
)

func MustGetConfig() *myncer_pb.Config {
	// First try getting the config from local filesystem (for dev).
	devConfig, err := maybeGetDevConfig()
	if err == nil {
		Printf("Dev config loaded!")
		return devConfig
	}
	Warningf(
		fmt.Sprintf(
			"failed to get dev config because %s. Falling back to production config",
			err.Error(),
		),
	)
	// Fallback to prod config.
	prodConfig, err := getProdConfig()
	if err != nil {
		Errorf("failed to get production config")
		panic(err)
	}
	return prodConfig
}

func maybeGetDevConfig() (*myncer_pb.Config, error) {
	bytes, err := os.ReadFile(cDevConfigPath)
	if err != nil {
		return nil, WrappedError(err, "failed to read dev config file")
	}
	config, err := parseConfigFileBytes(bytes)
	if err != nil {
		return nil, WrappedError(err, "failed to parse dev config file bytes")
	}
	return config, nil
}

func getProdConfig() (*myncer_pb.Config, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, WrappedError(err, "failed to create GCS client")
	}
	defer client.Close()

	rc, err := client.Bucket(cProdBucketName).Object(cProdObjectPath).NewReader(ctx)
	if err != nil {
		return nil, WrappedError(err, "failed to open config file from GCS")
	}
	defer rc.Close()

	bytes, err := io.ReadAll(rc)
	if err != nil {
		return nil, WrappedError(err, "failed to read config file bytes from GCS")
	}

	config, err := parseConfigFileBytes(bytes)
	if err != nil {
		return nil, WrappedError(err, "failed to parse prod config bytes")
	}

	return config, nil
}

func parseConfigFileBytes(bytes []byte /*const*/) (*myncer_pb.Config, error) {
	cs := &myncer_pb.Configs{}
	if err := prototext.Unmarshal(bytes, cs); err != nil {
		return nil, WrappedError(err, "failed to unmarshal config proto bytes")
	}
	configs := cs.GetConfig()
	switch len(configs) {
	case 0:
		return nil, NewError("configs file had no valid configs")
	case 1:
		return configs[0], nil
	default:
		return nil, NewError("config file should declare only one config but had multiple")
	}
}
