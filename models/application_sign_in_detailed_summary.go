package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInDetailedSummary 
type ApplicationSignInDetailedSummary struct {
    Entity
}
// NewApplicationSignInDetailedSummary instantiates a new applicationSignInDetailedSummary and sets the default values.
func NewApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummary) {
    m := &ApplicationSignInDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApplicationSignInDetailedSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationSignInDetailedSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationSignInDetailedSummary(), nil
}
// GetAggregatedEventDateTime gets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("aggregatedEventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppId gets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInDetailedSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedEventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregatedEventDateTime(val)
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["signInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInCount(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignInStatusable))
        }
        return nil
    }
    return res
}
// GetSignInCount gets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) GetSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("signInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetStatus gets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) GetStatus()(SignInStatusable) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApplicationSignInDetailedSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("aggregatedEventDateTime", m.GetAggregatedEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("signInCount", m.GetSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedEventDateTime sets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("aggregatedEventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInCount sets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) SetSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("signInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) SetStatus(value SignInStatusable)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// ApplicationSignInDetailedSummaryable 
type ApplicationSignInDetailedSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetSignInCount()(*int64)
    GetStatus()(SignInStatusable)
    SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetSignInCount(value *int64)()
    SetStatus(value SignInStatusable)()
}
