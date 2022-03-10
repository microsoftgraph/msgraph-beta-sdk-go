package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionLabelable 
type InformationProtectionLabelable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColor()(*string)
    GetDescription()(*string)
    GetIsActive()(*bool)
    GetName()(*string)
    GetParent()(ParentLabelDetailsable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetDescription(value *string)()
    SetIsActive(value *bool)()
    SetName(value *string)()
    SetParent(value ParentLabelDetailsable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
