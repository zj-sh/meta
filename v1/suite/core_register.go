package suite

import (
	"github.com/hashicorp/go-plugin"
)

func GRPCSuiteRegister(impl SuiteServer) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "suite",
			MagicCookieValue: "mored",
		},
		Plugins: map[string]plugin.Plugin{
			"suite": &GRPCSuite{
				Suite: impl,
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
