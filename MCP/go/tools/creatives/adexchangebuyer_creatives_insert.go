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

func Adexchangebuyer_creatives_insertHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		var requestBody models.Creative
		
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
		url := fmt.Sprintf("%s/creatives%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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

func CreateAdexchangebuyer_creatives_insertTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_creatives",
		mcp.WithDescription("Submit a new creative."),
		mcp.WithNumber("accountId", mcp.Description("Input parameter: Account id.")),
		mcp.WithString("videoURL", mcp.Description("Input parameter: The url to fetch a video ad. If set, HTMLSnippet should not be set.")),
		mcp.WithString("apiUploadTimestamp", mcp.Description("Input parameter: The last upload timestamp of this creative if it was uploaded via API. Read-only. The value of this field is generated, and will be ignored for uploads. (formatted RFC 3339 timestamp).")),
		mcp.WithArray("corrections", mcp.Description("Input parameter: Shows any corrections that were applied to this creative. Read-only. This field should not be set in requests.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Resource type.")),
		mcp.WithString("advertiserName", mcp.Description("Input parameter: The name of the company being advertised in the creative.")),
		mcp.WithNumber("width", mcp.Description("Input parameter: Ad width.")),
		mcp.WithArray("vendorType", mcp.Description("Input parameter: All vendor types for the ads that may be shown from this snippet.")),
		mcp.WithNumber("version", mcp.Description("Input parameter: The version for this creative. Read-only. This field should not be set in requests.")),
		mcp.WithArray("disapprovalReasons", mcp.Description("Input parameter: The reasons for disapproval, if any. Note that not all disapproval reasons may be categorized, so it is possible for the creative to have a status of DISAPPROVED with an empty list for disapproval_reasons. In this case, please reach out to your TAM to help debug the issue. Read-only. This field should not be set in requests.")),
		mcp.WithString("HTMLSnippet", mcp.Description("Input parameter: The HTML snippet that displays the ad when inserted in the web page. If set, videoURL should not be set.")),
		mcp.WithArray("productCategories", mcp.Description("Input parameter: Detected product categories, if any. Read-only. This field should not be set in requests.")),
		mcp.WithArray("attribute", mcp.Description("Input parameter: All attributes for the ads that may be shown from this snippet.")),
		mcp.WithArray("restrictedCategories", mcp.Description("Input parameter: All restricted categories for the ads that may be shown from this snippet.")),
		mcp.WithString("status", mcp.Description("Input parameter: Creative serving status. Read-only. This field should not be set in requests.")),
		mcp.WithString("agencyId", mcp.Description("Input parameter: The agency id for this creative.")),
		mcp.WithString("buyerCreativeId", mcp.Description("Input parameter: A buyer-specific id identifying the creative in this ad.")),
		mcp.WithArray("clickThroughUrl", mcp.Description("Input parameter: The set of destination urls for the snippet.")),
		mcp.WithObject("filteringReasons", mcp.Description("Input parameter: The filtering reasons for the creative. Read-only. This field should not be set in requests.")),
		mcp.WithNumber("height", mcp.Description("Input parameter: Ad height.")),
		mcp.WithArray("sensitiveCategories", mcp.Description("Input parameter: Detected sensitive categories, if any. Read-only. This field should not be set in requests.")),
		mcp.WithArray("impressionTrackingUrl", mcp.Description("Input parameter: The set of urls to be called to record an impression.")),
		mcp.WithArray("advertiserId", mcp.Description("Input parameter: Detected advertiser id, if any. Read-only. This field should not be set in requests.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Adexchangebuyer_creatives_insertHandler(cfg),
	}
}
