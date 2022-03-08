package importresourceactions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportResourceActionsRequestBodyable 
type ImportResourceActionsRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(*string)
    GetOverwriteResourceNamespace()(*bool)
    GetValue()(*string)
    SetFormat(value *string)()
    SetOverwriteResourceNamespace(value *bool)()
    SetValue(value *string)()
}
