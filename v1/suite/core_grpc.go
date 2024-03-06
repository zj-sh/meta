package suite

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type GRPCSuite struct {
	plugin.Plugin
	Suite SuiteServer
}

func (p *GRPCSuite) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterSuiteServer(s, p.Suite)
	return nil
}

func (p *GRPCSuite) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return NewSuiteClient(c), nil
}
