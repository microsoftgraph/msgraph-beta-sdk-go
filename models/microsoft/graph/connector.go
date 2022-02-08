package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Connector 
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
// NewConnector instantiates a new connector and sets the default values.
func NewConnector()(*Connector) {
    m := &Connector{
        Entity: *NewEntity(),
    }
    return m
}
// GetExternalIp gets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
func (m *Connector) GetExternalIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalIp
    }
}
// GetMachineName gets the machineName property value. The machine name the connector is installed and running on.
func (m *Connector) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
// GetMemberOf gets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
func (m *Connector) GetMemberOf()([]ConnectorGroup) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// GetStatus gets the status property value. Indicates the status of the connector. Possible values are: active, inactive. Read-only.
func (m *Connector) GetStatus()(*ConnectorStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Connector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["externalIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIp(val)
        }
        return nil
    }
    res["machineName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineName(val)
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectorGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConnectorGroup))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectorStatus))
        }
        return nil
    }
    return res
}
func (m *Connector) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetMemberOf() != nil {
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
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExternalIp sets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
func (m *Connector) SetExternalIp(value *string)() {
    if m != nil {
        m.externalIp = value
    }
}
// SetMachineName sets the machineName property value. The machine name the connector is installed and running on.
func (m *Connector) SetMachineName(value *string)() {
    if m != nil {
        m.machineName = value
    }
}
// SetMemberOf sets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
func (m *Connector) SetMemberOf(value []ConnectorGroup)() {
    if m != nil {
        m.memberOf = value
    }
}
// SetStatus sets the status property value. Indicates the status of the connector. Possible values are: active, inactive. Read-only.
func (m *Connector) SetStatus(value *ConnectorStatus)() {
    if m != nil {
        m.status = value
    }
}
