package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementTemplateDetailedInfoable 
type ManagementTemplateDetailedInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(*ManagementCategory)
    GetDisplayName()(*string)
    GetManagementTemplateId()(*string)
    GetVersion()(*int32)
    SetCategory(value *ManagementCategory)()
    SetDisplayName(value *string)()
    SetManagementTemplateId(value *string)()
    SetVersion(value *int32)()
}
