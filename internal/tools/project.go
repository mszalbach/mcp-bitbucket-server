package tools

import (
	"context"
	"encoding/json"
	"io"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/mszalbach/mcp-bitbucket-server/internal/bitbucket"
)

type ProjectsTools struct {
	client bitbucket.Client
}

// TODO support pagination
type GetProjectsInput struct{}

// TODO only return most needed fields for an mcp client
type GetProjectsContent struct {
	Size       int  `json:"size" jsonschema:"the message to convey"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		Key    string `json:"key"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Public bool   `json:"public"`
		Type   string `json:"type"`
		Links  struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
		Description string `json:"description,omitempty"`
	} `json:"values"`
	Start         int `json:"start"`
	NextPageStart int `json:"nextPageStart"`
}

func (p *ProjectsTools) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_projects",
		Description: "Get Bitbucket projects",
	}, p.GetProjects)
}

// TODO check for http errors
func (p *ProjectsTools) GetProjects(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GetProjectsInput,
) (*mcp.CallToolResult, GetProjectsContent, error) {
	res, err := p.client.Do("GET", "/projects", nil)
	if err != nil {
		return errorResult(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errorResult(err)
	}

	var content GetProjectsContent
	err = json.Unmarshal(body, &content)
	if err != nil {
		return errorResult(err)
	}

	return nil, content, nil
}

func errorResult(err error) (*mcp.CallToolResult, GetProjectsContent, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: err.Error()},
		},
	}, GetProjectsContent{}, nil
}
