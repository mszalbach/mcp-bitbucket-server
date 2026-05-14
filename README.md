# mcp-bitbucket-server

## Description

This is a mcp server allowing mcp clients connecting to Bitbucket Server or Bitbucket Data Center.

# Manual testing

Install the mcp insecptor. This needs node and npx.

```bash
npx @modelcontextprotocol/inspector
```
Use the follwing settings:

| Setting        | Value                               |
| -------------- | ----------------------------------- |
| Transport Type | STDIO                               |
| Command        | go                                  |
| Arguments      | un cmd/mcp-bitbucket-server/main.go |