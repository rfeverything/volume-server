package server

import (
	"context"

	volume "github.com/rfeverything/volume-server/proto"
)

func NewVolumeServer() *VolumeServer {
	return &VolumeServer{}
}

type VolumeServer struct {
}

func (v *VolumeServer) Write(ctx context.Context, req *volume.WriteReq) (*volume.WriteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VolumeServer) Read(ctx context.Context, req *volume.ReadReq) (*volume.ReadResp, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VolumeServer) Status(ctx context.Context, req *volume.StatusReq) (*volume.StatusResp, error) {
	//TODO implement me
	panic("implement me")
}
