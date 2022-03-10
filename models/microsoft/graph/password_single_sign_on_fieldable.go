package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PasswordSingleSignOnFieldable 
type PasswordSingleSignOnFieldable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCustomizedLabel()(*string)
    GetDefaultLabel()(*string)
    GetFieldId()(*string)
    GetType()(*string)
    SetCustomizedLabel(value *string)()
    SetDefaultLabel(value *string)()
    SetFieldId(value *string)()
    SetType(value *string)()
}
