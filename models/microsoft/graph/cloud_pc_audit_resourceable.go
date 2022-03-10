package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditResourceable 
type CloudPcAuditResourceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetModifiedProperties()([]CloudPcAuditPropertyable)
    GetResourceId()(*string)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetModifiedProperties(value []CloudPcAuditPropertyable)()
    SetResourceId(value *string)()
    SetType(value *string)()
}
