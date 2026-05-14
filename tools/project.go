package tools

import (
	"context"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ProjectTool struct {
}

type GetProjectsInput struct {
}

type test struct {
	Name string `json:"name"`
}

func (p *ProjectTool) GetProjects(ctx context.Context, req *mcp.CallToolRequest, input GetProjectsInput) (*mcp.CallToolResult, any, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("Hello2")
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Unable to fetch projects."},
		},
	}, test{Name: "Test Project"}, nil
}
