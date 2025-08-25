package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreativesList represents the CreativesList schema from the OpenAPI specification
type CreativesList struct {
	Kind string `json:"kind,omitempty"` // Resource type.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Continuation token used to page through creatives. To retrieve the next page of results, set the next request's "pageToken" value to this.
	Items []Creative `json:"items,omitempty"` // A list of creatives.
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Numberactivecreatives int `json:"numberActiveCreatives,omitempty"` // The number of creatives that this account inserted or bid with in the last 30 days.
	Bidderlocation []map[string]interface{} `json:"bidderLocation,omitempty"` // Your bidder locations that have distinct URLs.
	Cookiematchingnid string `json:"cookieMatchingNid,omitempty"` // The nid parameter value used in cookie match requests. Please contact your technical account manager if you need to change this.
	Cookiematchingurl string `json:"cookieMatchingUrl,omitempty"` // The base URL used in cookie match requests.
	Id int `json:"id,omitempty"` // Account id.
	Kind string `json:"kind,omitempty"` // Resource type.
	Maximumactivecreatives int `json:"maximumActiveCreatives,omitempty"` // The maximum number of active creatives that an account can have, where a creative is active if it was inserted or bid with in the last 30 days. Please contact your technical account manager if you need to change this.
	Maximumtotalqps int `json:"maximumTotalQps,omitempty"` // The sum of all bidderLocation.maximumQps values cannot exceed this. Please contact your technical account manager if you need to change this.
}

// AccountsList represents the AccountsList schema from the OpenAPI specification
type AccountsList struct {
	Items []Account `json:"items,omitempty"` // A list of accounts.
	Kind string `json:"kind,omitempty"` // Resource type.
}

// Creative represents the Creative schema from the OpenAPI specification
type Creative struct {
	Accountid int `json:"accountId,omitempty"` // Account id.
	Videourl string `json:"videoURL,omitempty"` // The url to fetch a video ad. If set, HTMLSnippet should not be set.
	Apiuploadtimestamp string `json:"apiUploadTimestamp,omitempty"` // The last upload timestamp of this creative if it was uploaded via API. Read-only. The value of this field is generated, and will be ignored for uploads. (formatted RFC 3339 timestamp).
	Corrections []map[string]interface{} `json:"corrections,omitempty"` // Shows any corrections that were applied to this creative. Read-only. This field should not be set in requests.
	Kind string `json:"kind,omitempty"` // Resource type.
	Advertisername string `json:"advertiserName,omitempty"` // The name of the company being advertised in the creative.
	Width int `json:"width,omitempty"` // Ad width.
	Vendortype []int `json:"vendorType,omitempty"` // All vendor types for the ads that may be shown from this snippet.
	Version int `json:"version,omitempty"` // The version for this creative. Read-only. This field should not be set in requests.
	Disapprovalreasons []map[string]interface{} `json:"disapprovalReasons,omitempty"` // The reasons for disapproval, if any. Note that not all disapproval reasons may be categorized, so it is possible for the creative to have a status of DISAPPROVED with an empty list for disapproval_reasons. In this case, please reach out to your TAM to help debug the issue. Read-only. This field should not be set in requests.
	Htmlsnippet string `json:"HTMLSnippet,omitempty"` // The HTML snippet that displays the ad when inserted in the web page. If set, videoURL should not be set.
	Productcategories []int `json:"productCategories,omitempty"` // Detected product categories, if any. Read-only. This field should not be set in requests.
	Attribute []int `json:"attribute,omitempty"` // All attributes for the ads that may be shown from this snippet.
	Restrictedcategories []int `json:"restrictedCategories,omitempty"` // All restricted categories for the ads that may be shown from this snippet.
	Status string `json:"status,omitempty"` // Creative serving status. Read-only. This field should not be set in requests.
	Agencyid string `json:"agencyId,omitempty"` // The agency id for this creative.
	Buyercreativeid string `json:"buyerCreativeId,omitempty"` // A buyer-specific id identifying the creative in this ad.
	Clickthroughurl []string `json:"clickThroughUrl,omitempty"` // The set of destination urls for the snippet.
	Filteringreasons map[string]interface{} `json:"filteringReasons,omitempty"` // The filtering reasons for the creative. Read-only. This field should not be set in requests.
	Height int `json:"height,omitempty"` // Ad height.
	Sensitivecategories []int `json:"sensitiveCategories,omitempty"` // Detected sensitive categories, if any. Read-only. This field should not be set in requests.
	Impressiontrackingurl []string `json:"impressionTrackingUrl,omitempty"` // The set of urls to be called to record an impression.
	Advertiserid []string `json:"advertiserId,omitempty"` // Detected advertiser id, if any. Read-only. This field should not be set in requests.
}
