package tools

import (
	"context"
	"io"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/mszalbach/mcp-bitbucket-server/internal/bitbucket"
)

type ProjectsTools struct {
	client bitbucket.Client
}

type GetProjectsInput struct{}

type GetProjectsContent struct {
	Name string `json:"name"`
}

func (p *ProjectsTools) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_projects",
		Description: "Get Bitbucket projects",
	}, p.GetProjects)
}

func (p *ProjectsTools) GetProjects(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GetProjectsInput,
) (*mcp.CallToolResult, any, error) {
	res, err := p.client.Do("GET", "/projects", nil)
	if err != nil {
		return nil, nil, err
	}
	res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, nil, nil
}
