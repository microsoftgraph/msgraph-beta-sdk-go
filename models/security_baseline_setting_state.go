package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineSettingState the security baseline compliance state of a setting for a device
type SecurityBaselineSettingState struct {
    Entity
}
// NewSecurityBaselineSettingState instantiates a new SecurityBaselineSettingState and sets the default values.
func NewSecurityBaselineSettingState()(*SecurityBaselineSettingState) {
    m := &SecurityBaselineSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityBaselineSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineSettingState(), nil
}
// GetContributingPolicies gets the contributingPolicies property value. The policies that contribute to this setting instance
// returns a []SecurityBaselineContributingPolicyable when successful
func (m *SecurityBaselineSettingState) GetContributingPolicies()([]SecurityBaselineContributingPolicyable) {
    val, err := m.GetBackingStore().Get("contributingPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineContributingPolicyable)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. The error code if the setting is in error state
// returns a *string when successful
func (m *SecurityBaselineSettingState) GetErrorCode()(*string) {
    val, err := m.GetBackingStore().Get("errorCode")
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
func (m *SecurityBaselineSettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineContributingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineContributingPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityBaselineContributingPolicyable)
                }
            }
            m.SetContributingPolicies(res)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["settingCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryId(val)
        }
        return nil
    }
    res["settingCategoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryName(val)
        }
        return nil
    }
    res["settingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["sourcePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SettingSourceable)
                }
            }
            m.SetSourcePolicies(res)
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
    return res
}
// GetSettingCategoryId gets the settingCategoryId property value. The setting category id which this setting belongs to
// returns a *string when successful
func (m *SecurityBaselineSettingState) GetSettingCategoryId()(*string) {
    val, err := m.GetBackingStore().Get("settingCategoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingCategoryName gets the settingCategoryName property value. The setting category name which this setting belongs to
// returns a *string when successful
func (m *SecurityBaselineSettingState) GetSettingCategoryName()(*string) {
    val, err := m.GetBackingStore().Get("settingCategoryName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingId gets the settingId property value. The setting id guid
// returns a *string when successful
func (m *SecurityBaselineSettingState) GetSettingId()(*string) {
    val, err := m.GetBackingStore().Get("settingId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingName gets the settingName property value. The setting name that is being reported
// returns a *string when successful
func (m *SecurityBaselineSettingState) GetSettingName()(*string) {
    val, err := m.GetBackingStore().Get("settingName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSourcePolicies gets the sourcePolicies property value. The policies that contribute to this setting instance
// returns a []SettingSourceable when successful
func (m *SecurityBaselineSettingState) GetSourcePolicies()([]SettingSourceable) {
    val, err := m.GetBackingStore().Get("sourcePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SettingSourceable)
    }
    return nil
}
// GetState gets the state property value. Security Baseline Compliance State
// returns a *SecurityBaselineComplianceState when successful
func (m *SecurityBaselineSettingState) GetState()(*SecurityBaselineComplianceState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityBaselineComplianceState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContributingPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContributingPolicies()))
        for i, v := range m.GetContributingPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contributingPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategoryId", m.GetSettingCategoryId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategoryName", m.GetSettingCategoryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingId", m.GetSettingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetSourcePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSourcePolicies()))
        for i, v := range m.GetSourcePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sourcePolicies", cast)
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
    return nil
}
// SetContributingPolicies sets the contributingPolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) SetContributingPolicies(value []SecurityBaselineContributingPolicyable)() {
    err := m.GetBackingStore().Set("contributingPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) SetErrorCode(value *string)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingCategoryId sets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryId(value *string)() {
    err := m.GetBackingStore().Set("settingCategoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingCategoryName sets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryName(value *string)() {
    err := m.GetBackingStore().Set("settingCategoryName", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingId sets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) SetSettingId(value *string)() {
    err := m.GetBackingStore().Set("settingId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingName sets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) SetSettingName(value *string)() {
    err := m.GetBackingStore().Set("settingName", value)
    if err != nil {
        panic(err)
    }
}
// SetSourcePolicies sets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) SetSourcePolicies(value []SettingSourceable)() {
    err := m.GetBackingStore().Set("sourcePolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineSettingState) SetState(value *SecurityBaselineComplianceState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
type SecurityBaselineSettingStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContributingPolicies()([]SecurityBaselineContributingPolicyable)
    GetErrorCode()(*string)
    GetSettingCategoryId()(*string)
    GetSettingCategoryName()(*string)
    GetSettingId()(*string)
    GetSettingName()(*string)
    GetSourcePolicies()([]SettingSourceable)
    GetState()(*SecurityBaselineComplianceState)
    SetContributingPolicies(value []SecurityBaselineContributingPolicyable)()
    SetErrorCode(value *string)()
    SetSettingCategoryId(value *string)()
    SetSettingCategoryName(value *string)()
    SetSettingId(value *string)()
    SetSettingName(value *string)()
    SetSourcePolicies(value []SettingSourceable)()
    SetState(value *SecurityBaselineComplianceState)()
}
