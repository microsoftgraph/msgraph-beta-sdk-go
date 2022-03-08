package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcBulkRemoteActionResultable 
type CloudPcBulkRemoteActionResultable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFailedDeviceIds()([]string)
    GetNotFoundDeviceIds()([]string)
    GetNotSupportedDeviceIds()([]string)
    GetSuccessfulDeviceIds()([]string)
    SetFailedDeviceIds(value []string)()
    SetNotFoundDeviceIds(value []string)()
    SetNotSupportedDeviceIds(value []string)()
    SetSuccessfulDeviceIds(value []string)()
}
