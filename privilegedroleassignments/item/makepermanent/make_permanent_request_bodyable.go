package makepermanent

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MakePermanentRequestBodyable 
type MakePermanentRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetReason()(*string)
    GetTicketNumber()(*string)
    GetTicketSystem()(*string)
    SetReason(value *string)()
    SetTicketNumber(value *string)()
    SetTicketSystem(value *string)()
}
