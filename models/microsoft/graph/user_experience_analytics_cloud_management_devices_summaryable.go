package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsCloudManagementDevicesSummaryable 
type UserExperienceAnalyticsCloudManagementDevicesSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCoManagedDeviceCount()(*int32)
    GetIntuneDeviceCount()(*int32)
    GetTenantAttachDeviceCount()(*int32)
    SetCoManagedDeviceCount(value *int32)()
    SetIntuneDeviceCount(value *int32)()
    SetTenantAttachDeviceCount(value *int32)()
}
