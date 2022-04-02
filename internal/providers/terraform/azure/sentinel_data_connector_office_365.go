package azure

import (
	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getSentinelDataConnectorOffice365RegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_sentinel_data_connector_office_365",
		RFunc: newSentinelDataConnectorOffice365,
		ReferenceAttributes: []string{
			"resource_group_name",
			"log_analytics_workspace_id",
		},
	}
}

func newSentinelDataConnectorOffice365(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := lookupRegion(d, []string{"resource_group_name", "log_analytics_workspace_id"})
	r := &azure.SentinelDataConnectorOffice365{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
