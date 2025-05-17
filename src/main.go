package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	mcpServer := server.NewMCPServer(
		"Hello World MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	helloWorldTool := mcp.NewTool("hello_world",
		mcp.WithDescription("Greet someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Person’s name"),
		),
	)

	mcpServer.AddTool(helloWorldTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, ok := req.Params.Arguments["name"].(string)

		if !ok {
			return nil, errors.New("name must be a string")
		}

		return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
	})

	if err := server.ServeStdio(mcpServer); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
