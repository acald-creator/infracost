package azure

import (
	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getSentinelDataConnectorThreatIntelligenceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_sentinel_data_connector_threat_intelligence",
		RFunc: newSentinelDataConnectorThreatIntelligence,
		ReferenceAttributes: []string{
			"resource_group_name",
			"log_analytics_workspace_id",
		},
	}
}

func newSentinelDataConnectorThreatIntelligence(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := lookupRegion(d, []string{"resource_group_name", "log_analytics_workspace_id"})
	r := &azure.SentinelDataConnectorThreatIntelligence{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
