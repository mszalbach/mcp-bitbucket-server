package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/mszalbach/mcp-bitbucket-server/tools"
)

func main() {

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "bitbucket-server",
		Version: "0.0.1",
	}, nil)

	projectTool := &tools.ProjectTool{}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_projects",
		Description: "Get Bitbucket projects",
	}, projectTool.GetProjects)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}

}
