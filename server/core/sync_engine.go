package core

import myncer_pb "github.com/hansbala/myncer/proto"

type SyncEngine interface {
	RunSync(sync *myncer_pb.Sync /*const*/) error
}
