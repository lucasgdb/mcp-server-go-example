# Hello World - MCP Server in Go

## Introduction

This example demonstrates how to build a minimal Model Context Protocol (MCP) server in Go, using the [mcp-go SDK](https://github.com/mark3labs/mcp-go). MCP is a lightweight JSON‑RPC‑based protocol that lets language model clients discover and invoke external tools in a structured way.

With just a few lines of Go code, you can:

1. Define an MCP server with a human‑readable name and version.
2. Declare one or more tools (commands) with typed parameters and descriptions.
3. Register handler functions that execute your custom logic when a client calls a tool.
4. Expose the server over standard I/O (stdin/stdout) so any MCP‑compatible client (like Cursor, Claude Desktop, or other IDE extensions) can launch and communicate with it without extra setup.

## How It Works

1. **Server Initialization**

   ```go
   mcpServer := server.NewMCPServer(
     "Hello World MCP Server",
     "1.0.0",
     server.WithToolCapabilities(false),
   )
   ```

   - Creates an MCP server named “Hello World MCP Server” at version 1.0.0.
   - Disables automatic tool capability announcements (so only your explicit tools appear).

2. **Tool Definition**

   ```go
   helloWorldTool := mcp.NewTool("hello_world",
     mcp.WithDescription("Greet someone"),
     mcp.WithString("name",
      mcp.Required(),
      mcp.Description("Person’s name"),
     ),
   )
   ```

   - Defines a `hello_world` tool that accepts a single required string parameter `name`.
   - Descriptions help clients display usage information and generate prompts.

3. **Handler Registration**

   ```go
   mcpServer.AddTool(helloWorldTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
     name, ok := req.Params.Arguments["name"].(string)

     if !ok {
      return nil, errors.New("name must be a string")
     }

     return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
   })
   ```

   - Associates the `hello_world` tool with a Go function.
   - Extracts the `name` argument, validates it, and returns a greeting.
   - Errors are surfaced back to the client as JSON‑RPC errors.

4. **Serving Over Stdio**

   ```go
   if err := server.ServeStdio(mcpServer); err != nil {
    fmt.Printf("Server error: %v\n", err)
   }
   ```

   - Starts an event loop reading JSON‑RPC requests from stdin and writing responses to stdout.
   - Allows any MCP client to launch the binary directly and speak MCP over pipes.

## Getting Started

1. **Install Dependencies**

   ```bash
   go mod tidy
   ```

2. **Build the Server**

   ```bash
   go build -o hello-mcp ./src/main.go
   ```

3. **Connect from an MCP Client**

   - In your client’s configuration, point to the `hello-mcp` executable.
   - The client will automatically discover the `hello_world` tool schema and let you invoke it.

   Manually example:

   ```json
   {
     "mcpServers": {
       "demo-go-mcp": {
         "command": "path-to-server/hello-mcp"
       }
     }
   }
   ```

Enjoy building more complex workflows by adding additional tools, parameters, and transports (e.g., HTTP/SSE) as needed!
