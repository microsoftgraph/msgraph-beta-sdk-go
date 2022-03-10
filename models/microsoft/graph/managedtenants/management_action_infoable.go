package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementActionInfoable 
type ManagementActionInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetManagementActionId()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateVersion()(*int32)
    SetManagementActionId(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateVersion(value *int32)()
}
