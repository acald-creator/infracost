package azure

import (
	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getDataFactoryIntegrationRuntimeSelfHostedRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_data_factory_integration_runtime_self_hosted",
		CoreRFunc: newDataFactoryIntegrationRuntimeSelfHosted,
		ReferenceAttributes: []string{
			"data_factory_id",
		},
	}
}

func newDataFactoryIntegrationRuntimeSelfHosted(d *schema.ResourceData) schema.CoreResource {
	dataFactoryRefs := d.References("data_factory_id")
	var region string

	if len(dataFactoryRefs) > 0 {
		region = lookupRegion(dataFactoryRefs[0], []string{"resource_group_name"})
	}

	r := &azure.DataFactoryIntegrationRuntimeSelfHosted{
		Address: d.Address,
		Region:  region,
	}
	return r
}
