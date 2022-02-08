package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NdesConnector 
type NdesConnector struct {
    Entity
    // The friendly name of the Ndes Connector.
    displayName *string;
    // Last connection time for the Ndes Connector
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Ndes Connector Status. Possible values are: none, active, inactive.
    state *NdesConnectorState;
}
// NewNdesConnector instantiates a new ndesConnector and sets the default values.
func NewNdesConnector()(*NdesConnector) {
    m := &NdesConnector{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// GetState gets the state property value. Ndes Connector Status. Possible values are: none, active, inactive.
func (m *NdesConnector) GetState()(*NdesConnectorState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NdesConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastConnectionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectionDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNdesConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*NdesConnectorState))
        }
        return nil
    }
    return res
}
func (m *NdesConnector) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *NdesConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastConnectionDateTime", m.GetLastConnectionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConnectionDateTime = value
    }
}
// SetState sets the state property value. Ndes Connector Status. Possible values are: none, active, inactive.
func (m *NdesConnector) SetState(value *NdesConnectorState)() {
    if m != nil {
        m.state = value
    }
}
