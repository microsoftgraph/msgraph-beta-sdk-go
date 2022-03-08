package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HasPayloadLinkResultItemable 
type HasPayloadLinkResultItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetError()(*string)
    GetHasLink()(*bool)
    GetPayloadId()(*string)
    GetSources()([]DeviceAndAppManagementAssignmentSource)
    SetError(value *string)()
    SetHasLink(value *bool)()
    SetPayloadId(value *string)()
    SetSources(value []DeviceAndAppManagementAssignmentSource)()
}
