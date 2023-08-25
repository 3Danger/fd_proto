package cashdeskpb

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

const (
	metadataWorkstationID = "workstation-id"
)

func GetWorkstationID(ctx context.Context) (string, bool) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return "", false
	}

	mdField := md.Get(metadataWorkstationID)

	if len(mdField) == 0 {
		return "", false
	}

	return mdField[0], true
}

func checkMetadata(ctx context.Context) error {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return errors.New("no metadata was found")
	}

	mdField := md.Get(metadataWorkstationID)

	if len(mdField) == 0 {
		return fmt.Errorf("no %s field in metadata was found", metadataWorkstationID)
	}

	return nil
}
