package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/mszalbach/mcp-bitbucket-server/internal/bitbucket"
)

const (
	baseURL = "https://api.bitbucket.org"
	url     = baseURL + "/rest/api/latest"
	token   = "your-token"
)

type ToolRegister interface {
	Register(server *mcp.Server)
}

var tools = []ToolRegister{
	&ProjectsTools{client: bitbucket.Client{BaseURL: url, Token: token}},
}

func RegisterTools(server *mcp.Server) {
	for _, t := range tools {
		t.Register(server)
	}
}
