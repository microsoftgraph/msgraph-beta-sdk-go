package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInSummary 
type ApplicationSignInSummary struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewApplicationSignInSummary instantiates a new applicationSignInSummary and sets the default values.
func NewApplicationSignInSummary()(*ApplicationSignInSummary) {
    m := &ApplicationSignInSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApplicationSignInSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationSignInSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationSignInSummary(), nil
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailedSignInCount gets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) GetFailedSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("failedSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["failedSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedSignInCount(val)
        }
        return nil
    }
    res["successfulSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulSignInCount(val)
        }
        return nil
    }
    res["successPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessPercentage(val)
        }
        return nil
    }
    return res
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessfulSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("successfulSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSuccessPercentage gets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessPercentage()(*float64) {
    val, err := m.GetBackingStore().Get("successPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApplicationSignInSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("failedSignInCount", m.GetFailedSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulSignInCount", m.GetSuccessfulSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("successPercentage", m.GetSuccessPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedSignInCount sets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) SetFailedSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("failedSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessfulSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("successfulSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessPercentage sets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessPercentage(value *float64)() {
    err := m.GetBackingStore().Set("successPercentage", value)
    if err != nil {
        panic(err)
    }
}
// ApplicationSignInSummaryable 
type ApplicationSignInSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetFailedSignInCount()(*int64)
    GetSuccessfulSignInCount()(*int64)
    GetSuccessPercentage()(*float64)
    SetAppDisplayName(value *string)()
    SetFailedSignInCount(value *int64)()
    SetSuccessfulSignInCount(value *int64)()
    SetSuccessPercentage(value *float64)()
}
