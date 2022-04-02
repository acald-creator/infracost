package azure

import (
	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getSentinelDataConnectorAwsCloudTrailRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_sentinel_data_connector_aws_cloud_trail",
		RFunc: newSentinelDataConnectorAwsCloudTrail,
		ReferenceAttributes: []string{
			"resource_group_name",
			"log_analytics_workspace_id",
		},
	}
}

func newSentinelDataConnectorAwsCloudTrail(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := lookupRegion(d, []string{"resource_group_name", "log_analytics_workspace_id"})
	r := &azure.SentinelDataConnectorAwsCloudTrail{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
