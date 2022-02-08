package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineSettingState 
type SecurityBaselineSettingState struct {
    Entity
    // The policies that contribute to this setting instance
    contributingPolicies []SecurityBaselineContributingPolicy;
    // The error code if the setting is in error state
    errorCode *string;
    // The setting category id which this setting belongs to
    settingCategoryId *string;
    // The setting category name which this setting belongs to
    settingCategoryName *string;
    // The setting id guid
    settingId *string;
    // The setting name that is being reported
    settingName *string;
    // The policies that contribute to this setting instance
    sourcePolicies []SettingSource;
    // The compliance state of the security baseline setting
    state *SecurityBaselineComplianceState;
}
// NewSecurityBaselineSettingState instantiates a new securityBaselineSettingState and sets the default values.
func NewSecurityBaselineSettingState()(*SecurityBaselineSettingState) {
    m := &SecurityBaselineSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// GetContributingPolicies gets the contributingPolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetContributingPolicies()([]SecurityBaselineContributingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.contributingPolicies
    }
}
// GetErrorCode gets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetSettingCategoryId gets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryId
    }
}
// GetSettingCategoryName gets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryName
    }
}
// GetSettingId gets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) GetSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingId
    }
}
// GetSettingName gets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetSourcePolicies gets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetSourcePolicies()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.sourcePolicies
    }
}
// GetState gets the state property value. The compliance state of the security baseline setting
func (m *SecurityBaselineSettingState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineSettingState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityBaselineContributingPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineContributingPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecurityBaselineContributingPolicy))
            }
            m.SetContributingPolicies(res)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["settingCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryId(val)
        }
        return nil
    }
    res["settingCategoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryName(val)
        }
        return nil
    }
    res["settingId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
        }
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["sourcePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingSource))
            }
            m.SetSourcePolicies(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *SecurityBaselineSettingState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineSettingState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContributingPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContributingPolicies()))
        for i, v := range m.GetContributingPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSourcePolicies()))
        for i, v := range m.GetSourcePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *SecurityBaselineSettingState) SetContributingPolicies(value []SecurityBaselineContributingPolicy)() {
    if m != nil {
        m.contributingPolicies = value
    }
}
// SetErrorCode sets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) SetErrorCode(value *string)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetSettingCategoryId sets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryId(value *string)() {
    if m != nil {
        m.settingCategoryId = value
    }
}
// SetSettingCategoryName sets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryName(value *string)() {
    if m != nil {
        m.settingCategoryName = value
    }
}
// SetSettingId sets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) SetSettingId(value *string)() {
    if m != nil {
        m.settingId = value
    }
}
// SetSettingName sets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetSourcePolicies sets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) SetSourcePolicies(value []SettingSource)() {
    if m != nil {
        m.sourcePolicies = value
    }
}
// SetState sets the state property value. The compliance state of the security baseline setting
func (m *SecurityBaselineSettingState) SetState(value *SecurityBaselineComplianceState)() {
    if m != nil {
        m.state = value
    }
}
