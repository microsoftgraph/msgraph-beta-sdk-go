package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageTraceable 
type MessageTraceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDestinationIPAddress()(*string)
    GetMessageId()(*string)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipients()([]MessageRecipientable)
    GetSenderEmail()(*string)
    GetSize()(*int32)
    GetSourceIPAddress()(*string)
    GetSubject()(*string)
    SetDestinationIPAddress(value *string)()
    SetMessageId(value *string)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipients(value []MessageRecipientable)()
    SetSenderEmail(value *string)()
    SetSize(value *int32)()
    SetSourceIPAddress(value *string)()
    SetSubject(value *string)()
}
