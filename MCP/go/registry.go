package main

import (
	"github.com/ad-exchange-buyer-api/mcp-server/config"
	"github.com/ad-exchange-buyer-api/mcp-server/models"
	tools_accounts "github.com/ad-exchange-buyer-api/mcp-server/tools/accounts"
	tools_creatives "github.com/ad-exchange-buyer-api/mcp-server/tools/creatives"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_accounts.CreateAdexchangebuyer_accounts_listTool(cfg),
		tools_accounts.CreateAdexchangebuyer_accounts_getTool(cfg),
		tools_accounts.CreateAdexchangebuyer_accounts_patchTool(cfg),
		tools_accounts.CreateAdexchangebuyer_accounts_updateTool(cfg),
		tools_creatives.CreateAdexchangebuyer_creatives_listTool(cfg),
		tools_creatives.CreateAdexchangebuyer_creatives_insertTool(cfg),
		tools_creatives.CreateAdexchangebuyer_creatives_getTool(cfg),
	}
}
