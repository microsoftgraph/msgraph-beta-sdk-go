package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantContactInformationable 
type TenantContactInformationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEmail()(*string)
    GetName()(*string)
    GetNotes()(*string)
    GetPhone()(*string)
    GetTitle()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetPhone(value *string)()
    SetTitle(value *string)()
}
