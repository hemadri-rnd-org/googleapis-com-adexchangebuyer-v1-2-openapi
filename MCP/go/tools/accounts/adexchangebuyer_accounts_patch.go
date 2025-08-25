package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/ad-exchange-buyer-api/mcp-server/config"
	"github.com/ad-exchange-buyer-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Adexchangebuyer_accounts_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Account
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/accounts/%s%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Account
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateAdexchangebuyer_accounts_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_accounts_id",
		mcp.WithDescription("Updates an existing account. This method supports patch semantics."),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("The account id")),
		mcp.WithArray("bidderLocation", mcp.Description("Input parameter: Your bidder locations that have distinct URLs.")),
		mcp.WithString("cookieMatchingNid", mcp.Description("Input parameter: The nid parameter value used in cookie match requests. Please contact your technical account manager if you need to change this.")),
		mcp.WithString("cookieMatchingUrl", mcp.Description("Input parameter: The base URL used in cookie match requests.")),
		mcp.WithNumber("id", mcp.Description("Input parameter: Account id.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Resource type.")),
		mcp.WithNumber("maximumActiveCreatives", mcp.Description("Input parameter: The maximum number of active creatives that an account can have, where a creative is active if it was inserted or bid with in the last 30 days. Please contact your technical account manager if you need to change this.")),
		mcp.WithNumber("maximumTotalQps", mcp.Description("Input parameter: The sum of all bidderLocation.maximumQps values cannot exceed this. Please contact your technical account manager if you need to change this.")),
		mcp.WithNumber("numberActiveCreatives", mcp.Description("Input parameter: The number of creatives that this account inserted or bid with in the last 30 days.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Adexchangebuyer_accounts_patchHandler(cfg),
	}
}
