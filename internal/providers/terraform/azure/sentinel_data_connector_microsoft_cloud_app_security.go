package azure

import (
	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getSentinelDataConnectorMicrosoftCloudAppSecurityRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_sentinel_data_connector_microsoft_cloud_app_security",
		RFunc: newSentinelDataConnectorMicrosoftCloudAppSecurity,
		ReferenceAttributes: []string{
			"resource_group_name",
			"log_analytics_workspace_id",
		},
	}
}

func newSentinelDataConnectorMicrosoftCloudAppSecurity(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := lookupRegion(d, []string{"resource_group_name", "log_analytics_workspace_id"})
	r := &azure.SentinelDataConnectorMicrosoftCloudAppSecurity{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
