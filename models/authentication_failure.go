package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationFailure struct {
    Entity
}
// NewAuthenticationFailure instantiates a new AuthenticationFailure and sets the default values.
func NewAuthenticationFailure()(*AuthenticationFailure) {
    m := &AuthenticationFailure{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationFailureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationFailureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationFailure(), nil
}
// GetCount gets the count property value. The count property
// returns a *int64 when successful
func (m *AuthenticationFailure) GetCount()(*int64) {
    val, err := m.GetBackingStore().Get("count")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationFailure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["reasonCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationFailureReasonCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReasonCode(val.(*AuthenticationFailureReasonCode))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *AuthenticationFailure) GetReason()(*string) {
    val, err := m.GetBackingStore().Get("reason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReasonCode gets the reasonCode property value. The reasonCode property
// returns a *AuthenticationFailureReasonCode when successful
func (m *AuthenticationFailure) GetReasonCode()(*AuthenticationFailureReasonCode) {
    val, err := m.GetBackingStore().Get("reasonCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationFailureReasonCode)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationFailure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetReasonCode() != nil {
        cast := (*m.GetReasonCode()).String()
        err = writer.WriteStringValue("reasonCode", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCount sets the count property value. The count property
func (m *AuthenticationFailure) SetCount(value *int64)() {
    err := m.GetBackingStore().Set("count", value)
    if err != nil {
        panic(err)
    }
}
// SetReason sets the reason property value. The reason property
func (m *AuthenticationFailure) SetReason(value *string)() {
    err := m.GetBackingStore().Set("reason", value)
    if err != nil {
        panic(err)
    }
}
// SetReasonCode sets the reasonCode property value. The reasonCode property
func (m *AuthenticationFailure) SetReasonCode(value *AuthenticationFailureReasonCode)() {
    err := m.GetBackingStore().Set("reasonCode", value)
    if err != nil {
        panic(err)
    }
}
type AuthenticationFailureable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetReason()(*string)
    GetReasonCode()(*AuthenticationFailureReasonCode)
    SetCount(value *int64)()
    SetReason(value *string)()
    SetReasonCode(value *AuthenticationFailureReasonCode)()
}
