package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicableTo()(*SensitivityLabelTarget)
    GetApplicationMode()(*ApplicationMode)
    GetAssignedPolicies()([]LabelPolicyable)
    GetAutoLabeling()(AutoLabelingable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetIsEndpointProtectionEnabled()(*bool)
    GetLabelActions()([]LabelActionBaseable)
    GetName()(*string)
    GetPriority()(*int32)
    GetSublabels()([]SensitivityLabelable)
    GetToolTip()(*string)
    SetApplicableTo(value *SensitivityLabelTarget)()
    SetApplicationMode(value *ApplicationMode)()
    SetAssignedPolicies(value []LabelPolicyable)()
    SetAutoLabeling(value AutoLabelingable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetLabelActions(value []LabelActionBaseable)()
    SetName(value *string)()
    SetPriority(value *int32)()
    SetSublabels(value []SensitivityLabelable)()
    SetToolTip(value *string)()
}
