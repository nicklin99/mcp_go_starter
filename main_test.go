// main_test.go
package main

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name       string
		arguments  map[string]interface{}
		wantResult string
		wantErr    bool
	}{
		{
			name: "Valid name",
			arguments: map[string]interface{}{
				"name": "Alice",
			},
			wantResult: "Hello, Alice!",
			wantErr:    false,
		},
		{
			name:      "Missing name",
			arguments: map[string]interface{}{},
			wantErr:   true,
		},
		{
			name: "Invalid name type",
			arguments: map[string]interface{}{
				"name": 123,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := mcp.CallToolRequest{
				Request: mcp.Request{
                    Method: "hello_world",
                },
			}
            request.Params.Arguments = tt.arguments

			result, err := helloHandler(context.Background(), request)
			t.Logf("Result: %v, Error: %v", result, err)
		})
	}
}
