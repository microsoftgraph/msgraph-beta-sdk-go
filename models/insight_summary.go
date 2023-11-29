package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightSummary 
type InsightSummary struct {
    Entity
}
// NewInsightSummary instantiates a new insightSummary and sets the default values.
func NewInsightSummary()(*InsightSummary) {
    m := &InsightSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInsightSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightSummary(), nil
}
// GetActiveUsers gets the activeUsers property value. Daily active users.
func (m *InsightSummary) GetActiveUsers()(*int64) {
    val, err := m.GetBackingStore().Get("activeUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetAppId gets the appId property value. The ID of the Microsoft Entra application.
func (m *InsightSummary) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationCompletions gets the authenticationCompletions property value. Daily authentication completions.
func (m *InsightSummary) GetAuthenticationCompletions()(*int64) {
    val, err := m.GetBackingStore().Get("authenticationCompletions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetAuthenticationRequests gets the authenticationRequests property value. Daily authentication requests.
func (m *InsightSummary) GetAuthenticationRequests()(*int64) {
    val, err := m.GetBackingStore().Get("authenticationRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFactDate gets the factDate property value. The date of the insight.
func (m *InsightSummary) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *InsightSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveUsers(val)
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
    res["authenticationCompletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationCompletions(val)
        }
        return nil
    }
    res["authenticationRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRequests(val)
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
    res["securityTextCompletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityTextCompletions(val)
        }
        return nil
    }
    res["securityTextRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityTextRequests(val)
        }
        return nil
    }
    res["securityVoiceCompletions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityVoiceCompletions(val)
        }
        return nil
    }
    res["securityVoiceRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityVoiceRequests(val)
        }
        return nil
    }
    return res
}
// GetOs gets the os property value. The platform for the device that the customers used. Supports $filter (eq).
func (m *InsightSummary) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecurityTextCompletions gets the securityTextCompletions property value. Daily MFA SMS completions.
func (m *InsightSummary) GetSecurityTextCompletions()(*int64) {
    val, err := m.GetBackingStore().Get("securityTextCompletions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSecurityTextRequests gets the securityTextRequests property value. Daily MFA SMS requests.
func (m *InsightSummary) GetSecurityTextRequests()(*int64) {
    val, err := m.GetBackingStore().Get("securityTextRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSecurityVoiceCompletions gets the securityVoiceCompletions property value. Daily MFA Voice completions.
func (m *InsightSummary) GetSecurityVoiceCompletions()(*int64) {
    val, err := m.GetBackingStore().Get("securityVoiceCompletions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSecurityVoiceRequests gets the securityVoiceRequests property value. Daily MFA Voice requests.
func (m *InsightSummary) GetSecurityVoiceRequests()(*int64) {
    val, err := m.GetBackingStore().Get("securityVoiceRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InsightSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("activeUsers", m.GetActiveUsers())
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
        err = writer.WriteInt64Value("authenticationCompletions", m.GetAuthenticationCompletions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("authenticationRequests", m.GetAuthenticationRequests())
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
        err = writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("securityTextCompletions", m.GetSecurityTextCompletions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("securityTextRequests", m.GetSecurityTextRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("securityVoiceCompletions", m.GetSecurityVoiceCompletions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("securityVoiceRequests", m.GetSecurityVoiceRequests())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveUsers sets the activeUsers property value. Daily active users.
func (m *InsightSummary) SetActiveUsers(value *int64)() {
    err := m.GetBackingStore().Set("activeUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. The ID of the Microsoft Entra application.
func (m *InsightSummary) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationCompletions sets the authenticationCompletions property value. Daily authentication completions.
func (m *InsightSummary) SetAuthenticationCompletions(value *int64)() {
    err := m.GetBackingStore().Set("authenticationCompletions", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationRequests sets the authenticationRequests property value. Daily authentication requests.
func (m *InsightSummary) SetAuthenticationRequests(value *int64)() {
    err := m.GetBackingStore().Set("authenticationRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The date of the insight.
func (m *InsightSummary) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The platform for the device that the customers used. Supports $filter (eq).
func (m *InsightSummary) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityTextCompletions sets the securityTextCompletions property value. Daily MFA SMS completions.
func (m *InsightSummary) SetSecurityTextCompletions(value *int64)() {
    err := m.GetBackingStore().Set("securityTextCompletions", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityTextRequests sets the securityTextRequests property value. Daily MFA SMS requests.
func (m *InsightSummary) SetSecurityTextRequests(value *int64)() {
    err := m.GetBackingStore().Set("securityTextRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityVoiceCompletions sets the securityVoiceCompletions property value. Daily MFA Voice completions.
func (m *InsightSummary) SetSecurityVoiceCompletions(value *int64)() {
    err := m.GetBackingStore().Set("securityVoiceCompletions", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityVoiceRequests sets the securityVoiceRequests property value. Daily MFA Voice requests.
func (m *InsightSummary) SetSecurityVoiceRequests(value *int64)() {
    err := m.GetBackingStore().Set("securityVoiceRequests", value)
    if err != nil {
        panic(err)
    }
}
// InsightSummaryable 
type InsightSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveUsers()(*int64)
    GetAppId()(*string)
    GetAuthenticationCompletions()(*int64)
    GetAuthenticationRequests()(*int64)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetOs()(*string)
    GetSecurityTextCompletions()(*int64)
    GetSecurityTextRequests()(*int64)
    GetSecurityVoiceCompletions()(*int64)
    GetSecurityVoiceRequests()(*int64)
    SetActiveUsers(value *int64)()
    SetAppId(value *string)()
    SetAuthenticationCompletions(value *int64)()
    SetAuthenticationRequests(value *int64)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetOs(value *string)()
    SetSecurityTextCompletions(value *int64)()
    SetSecurityTextRequests(value *int64)()
    SetSecurityVoiceCompletions(value *int64)()
    SetSecurityVoiceRequests(value *int64)()
}
