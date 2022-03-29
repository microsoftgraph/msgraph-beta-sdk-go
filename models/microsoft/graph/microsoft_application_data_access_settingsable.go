package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftApplicationDataAccessSettingsable 
type MicrosoftApplicationDataAccessSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisabledForGroup()(*string)
    GetIsEnabledForAllMicrosoftApplications()(*bool)
    SetDisabledForGroup(value *string)()
    SetIsEnabledForAllMicrosoftApplications(value *bool)()
}
