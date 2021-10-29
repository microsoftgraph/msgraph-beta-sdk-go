package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Connector struct {
    Entity
    // The external IP address as detected by the the connector server. Read-only.
    externalIp *string;
    // The machine name the connector is installed and running on.
    machineName *string;
    // The connectorGroup that the connector is a member of. Read-only.
    memberOf []ConnectorGroup;
    // Indicates the status of the connector. Possible values are: active, inactive. Read-only.
    status *ConnectorStatus;
}
// Instantiates a new connector and sets the default values.
func NewConnector()(*Connector) {
    m := &Connector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
func (m *Connector) GetExternalIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalIp
    }
}
// Gets the machineName property value. The machine name the connector is installed and running on.
func (m *Connector) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
// Gets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
func (m *Connector) GetMemberOf()([]ConnectorGroup) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the status property value. Indicates the status of the connector. Possible values are: active, inactive. Read-only.
func (m *Connector) GetStatus()(*ConnectorStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *Connector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["externalIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalIp(val)
        return nil
    }
    res["machineName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMachineName(val)
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectorGroup() })
        if err != nil {
            return err
        }
        res := make([]ConnectorGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConnectorGroup))
        }
        m.SetMemberOf(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorStatus)
        if err != nil {
            return err
        }
        cast := val.(ConnectorStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *Connector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Connector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("externalIp", m.GetExternalIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
// Parameters:
//  - value : Value to set for the externalIp property.
func (m *Connector) SetExternalIp(value *string)() {
    m.externalIp = value
}
// Sets the machineName property value. The machine name the connector is installed and running on.
// Parameters:
//  - value : Value to set for the machineName property.
func (m *Connector) SetMachineName(value *string)() {
    m.machineName = value
}
// Sets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *Connector) SetMemberOf(value []ConnectorGroup)() {
    m.memberOf = value
}
// Sets the status property value. Indicates the status of the connector. Possible values are: active, inactive. Read-only.
// Parameters:
//  - value : Value to set for the status property.
func (m *Connector) SetStatus(value *ConnectorStatus)() {
    m.status = value
}
