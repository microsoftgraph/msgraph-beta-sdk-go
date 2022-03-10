package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColor()(*string)
    GetContentFormats()([]string)
    GetDescription()(*string)
    GetHasProtection()(*bool)
    GetIsActive()(*bool)
    GetIsAppliable()(*bool)
    GetName()(*string)
    GetParent()(SensitivityLabelable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetContentFormats(value []string)()
    SetDescription(value *string)()
    SetHasProtection(value *bool)()
    SetIsActive(value *bool)()
    SetIsAppliable(value *bool)()
    SetName(value *string)()
    SetParent(value SensitivityLabelable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
