package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcStatusDetailsable 
type CloudPcStatusDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalInformation()([]KeyValuePairable)
    GetCode()(*string)
    GetMessage()(*string)
    SetAdditionalInformation(value []KeyValuePairable)()
    SetCode(value *string)()
    SetMessage(value *string)()
}
