package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MfaCompletionMetric struct {
    Entity
}
// NewMfaCompletionMetric instantiates a new MfaCompletionMetric and sets the default values.
func NewMfaCompletionMetric()(*MfaCompletionMetric) {
    m := &MfaCompletionMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMfaCompletionMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMfaCompletionMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMfaCompletionMetric(), nil
}
// GetAppId gets the appId property value. The ID of the Microsoft Entra application. Supports $filter (eq).
// returns a *string when successful
func (m *MfaCompletionMetric) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAttemptsCount gets the attemptsCount property value. Number of users who attempted to sign up. Supports $filter (eq).
// returns a *int64 when successful
func (m *MfaCompletionMetric) GetAttemptsCount()(*int64) {
    val, err := m.GetBackingStore().Get("attemptsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFactDate gets the factDate property value. The date of the user insight.
// returns a *DateOnly when successful
func (m *MfaCompletionMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("factDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MfaCompletionMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["attemptsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttemptsCount(val)
        }
        return nil
    }
    res["factDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactDate(val)
        }
        return nil
    }
    res["mfaMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaMethod(val)
        }
        return nil
    }
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetMfaMethod gets the mfaMethod property value. The MFA authentication method used by the customers. Supports $filter (eq).
// returns a *string when successful
func (m *MfaCompletionMetric) GetMfaMethod()(*string) {
    val, err := m.GetBackingStore().Get("mfaMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOs gets the os property value. The platform of the device that the customers used. Supports $filter (eq).
// returns a *string when successful
func (m *MfaCompletionMetric) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessCount gets the successCount property value. Number of users who signed up successfully. Supports $filter (eq).
// returns a *int64 when successful
func (m *MfaCompletionMetric) GetSuccessCount()(*int64) {
    val, err := m.GetBackingStore().Get("successCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MfaCompletionMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("attemptsCount", m.GetAttemptsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("factDate", m.GetFactDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mfaMethod", m.GetMfaMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The ID of the Microsoft Entra application. Supports $filter (eq).
func (m *MfaCompletionMetric) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAttemptsCount sets the attemptsCount property value. Number of users who attempted to sign up. Supports $filter (eq).
func (m *MfaCompletionMetric) SetAttemptsCount(value *int64)() {
    err := m.GetBackingStore().Set("attemptsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The date of the user insight.
func (m *MfaCompletionMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaMethod sets the mfaMethod property value. The MFA authentication method used by the customers. Supports $filter (eq).
func (m *MfaCompletionMetric) SetMfaMethod(value *string)() {
    err := m.GetBackingStore().Set("mfaMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The platform of the device that the customers used. Supports $filter (eq).
func (m *MfaCompletionMetric) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessCount sets the successCount property value. Number of users who signed up successfully. Supports $filter (eq).
func (m *MfaCompletionMetric) SetSuccessCount(value *int64)() {
    err := m.GetBackingStore().Set("successCount", value)
    if err != nil {
        panic(err)
    }
}
type MfaCompletionMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAttemptsCount()(*int64)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetMfaMethod()(*string)
    GetOs()(*string)
    GetSuccessCount()(*int64)
    SetAppId(value *string)()
    SetAttemptsCount(value *int64)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetMfaMethod(value *string)()
    SetOs(value *string)()
    SetSuccessCount(value *int64)()
}
