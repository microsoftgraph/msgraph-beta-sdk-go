package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChromeOSDevicePropertyable 
type ChromeOSDevicePropertyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetName()(*string)
    GetUpdatable()(*bool)
    GetValue()(*string)
    GetValueType()(*string)
    SetName(value *string)()
    SetUpdatable(value *bool)()
    SetValue(value *string)()
    SetValueType(value *string)()
}
