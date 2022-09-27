package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NdesConnector entity which represents an OnPrem Ndes connector.
type NdesConnector struct {
    Entity
    // The build version of the Ndes Connector.
    connectorVersion *string
    // The friendly name of the Ndes Connector.
    displayName *string
    // Timestamp when on-prem certificate connector was enrolled in Intune.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last connection time for the Ndes Connector
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the machine running on-prem certificate connector service.
    machineName *string
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The current status of the Ndes Connector.
    state *NdesConnectorState
}
// NewNdesConnector instantiates a new ndesConnector and sets the default values.
func NewNdesConnector()(*NdesConnector) {
    m := &NdesConnector{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ndesConnector";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateNdesConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNdesConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNdesConnector(), nil
}
// GetConnectorVersion gets the connectorVersion property value. The build version of the Ndes Connector.
func (m *NdesConnector) GetConnectorVersion()(*string) {
    return m.connectorVersion
}
// GetDisplayName gets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Timestamp when on-prem certificate connector was enrolled in Intune.
func (m *NdesConnector) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NdesConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectorVersion)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["enrolledDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEnrolledDateTime)
    res["lastConnectionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastConnectionDateTime)
    res["machineName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMachineName)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseNdesConnectorState , m.SetState)
    return res
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConnectionDateTime
}
// GetMachineName gets the machineName property value. Name of the machine running on-prem certificate connector service.
func (m *NdesConnector) GetMachineName()(*string) {
    return m.machineName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NdesConnector) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetState gets the state property value. The current status of the Ndes Connector.
func (m *NdesConnector) GetState()(*NdesConnectorState) {
    return m.state
}
// Serialize serializes information the current object
func (m *NdesConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrolledDateTime", m.GetEnrolledDateTime())
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
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
// SetConnectorVersion sets the connectorVersion property value. The build version of the Ndes Connector.
func (m *NdesConnector) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
// SetDisplayName sets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Timestamp when on-prem certificate connector was enrolled in Intune.
func (m *NdesConnector) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// SetMachineName sets the machineName property value. Name of the machine running on-prem certificate connector service.
func (m *NdesConnector) SetMachineName(value *string)() {
    m.machineName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NdesConnector) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetState sets the state property value. The current status of the Ndes Connector.
func (m *NdesConnector) SetState(value *NdesConnectorState)() {
    m.state = value
}
