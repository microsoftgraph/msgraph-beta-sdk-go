package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new securityBaselineSettingState and sets the default values.
func NewSecurityBaselineSettingState()(*SecurityBaselineSettingState) {
    m := &SecurityBaselineSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the contributingPolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetContributingPolicies()([]SecurityBaselineContributingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.contributingPolicies
    }
}
// Gets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryId
    }
}
// Gets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryName
    }
}
// Gets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) GetSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingId
    }
}
// Gets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// Gets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetSourcePolicies()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.sourcePolicies
    }
}
// Gets the state property value. The compliance state of the security baseline setting
func (m *SecurityBaselineSettingState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
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
            cast := val.(SecurityBaselineComplianceState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *SecurityBaselineSettingState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecurityBaselineSettingState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the contributingPolicies property value. The policies that contribute to this setting instance
// Parameters:
//  - value : Value to set for the contributingPolicies property.
func (m *SecurityBaselineSettingState) SetContributingPolicies(value []SecurityBaselineContributingPolicy)() {
    m.contributingPolicies = value
}
// Sets the errorCode property value. The error code if the setting is in error state
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *SecurityBaselineSettingState) SetErrorCode(value *string)() {
    m.errorCode = value
}
// Sets the settingCategoryId property value. The setting category id which this setting belongs to
// Parameters:
//  - value : Value to set for the settingCategoryId property.
func (m *SecurityBaselineSettingState) SetSettingCategoryId(value *string)() {
    m.settingCategoryId = value
}
// Sets the settingCategoryName property value. The setting category name which this setting belongs to
// Parameters:
//  - value : Value to set for the settingCategoryName property.
func (m *SecurityBaselineSettingState) SetSettingCategoryName(value *string)() {
    m.settingCategoryName = value
}
// Sets the settingId property value. The setting id guid
// Parameters:
//  - value : Value to set for the settingId property.
func (m *SecurityBaselineSettingState) SetSettingId(value *string)() {
    m.settingId = value
}
// Sets the settingName property value. The setting name that is being reported
// Parameters:
//  - value : Value to set for the settingName property.
func (m *SecurityBaselineSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// Sets the sourcePolicies property value. The policies that contribute to this setting instance
// Parameters:
//  - value : Value to set for the sourcePolicies property.
func (m *SecurityBaselineSettingState) SetSourcePolicies(value []SettingSource)() {
    m.sourcePolicies = value
}
// Sets the state property value. The compliance state of the security baseline setting
// Parameters:
//  - value : Value to set for the state property.
func (m *SecurityBaselineSettingState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
