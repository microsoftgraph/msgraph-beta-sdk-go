package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementDomainJoinConnector struct {
    Entity
    // The connector display name.
    displayName *string;
    // Last time connector contacted Intune.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The connector state. Possible values are: active, error, inactive.
    state *DeviceManagementDomainJoinConnectorState;
    // The version of the connector.
    version *string;
}
// Instantiates a new deviceManagementDomainJoinConnector and sets the default values.
func NewDeviceManagementDomainJoinConnector()(*DeviceManagementDomainJoinConnector) {
    m := &DeviceManagementDomainJoinConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// Gets the state property value. The connector state. Possible values are: active, error, inactive.
func (m *DeviceManagementDomainJoinConnector) GetState()(*DeviceManagementDomainJoinConnectorState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceManagementDomainJoinConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseDeviceManagementDomainJoinConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementDomainJoinConnectorState)
            m.SetState(&cast)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementDomainJoinConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementDomainJoinConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The connector display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementDomainJoinConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastConnectionDateTime property value. Last time connector contacted Intune.
// Parameters:
//  - value : Value to set for the lastConnectionDateTime property.
func (m *DeviceManagementDomainJoinConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// Sets the state property value. The connector state. Possible values are: active, error, inactive.
// Parameters:
//  - value : Value to set for the state property.
func (m *DeviceManagementDomainJoinConnector) SetState(value *DeviceManagementDomainJoinConnectorState)() {
    m.state = value
}
// Sets the version property value. The version of the connector.
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceManagementDomainJoinConnector) SetVersion(value *string)() {
    m.version = value
}
