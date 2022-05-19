package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineState security baseline state for a device.
type SecurityBaselineState struct {
    Entity
    // The display name of the security baseline
    displayName *string
    // The security baseline template id
    securityBaselineTemplateId *string
    // The security baseline state for different settings for a device
    settingStates []SecurityBaselineSettingStateable
    // Security baseline compliance state
    state *SecurityBaselineComplianceState
    // User Principal Name
    userPrincipalName *string
}
// NewSecurityBaselineState instantiates a new securityBaselineState and sets the default values.
func NewSecurityBaselineState()(*SecurityBaselineState) {
    m := &SecurityBaselineState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineState(), nil
}
// GetDisplayName gets the displayName property value. The display name of the security baseline
func (m *SecurityBaselineState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = v.(SecurityBaselineSettingStateable)
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
func (m *SecurityBaselineState) GetSecurityBaselineTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.securityBaselineTemplateId
    }
}
// GetSettingStates gets the settingStates property value. The security baseline state for different settings for a device
func (m *SecurityBaselineState) GetSettingStates()([]SecurityBaselineSettingStateable) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
// GetState gets the state property value. Security baseline compliance state
func (m *SecurityBaselineState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m != nil {
        m.displayName = value
    }
}
// SetSecurityBaselineTemplateId sets the securityBaselineTemplateId property value. The security baseline template id
func (m *SecurityBaselineState) SetSecurityBaselineTemplateId(value *string)() {
    if m != nil {
        m.securityBaselineTemplateId = value
    }
}
// SetSettingStates sets the settingStates property value. The security baseline state for different settings for a device
func (m *SecurityBaselineState) SetSettingStates(value []SecurityBaselineSettingStateable)() {
    if m != nil {
        m.settingStates = value
    }
}
// SetState sets the state property value. Security baseline compliance state
func (m *SecurityBaselineState) SetState(value *SecurityBaselineComplianceState)() {
    if m != nil {
        m.state = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
