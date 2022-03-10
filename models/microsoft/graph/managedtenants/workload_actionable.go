package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkloadActionable 
type WorkloadActionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionId()(*string)
    GetCategory()(*WorkloadActionCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLicenses()([]string)
    GetService()(*string)
    GetSettings()([]Settingable)
    SetActionId(value *string)()
    SetCategory(value *WorkloadActionCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLicenses(value []string)()
    SetService(value *string)()
    SetSettings(value []Settingable)()
}
