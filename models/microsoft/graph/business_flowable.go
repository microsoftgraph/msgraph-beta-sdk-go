package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BusinessFlowable 
type BusinessFlowable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCustomData()(*string)
    GetDeDuplicationId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicy()(GovernancePolicyable)
    GetPolicyTemplateId()(*string)
    GetRecordVersion()(*string)
    GetSchemaId()(*string)
    GetSettings()(BusinessFlowSettingsable)
    SetCustomData(value *string)()
    SetDeDuplicationId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicy(value GovernancePolicyable)()
    SetPolicyTemplateId(value *string)()
    SetRecordVersion(value *string)()
    SetSchemaId(value *string)()
    SetSettings(value BusinessFlowSettingsable)()
}
