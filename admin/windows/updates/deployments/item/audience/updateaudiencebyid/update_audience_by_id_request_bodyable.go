package updateaudiencebyid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateAudienceByIdRequestBodyable 
type UpdateAudienceByIdRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddExclusions()([]string)
    GetAddMembers()([]string)
    GetMemberEntityType()(*string)
    GetRemoveExclusions()([]string)
    GetRemoveMembers()([]string)
    SetAddExclusions(value []string)()
    SetAddMembers(value []string)()
    SetMemberEntityType(value *string)()
    SetRemoveExclusions(value []string)()
    SetRemoveMembers(value []string)()
}
