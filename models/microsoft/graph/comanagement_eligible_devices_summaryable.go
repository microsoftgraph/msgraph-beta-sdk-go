package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComanagementEligibleDevicesSummaryable 
type ComanagementEligibleDevicesSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetComanagedCount()(*int32)
    GetEligibleButNotAzureAdJoinedCount()(*int32)
    GetEligibleCount()(*int32)
    GetIneligibleCount()(*int32)
    GetNeedsOsUpdateCount()(*int32)
    SetComanagedCount(value *int32)()
    SetEligibleButNotAzureAdJoinedCount(value *int32)()
    SetEligibleCount(value *int32)()
    SetIneligibleCount(value *int32)()
    SetNeedsOsUpdateCount(value *int32)()
}
