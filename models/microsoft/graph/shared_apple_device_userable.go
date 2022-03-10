package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedAppleDeviceUserable 
type SharedAppleDeviceUserable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataQuota()(*int64)
    GetDataToSync()(*bool)
    GetDataUsed()(*int64)
    GetUserPrincipalName()(*string)
    SetDataQuota(value *int64)()
    SetDataToSync(value *bool)()
    SetDataUsed(value *int64)()
    SetUserPrincipalName(value *string)()
}
