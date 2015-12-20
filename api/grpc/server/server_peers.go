package server

import (
	"github.com/docker/containerd/api/grpc/types"
	"golang.org/x/net/context"
)

func (s *apiServer) ListPeers(ctx context.Context, r *types.PeersRequest) (*types.PeersResponse, error) {
	peers := []string{}
	for _, m := range s.serf.Members() {
		// log.Infof("Node(%s) = ADDR(%s)", m.Name, m.Tags["ADVERTISE_ADDR"])
		peers = append(peers, m.Tags["ADVERTISE_ADDR"])
	}
	return &types.PeersResponse{
		Peers: peers,
	}, nil
}
