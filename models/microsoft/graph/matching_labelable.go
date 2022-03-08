package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MatchingLabelable 
type MatchingLabelable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationMode()(*ApplicationMode)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetIsEndpointProtectionEnabled()(*bool)
    GetLabelActions()([]LabelActionBaseable)
    GetName()(*string)
    GetPolicyTip()(*string)
    GetPriority()(*int32)
    GetToolTip()(*string)
    SetApplicationMode(value *ApplicationMode)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetLabelActions(value []LabelActionBaseable)()
    SetName(value *string)()
    SetPolicyTip(value *string)()
    SetPriority(value *int32)()
    SetToolTip(value *string)()
}
