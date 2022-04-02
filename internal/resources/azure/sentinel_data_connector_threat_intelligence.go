package azure

import (
	"github.com/infracost/infracost/internal/resources"
	"github.com/infracost/infracost/internal/schema"
)

// SentinelDataConnectorThreatIntelligence struct represents <TODO: cloud service short description>.
//
// <TODO: Add any important information about the resource and links to the
// pricing pages or documentation that might be useful to developers in the future, e.g:>
//
// Resource information: https://azure.microsoft.com/<PATH/TO/RESOURCE>/
// Pricing information: https://azure.microsoft.com/<PATH/TO/PRICING>/
type SentinelDataConnectorThreatIntelligence struct {
	Address string
	Region  string
}

// SentinelDataConnectorThreatIntelligenceUsageSchema defines a list which represents the usage schema of SentinelDataConnectorThreatIntelligence.
var SentinelDataConnectorThreatIntelligenceUsageSchema = []*schema.UsageItem{}

// PopulateUsage parses the u schema.UsageData into the SentinelDataConnectorThreatIntelligence.
// It uses the `infracost_usage` struct tags to populate data into the SentinelDataConnectorThreatIntelligence.
func (r *SentinelDataConnectorThreatIntelligence) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

// BuildResource builds a schema.Resource from a valid SentinelDataConnectorThreatIntelligence struct.
// This method is called after the resource is initialised by an IaC provider.
// See providers folder for more information.
func (r *SentinelDataConnectorThreatIntelligence) BuildResource() *schema.Resource {
	return &schema.Resource{
		Name:        r.Address,
		IsSkipped:   true,
		NoPrice:     true,
		UsageSchema: SentinelDataConnectorThreatIntelligenceUsageSchema,
	}
}
