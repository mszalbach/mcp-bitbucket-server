# mcp-bitbucket-server

## Description

This is a test mcp server allowing mcp clients connecting to Bitbucket Server or Bitbucket Data Center.
THis is just for learning the patterns and will probally never be finished.

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