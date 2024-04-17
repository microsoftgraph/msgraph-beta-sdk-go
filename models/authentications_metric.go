package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationsMetric struct {
    Entity
}
// NewAuthenticationsMetric instantiates a new AuthenticationsMetric and sets the default values.
func NewAuthenticationsMetric()(*AuthenticationsMetric) {
    m := &AuthenticationsMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationsMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationsMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationsMetric(), nil
}
// GetAppid gets the appid property value. The ID of the Microsoft Entra application. Supports $filter (eq).
// returns a *string when successful
func (m *AuthenticationsMetric) GetAppid()(*string) {
    val, err := m.GetBackingStore().Get("appid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAttemptsCount gets the attemptsCount property value. The number of authentication requests made in the specified period. Supports $filter (eq).
// returns a *int64 when successful
func (m *AuthenticationsMetric) GetAttemptsCount()(*int64) {
    val, err := m.GetBackingStore().Get("attemptsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCountry gets the country property value. The location where the customers authenticated from. Supports $filter (eq).
// returns a *string when successful
func (m *AuthenticationsMetric) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFactDate gets the factDate property value. The date of the user insight.
// returns a *DateOnly when successful
func (m *AuthenticationsMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *AuthenticationsMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppid(val)
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
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
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
    res["identityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProvider(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
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
// GetIdentityProvider gets the identityProvider property value. The identityProvider property
// returns a *string when successful
func (m *AuthenticationsMetric) GetIdentityProvider()(*string) {
    val, err := m.GetBackingStore().Get("identityProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLanguage gets the language property value. The language property
// returns a *string when successful
func (m *AuthenticationsMetric) GetLanguage()(*string) {
    val, err := m.GetBackingStore().Get("language")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOs gets the os property value. The platform for the device that the customers used. Supports $filter (eq).
// returns a *string when successful
func (m *AuthenticationsMetric) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessCount gets the successCount property value. Number of successful authentication requests. Supports $filter (eq).
// returns a *int64 when successful
func (m *AuthenticationsMetric) GetSuccessCount()(*int64) {
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
func (m *AuthenticationsMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appid", m.GetAppid())
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
        err = writer.WriteStringValue("country", m.GetCountry())
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
        err = writer.WriteStringValue("identityProvider", m.GetIdentityProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
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
// SetAppid sets the appid property value. The ID of the Microsoft Entra application. Supports $filter (eq).
func (m *AuthenticationsMetric) SetAppid(value *string)() {
    err := m.GetBackingStore().Set("appid", value)
    if err != nil {
        panic(err)
    }
}
// SetAttemptsCount sets the attemptsCount property value. The number of authentication requests made in the specified period. Supports $filter (eq).
func (m *AuthenticationsMetric) SetAttemptsCount(value *int64)() {
    err := m.GetBackingStore().Set("attemptsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The location where the customers authenticated from. Supports $filter (eq).
func (m *AuthenticationsMetric) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The date of the user insight.
func (m *AuthenticationsMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityProvider sets the identityProvider property value. The identityProvider property
func (m *AuthenticationsMetric) SetIdentityProvider(value *string)() {
    err := m.GetBackingStore().Set("identityProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. The language property
func (m *AuthenticationsMetric) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The platform for the device that the customers used. Supports $filter (eq).
func (m *AuthenticationsMetric) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessCount sets the successCount property value. Number of successful authentication requests. Supports $filter (eq).
func (m *AuthenticationsMetric) SetSuccessCount(value *int64)() {
    err := m.GetBackingStore().Set("successCount", value)
    if err != nil {
        panic(err)
    }
}
type AuthenticationsMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppid()(*string)
    GetAttemptsCount()(*int64)
    GetCountry()(*string)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetIdentityProvider()(*string)
    GetLanguage()(*string)
    GetOs()(*string)
    GetSuccessCount()(*int64)
    SetAppid(value *string)()
    SetAttemptsCount(value *int64)()
    SetCountry(value *string)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetIdentityProvider(value *string)()
    SetLanguage(value *string)()
    SetOs(value *string)()
    SetSuccessCount(value *int64)()
}
