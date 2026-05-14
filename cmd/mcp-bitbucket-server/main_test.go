package main

import (
	"context"
	"os/exec"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mcp_server_has_tools(t *testing.T) {
	ctx := context.Background()

	client := mcp.NewClient(&mcp.Implementation{Name: "mcp-client", Version: "v1.0.0"}, nil)

	transport := &mcp.CommandTransport{Command: exec.Command("go", "run", "main.go")}
	session, err := client.Connect(ctx, transport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { session.Close() })

	res, err := session.ListTools(ctx, &mcp.ListToolsParams{})

	require.NoError(t, err)

	assert.Len(t, res.Tools, 1)

	assert.Equal(t, "get_projects", res.Tools[0].Name)
}
