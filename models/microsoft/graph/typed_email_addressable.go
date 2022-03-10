package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TypedEmailAddressable 
type TypedEmailAddressable interface {
    EmailAddressable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOtherLabel()(*string)
    GetType()(*EmailType)
    SetOtherLabel(value *string)()
    SetType(value *EmailType)()
}
