package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserSignUpMetric struct {
    Entity
}
// NewUserSignUpMetric instantiates a new UserSignUpMetric and sets the default values.
func NewUserSignUpMetric()(*UserSignUpMetric) {
    m := &UserSignUpMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSignUpMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserSignUpMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSignUpMetric(), nil
}
// GetAppId gets the appId property value. The appId property
// returns a *string when successful
func (m *UserSignUpMetric) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBrowser gets the browser property value. The browser property
// returns a *string when successful
func (m *UserSignUpMetric) GetBrowser()(*string) {
    val, err := m.GetBackingStore().Get("browser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCount gets the count property value. The total number of users who signed up in the specified period. Supports $filter (eq).
// returns a *int64 when successful
func (m *UserSignUpMetric) GetCount()(*int64) {
    val, err := m.GetBackingStore().Get("count")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *UserSignUpMetric) GetCountry()(*string) {
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
func (m *UserSignUpMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *UserSignUpMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["browser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowser(val)
        }
        return nil
    }
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
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
    return res
}
// GetIdentityProvider gets the identityProvider property value. The identityProvider property
// returns a *string when successful
func (m *UserSignUpMetric) GetIdentityProvider()(*string) {
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
func (m *UserSignUpMetric) GetLanguage()(*string) {
    val, err := m.GetBackingStore().Get("language")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOs gets the os property value. The device plaform that the customers used. Supports $filter (eq).
// returns a *string when successful
func (m *UserSignUpMetric) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserSignUpMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("count", m.GetCount())
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
    return nil
}
// SetAppId sets the appId property value. The appId property
func (m *UserSignUpMetric) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetBrowser sets the browser property value. The browser property
func (m *UserSignUpMetric) SetBrowser(value *string)() {
    err := m.GetBackingStore().Set("browser", value)
    if err != nil {
        panic(err)
    }
}
// SetCount sets the count property value. The total number of users who signed up in the specified period. Supports $filter (eq).
func (m *UserSignUpMetric) SetCount(value *int64)() {
    err := m.GetBackingStore().Set("count", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The country property
func (m *UserSignUpMetric) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The date of the user insight.
func (m *UserSignUpMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityProvider sets the identityProvider property value. The identityProvider property
func (m *UserSignUpMetric) SetIdentityProvider(value *string)() {
    err := m.GetBackingStore().Set("identityProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguage sets the language property value. The language property
func (m *UserSignUpMetric) SetLanguage(value *string)() {
    err := m.GetBackingStore().Set("language", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The device plaform that the customers used. Supports $filter (eq).
func (m *UserSignUpMetric) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
type UserSignUpMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetBrowser()(*string)
    GetCount()(*int64)
    GetCountry()(*string)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetIdentityProvider()(*string)
    GetLanguage()(*string)
    GetOs()(*string)
    SetAppId(value *string)()
    SetBrowser(value *string)()
    SetCount(value *int64)()
    SetCountry(value *string)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetIdentityProvider(value *string)()
    SetLanguage(value *string)()
    SetOs(value *string)()
}
