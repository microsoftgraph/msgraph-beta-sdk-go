package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MfaTelecomFraudMetric struct {
    Entity
}
// NewMfaTelecomFraudMetric instantiates a new MfaTelecomFraudMetric and sets the default values.
func NewMfaTelecomFraudMetric()(*MfaTelecomFraudMetric) {
    m := &MfaTelecomFraudMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMfaTelecomFraudMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMfaTelecomFraudMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMfaTelecomFraudMetric(), nil
}
// GetCaptchaFailureCount gets the captchaFailureCount property value. The captchaFailureCount property
// returns a *int64 when successful
func (m *MfaTelecomFraudMetric) GetCaptchaFailureCount()(*int64) {
    val, err := m.GetBackingStore().Get("captchaFailureCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCaptchaNotTriggeredUserCount gets the captchaNotTriggeredUserCount property value. The captchaNotTriggeredUserCount property
// returns a *int64 when successful
func (m *MfaTelecomFraudMetric) GetCaptchaNotTriggeredUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("captchaNotTriggeredUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCaptchaShownUserCount gets the captchaShownUserCount property value. The captchaShownUserCount property
// returns a *int64 when successful
func (m *MfaTelecomFraudMetric) GetCaptchaShownUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("captchaShownUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCaptchaSuccessCount gets the captchaSuccessCount property value. The captchaSuccessCount property
// returns a *int64 when successful
func (m *MfaTelecomFraudMetric) GetCaptchaSuccessCount()(*int64) {
    val, err := m.GetBackingStore().Get("captchaSuccessCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFactDate gets the factDate property value. The factDate property
// returns a *DateOnly when successful
func (m *MfaTelecomFraudMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *MfaTelecomFraudMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["captchaFailureCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptchaFailureCount(val)
        }
        return nil
    }
    res["captchaNotTriggeredUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptchaNotTriggeredUserCount(val)
        }
        return nil
    }
    res["captchaShownUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptchaShownUserCount(val)
        }
        return nil
    }
    res["captchaSuccessCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptchaSuccessCount(val)
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
    res["telecomBlockedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTelecomBlockedUserCount(val)
        }
        return nil
    }
    return res
}
// GetTelecomBlockedUserCount gets the telecomBlockedUserCount property value. The telecomBlockedUserCount property
// returns a *int64 when successful
func (m *MfaTelecomFraudMetric) GetTelecomBlockedUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("telecomBlockedUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MfaTelecomFraudMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("captchaFailureCount", m.GetCaptchaFailureCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("captchaNotTriggeredUserCount", m.GetCaptchaNotTriggeredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("captchaShownUserCount", m.GetCaptchaShownUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("captchaSuccessCount", m.GetCaptchaSuccessCount())
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
        err = writer.WriteInt64Value("telecomBlockedUserCount", m.GetTelecomBlockedUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCaptchaFailureCount sets the captchaFailureCount property value. The captchaFailureCount property
func (m *MfaTelecomFraudMetric) SetCaptchaFailureCount(value *int64)() {
    err := m.GetBackingStore().Set("captchaFailureCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCaptchaNotTriggeredUserCount sets the captchaNotTriggeredUserCount property value. The captchaNotTriggeredUserCount property
func (m *MfaTelecomFraudMetric) SetCaptchaNotTriggeredUserCount(value *int64)() {
    err := m.GetBackingStore().Set("captchaNotTriggeredUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCaptchaShownUserCount sets the captchaShownUserCount property value. The captchaShownUserCount property
func (m *MfaTelecomFraudMetric) SetCaptchaShownUserCount(value *int64)() {
    err := m.GetBackingStore().Set("captchaShownUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCaptchaSuccessCount sets the captchaSuccessCount property value. The captchaSuccessCount property
func (m *MfaTelecomFraudMetric) SetCaptchaSuccessCount(value *int64)() {
    err := m.GetBackingStore().Set("captchaSuccessCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The factDate property
func (m *MfaTelecomFraudMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTelecomBlockedUserCount sets the telecomBlockedUserCount property value. The telecomBlockedUserCount property
func (m *MfaTelecomFraudMetric) SetTelecomBlockedUserCount(value *int64)() {
    err := m.GetBackingStore().Set("telecomBlockedUserCount", value)
    if err != nil {
        panic(err)
    }
}
type MfaTelecomFraudMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCaptchaFailureCount()(*int64)
    GetCaptchaNotTriggeredUserCount()(*int64)
    GetCaptchaShownUserCount()(*int64)
    GetCaptchaSuccessCount()(*int64)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTelecomBlockedUserCount()(*int64)
    SetCaptchaFailureCount(value *int64)()
    SetCaptchaNotTriggeredUserCount(value *int64)()
    SetCaptchaShownUserCount(value *int64)()
    SetCaptchaSuccessCount(value *int64)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTelecomBlockedUserCount(value *int64)()
}
