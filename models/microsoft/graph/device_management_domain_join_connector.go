package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementDomainJoinConnector 
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
// NewDeviceManagementDomainJoinConnector instantiates a new deviceManagementDomainJoinConnector and sets the default values.
func NewDeviceManagementDomainJoinConnector()(*DeviceManagementDomainJoinConnector) {
    m := &DeviceManagementDomainJoinConnector{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// GetState gets the state property value. The connector state. Possible values are: active, error, inactive.
func (m *DeviceManagementDomainJoinConnector) GetState()(*DeviceManagementDomainJoinConnectorState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetVersion gets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetState(val.(*DeviceManagementDomainJoinConnectorState))
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
// Serialize serializes information the current object
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
        cast := (*m.GetState()).String()
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
// SetDisplayName sets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConnectionDateTime = value
    }
}
// SetState sets the state property value. The connector state. Possible values are: active, error, inactive.
func (m *DeviceManagementDomainJoinConnector) SetState(value *DeviceManagementDomainJoinConnectorState)() {
    if m != nil {
        m.state = value
    }
}
// SetVersion sets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
