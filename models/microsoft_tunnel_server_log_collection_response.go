package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServerLogCollectionResponse entity that stores the server log collection status.
type MicrosoftTunnelServerLogCollectionResponse struct {
    Entity
}
// NewMicrosoftTunnelServerLogCollectionResponse instantiates a new MicrosoftTunnelServerLogCollectionResponse and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponse()(*MicrosoftTunnelServerLogCollectionResponse) {
    m := &MicrosoftTunnelServerLogCollectionResponse{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelServerLogCollectionResponse(), nil
}
// GetEndDateTime gets the endDateTime property value. The end time of the logs collected
// returns a *Time when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExpiryDateTime gets the expiryDateTime property value. The time when the log collection is expired
// returns a *Time when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expiryDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["expiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDateTime(val)
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["serverId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerId(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftTunnelLogCollectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MicrosoftTunnelLogCollectionStatus))
        }
        return nil
    }
    return res
}
// GetRequestDateTime gets the requestDateTime property value. The time when the log collection was requested
// returns a *Time when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetServerId gets the serverId property value. ID of the server the log collection is requested upon
// returns a *string when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetServerId()(*string) {
    val, err := m.GetBackingStore().Get("serverId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSizeInBytes gets the sizeInBytes property value. The size of the logs in bytes
// returns a *int64 when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetSizeInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sizeInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The start time of the logs collected
// returns a *Time when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. Enum type that represent the status of log collection
// returns a *MicrosoftTunnelLogCollectionStatus when successful
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStatus()(*MicrosoftTunnelLogCollectionStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftTunnelLogCollectionStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelServerLogCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverId", m.GetServerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// SetEndDateTime sets the endDateTime property value. The end time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExpiryDateTime sets the expiryDateTime property value. The time when the log collection is expired
func (m *MicrosoftTunnelServerLogCollectionResponse) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expiryDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestDateTime sets the requestDateTime property value. The time when the log collection was requested
func (m *MicrosoftTunnelServerLogCollectionResponse) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetServerId sets the serverId property value. ID of the server the log collection is requested upon
func (m *MicrosoftTunnelServerLogCollectionResponse) SetServerId(value *string)() {
    err := m.GetBackingStore().Set("serverId", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInBytes sets the sizeInBytes property value. The size of the logs in bytes
func (m *MicrosoftTunnelServerLogCollectionResponse) SetSizeInBytes(value *int64)() {
    err := m.GetBackingStore().Set("sizeInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The start time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Enum type that represent the status of log collection
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStatus(value *MicrosoftTunnelLogCollectionStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type MicrosoftTunnelServerLogCollectionResponseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetServerId()(*string)
    GetSizeInBytes()(*int64)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*MicrosoftTunnelLogCollectionStatus)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetServerId(value *string)()
    SetSizeInBytes(value *int64)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *MicrosoftTunnelLogCollectionStatus)()
}
