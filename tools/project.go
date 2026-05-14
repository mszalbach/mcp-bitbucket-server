package tools

import (
	"context"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GetProjectsInput struct {
}

type GetProjectsOutput struct {
	Name string `json:"name"`
}

func AddProjectsTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_projects",
		Description: "Get Bitbucket projects",
	}, GetProjects)
}

func GetProjects(ctx context.Context, req *mcp.CallToolRequest, input GetProjectsInput) (*mcp.CallToolResult, any, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("Hello2")
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Unable to fetch projects."},
		},
	}, GetProjectsOutput{Name: "Test Project"}, nil
}
