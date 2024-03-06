package suite

import (
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/zohu/zlog"
)

func GRPCSuiteClient(dir, cmd, filename string) (SuiteClient, *plugin.Client, error) {
	if strings.Contains(cmd, "./{file}") && strings.HasPrefix(filename, "/") {
		cmd = strings.ReplaceAll(cmd, "./{file}", filename)
	} else {
		cmd = strings.ReplaceAll(cmd, "{file}", filename)
	}
	runCmd := exec.Command("sh", "-c", cmd)
	runCmd.Dir = dir
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "suite",
			MagicCookieValue: "mored",
		},
		Cmd: runCmd,
		Plugins: map[string]plugin.Plugin{
			"suite": &GRPCSuite{},
		},
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger: hclog.New(&hclog.LoggerOptions{
			Output:      zlog.SafeWriter(),
			Level:       hclog.Debug,
			DisableTime: true,
			Name:        "plugin",
		}),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, err
	}
	raw, err := rpcClient.Dispense("suite")
	if err != nil {
		return nil, nil, err
	}
	return raw.(SuiteClient), client, nil
}
