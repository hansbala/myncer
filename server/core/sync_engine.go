package core

import (
	"context"

	myncer_pb "github.com/hansbala/myncer/proto"
)

type SyncEngine interface {
	RunSync(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		sync *myncer_pb.Sync, /*const*/
	) error
}
