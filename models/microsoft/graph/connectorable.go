package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Connectorable 
type Connectorable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExternalIp()(*string)
    GetMachineName()(*string)
    GetMemberOf()([]ConnectorGroupable)
    GetStatus()(*ConnectorStatus)
    SetExternalIp(value *string)()
    SetMachineName(value *string)()
    SetMemberOf(value []ConnectorGroupable)()
    SetStatus(value *ConnectorStatus)()
}
