package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectorGroupable 
type ConnectorGroupable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplications()([]Applicationable)
    GetConnectorGroupType()(*ConnectorGroupType)
    GetIsDefault()(*bool)
    GetMembers()([]Connectorable)
    GetName()(*string)
    GetRegion()(*ConnectorGroupRegion)
    SetApplications(value []Applicationable)()
    SetConnectorGroupType(value *ConnectorGroupType)()
    SetIsDefault(value *bool)()
    SetMembers(value []Connectorable)()
    SetName(value *string)()
    SetRegion(value *ConnectorGroupRegion)()
}
