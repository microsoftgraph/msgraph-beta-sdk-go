package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VppTokenLicenseSummaryable 
type VppTokenLicenseSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppleId()(*string)
    GetAvailableLicenseCount()(*int32)
    GetOrganizationName()(*string)
    GetUsedLicenseCount()(*int32)
    GetVppTokenId()(*string)
    SetAppleId(value *string)()
    SetAvailableLicenseCount(value *int32)()
    SetOrganizationName(value *string)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenId(value *string)()
}
