package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineState security baseline state for a device.
type SecurityBaselineState struct {
    Entity
}
// NewSecurityBaselineState instantiates a new SecurityBaselineState and sets the default values.
func NewSecurityBaselineState()(*SecurityBaselineState) {
    m := &SecurityBaselineState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityBaselineStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineState(), nil
}
// GetDisplayName gets the displayName property value. The display name of the security baseline
// returns a *string when successful
func (m *SecurityBaselineState) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityBaselineState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["securityBaselineTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBaselineTemplateId(val)
        }
        return nil
    }
    res["settingStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineSettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineSettingStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityBaselineSettingStateable)
                }
            }
            m.SetSettingStates(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselineComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecurityBaselineComplianceState))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetSecurityBaselineTemplateId gets the securityBaselineTemplateId property value. The security baseline template id
// returns a *string when successful
func (m *SecurityBaselineState) GetSecurityBaselineTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("securityBaselineTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingStates gets the settingStates property value. The security baseline state for different settings for a device
// returns a []SecurityBaselineSettingStateable when successful
func (m *SecurityBaselineState) GetSettingStates()([]SecurityBaselineSettingStateable) {
    val, err := m.GetBackingStore().Get("settingStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineSettingStateable)
    }
    return nil
}
// GetState gets the state property value. Security Baseline Compliance State
// returns a *SecurityBaselineComplianceState when successful
func (m *SecurityBaselineState) GetState()(*SecurityBaselineComplianceState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityBaselineComplianceState)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
// returns a *string when successful
func (m *SecurityBaselineState) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityBaselineTemplateId", m.GetSecurityBaselineTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetSettingStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the security baseline
func (m *SecurityBaselineState) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityBaselineTemplateId sets the securityBaselineTemplateId property value. The security baseline template id
func (m *SecurityBaselineState) SetSecurityBaselineTemplateId(value *string)() {
    err := m.GetBackingStore().Set("securityBaselineTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingStates sets the settingStates property value. The security baseline state for different settings for a device
func (m *SecurityBaselineState) SetSettingStates(value []SecurityBaselineSettingStateable)() {
    err := m.GetBackingStore().Set("settingStates", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineState) SetState(value *SecurityBaselineComplianceState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineState) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type SecurityBaselineStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetSecurityBaselineTemplateId()(*string)
    GetSettingStates()([]SecurityBaselineSettingStateable)
    GetState()(*SecurityBaselineComplianceState)
    GetUserPrincipalName()(*string)
    SetDisplayName(value *string)()
    SetSecurityBaselineTemplateId(value *string)()
    SetSettingStates(value []SecurityBaselineSettingStateable)()
    SetState(value *SecurityBaselineComplianceState)()
    SetUserPrincipalName(value *string)()
}
