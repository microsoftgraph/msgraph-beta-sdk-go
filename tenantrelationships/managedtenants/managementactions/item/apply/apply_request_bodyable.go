package apply

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplyRequestBodyable 
type ApplyRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludeGroups()([]string)
    GetIncludeAllUsers()(*bool)
    GetIncludeGroups()([]string)
    GetManagementTemplateId()(*string)
    GetTenantGroupId()(*string)
    GetTenantId()(*string)
    SetExcludeGroups(value []string)()
    SetIncludeAllUsers(value *bool)()
    SetIncludeGroups(value []string)()
    SetManagementTemplateId(value *string)()
    SetTenantGroupId(value *string)()
    SetTenantId(value *string)()
}
