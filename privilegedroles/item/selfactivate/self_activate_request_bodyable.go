package selfactivate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SelfActivateRequestBodyable 
type SelfActivateRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDuration()(*string)
    GetReason()(*string)
    GetTicketNumber()(*string)
    GetTicketSystem()(*string)
    SetDuration(value *string)()
    SetReason(value *string)()
    SetTicketNumber(value *string)()
    SetTicketSystem(value *string)()
}
