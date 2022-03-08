package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeConfigurationable 
type OfficeConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClientConfigurations()([]OfficeClientConfigurationable)
    GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable)
    SetClientConfigurations(value []OfficeClientConfigurationable)()
    SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)()
}
