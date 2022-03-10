package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Settingable 
type Settingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetJsonValue()(*string)
    GetOverwriteAllowed()(*bool)
    GetSettingId()(*string)
    GetValueType()(*ManagementParameterValueType)
    SetDisplayName(value *string)()
    SetJsonValue(value *string)()
    SetOverwriteAllowed(value *bool)()
    SetSettingId(value *string)()
    SetValueType(value *ManagementParameterValueType)()
}
