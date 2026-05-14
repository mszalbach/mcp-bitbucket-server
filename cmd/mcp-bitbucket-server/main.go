package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/mszalbach/mcp-bitbucket-server/tools"
)

func main() {
	// print to stderr to not interfere with the mcp protocol communication on stdout
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "bitbucket-server",
		Version: "0.0.1",
	}, nil)

	tools.InstallTools(server)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
