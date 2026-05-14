package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type toolAdder struct {
	adder func(server *mcp.Server)
}

var tools = []toolAdder{
	{AddProjectsTools},
}

func InstallTools(server *mcp.Server) {
	for _, t := range tools {
		t.adder(server)
	}
}
