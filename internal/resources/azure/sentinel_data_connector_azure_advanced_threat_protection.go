package azure

import (
	"github.com/infracost/infracost/internal/resources"
	"github.com/infracost/infracost/internal/schema"
)

// SentinelDataConnectorAzureAdvancedThreatProtection struct represents <TODO: cloud service short description>.
//
// <TODO: Add any important information about the resource and links to the
// pricing pages or documentation that might be useful to developers in the future, e.g:>
//
// Resource information: https://azure.microsoft.com/<PATH/TO/RESOURCE>/
// Pricing information: https://azure.microsoft.com/<PATH/TO/PRICING>/
type SentinelDataConnectorAzureAdvancedThreatProtection struct {
	Address string
	Region  string
}

// SentinelDataConnectorAzureAdvancedThreatProtectionUsageSchema defines a list which represents the usage schema of SentinelDataConnectorAzureAdvancedThreatProtection.
var SentinelDataConnectorAzureAdvancedThreatProtectionUsageSchema = []*schema.UsageItem{
}

// PopulateUsage parses the u schema.UsageData into the SentinelDataConnectorAzureAdvancedThreatProtection.
// It uses the `infracost_usage` struct tags to populate data into the SentinelDataConnectorAzureAdvancedThreatProtection.
func (r *SentinelDataConnectorAzureAdvancedThreatProtection) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

// BuildResource builds a schema.Resource from a valid SentinelDataConnectorAzureAdvancedThreatProtection struct.
// This method is called after the resource is initialised by an IaC provider.
// See providers folder for more information.
func (r *SentinelDataConnectorAzureAdvancedThreatProtection) BuildResource() *schema.Resource {
	costComponents := []*schema.CostComponent{
		// TODO: add cost components
	}

	return &schema.Resource{
		Name:           r.Address,
		UsageSchema:    SentinelDataConnectorAzureAdvancedThreatProtectionUsageSchema,
		CostComponents: costComponents,
	}
}

