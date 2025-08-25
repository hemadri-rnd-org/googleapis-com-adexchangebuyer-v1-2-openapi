package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ad-exchange-buyer-api/mcp-server/config"
	"github.com/ad-exchange-buyer-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Adexchangebuyer_creatives_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		accountIdVal, ok := args["accountId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: accountId"), nil
		}
		accountId, ok := accountIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: accountId"), nil
		}
		buyerCreativeIdVal, ok := args["buyerCreativeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: buyerCreativeId"), nil
		}
		buyerCreativeId, ok := buyerCreativeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: buyerCreativeId"), nil
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
		url := fmt.Sprintf("%s/creatives/%s/%s%s", cfg.BaseURL, accountId, buyerCreativeId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.Creative
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

func CreateAdexchangebuyer_creatives_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_creatives_accountId_buyerCreativeId",
		mcp.WithDescription("Gets the status for a single creative. A creative will be available 30-40 minutes after submission."),
		mcp.WithNumber("accountId", mcp.Required(), mcp.Description("The id for the account that will serve this creative.")),
		mcp.WithString("buyerCreativeId", mcp.Required(), mcp.Description("The buyer-specific id for this creative.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Adexchangebuyer_creatives_getHandler(cfg),
	}
}
