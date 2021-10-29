package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NdesConnector struct {
    Entity
    // The friendly name of the Ndes Connector.
    displayName *string;
    // Last connection time for the Ndes Connector
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Ndes Connector Status. Possible values are: none, active, inactive.
    state *NdesConnectorState;
}
// Instantiates a new ndesConnector and sets the default values.
func NewNdesConnector()(*NdesConnector) {
    m := &NdesConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// Gets the state property value. Ndes Connector Status. Possible values are: none, active, inactive.
func (m *NdesConnector) GetState()(*NdesConnectorState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *NdesConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastConnectionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastConnectionDateTime(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNdesConnectorState)
        if err != nil {
            return err
        }
        cast := val.(NdesConnectorState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *NdesConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The friendly name of the Ndes Connector.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *NdesConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
// Parameters:
//  - value : Value to set for the lastConnectionDateTime property.
func (m *NdesConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// Sets the state property value. Ndes Connector Status. Possible values are: none, active, inactive.
// Parameters:
//  - value : Value to set for the state property.
func (m *NdesConnector) SetState(value *NdesConnectorState)() {
    m.state = value
}
