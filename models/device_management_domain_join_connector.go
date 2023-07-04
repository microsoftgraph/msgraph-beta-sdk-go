package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDomainJoinConnector a Domain Join Connector is a connector that is responsible to allocate (and delete) machine account blobs
type DeviceManagementDomainJoinConnector struct {
    Entity
}
// NewDeviceManagementDomainJoinConnector instantiates a new deviceManagementDomainJoinConnector and sets the default values.
func NewDeviceManagementDomainJoinConnector()(*DeviceManagementDomainJoinConnector) {
    m := &DeviceManagementDomainJoinConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDomainJoinConnector(), nil
}
// GetDisplayName gets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDomainJoinConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastConnectionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectionDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDomainJoinConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*DeviceManagementDomainJoinConnectorState))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastConnectionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetState gets the state property value. The ODJ request states.
func (m *DeviceManagementDomainJoinConnector) GetState()(*DeviceManagementDomainJoinConnectorState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementDomainJoinConnectorState)
    }
    return nil
}
// GetVersion gets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementDomainJoinConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastConnectionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The ODJ request states.
func (m *DeviceManagementDomainJoinConnector) SetState(value *DeviceManagementDomainJoinConnectorState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementDomainJoinConnectorable 
type DeviceManagementDomainJoinConnectorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetState()(*DeviceManagementDomainJoinConnectorState)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetState(value *DeviceManagementDomainJoinConnectorState)()
    SetVersion(value *string)()
}
