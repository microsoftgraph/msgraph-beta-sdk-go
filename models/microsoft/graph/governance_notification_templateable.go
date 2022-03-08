package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceNotificationTemplateable 
type GovernanceNotificationTemplateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCulture()(*string)
    GetId()(*string)
    GetSource()(*string)
    GetType()(*string)
    GetVersion()(*string)
    SetCulture(value *string)()
    SetId(value *string)()
    SetSource(value *string)()
    SetType(value *string)()
    SetVersion(value *string)()
}
