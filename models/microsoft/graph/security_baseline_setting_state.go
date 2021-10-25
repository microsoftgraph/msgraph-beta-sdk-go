package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityBaselineSettingState struct {
    Entity
    contributingPolicies []SecurityBaselineContributingPolicy;
    errorCode *string;
    settingCategoryId *string;
    settingCategoryName *string;
    settingId *string;
    settingName *string;
    sourcePolicies []SettingSource;
    state *SecurityBaselineComplianceState;
}
func NewSecurityBaselineSettingState()(*SecurityBaselineSettingState) {
    m := &SecurityBaselineSettingState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SecurityBaselineSettingState) GetContributingPolicies()([]SecurityBaselineContributingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.contributingPolicies
    }
}
func (m *SecurityBaselineSettingState) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *SecurityBaselineSettingState) GetSettingCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryId
    }
}
func (m *SecurityBaselineSettingState) GetSettingCategoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategoryName
    }
}
func (m *SecurityBaselineSettingState) GetSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingId
    }
}
func (m *SecurityBaselineSettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *SecurityBaselineSettingState) GetSourcePolicies()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.sourcePolicies
    }
}
func (m *SecurityBaselineSettingState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *SecurityBaselineSettingState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityBaselineContributingPolicy() })
        if err != nil {
            return err
        }
        res := make([]SecurityBaselineContributingPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityBaselineContributingPolicy))
        }
        m.SetContributingPolicies(res)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["settingCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingCategoryId(val)
        return nil
    }
    res["settingCategoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingCategoryName(val)
        return nil
    }
    res["settingId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingId(val)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    res["sourcePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        res := make([]SettingSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingSource))
        }
        m.SetSourcePolicies(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselineComplianceState)
        if err != nil {
            return err
        }
        cast := val.(SecurityBaselineComplianceState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *SecurityBaselineSettingState) IsNil()(bool) {
    return m == nil
}
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
func (m *SecurityBaselineSettingState) SetContributingPolicies(value []SecurityBaselineContributingPolicy)() {
    m.contributingPolicies = value
}
func (m *SecurityBaselineSettingState) SetErrorCode(value *string)() {
    m.errorCode = value
}
func (m *SecurityBaselineSettingState) SetSettingCategoryId(value *string)() {
    m.settingCategoryId = value
}
func (m *SecurityBaselineSettingState) SetSettingCategoryName(value *string)() {
    m.settingCategoryName = value
}
func (m *SecurityBaselineSettingState) SetSettingId(value *string)() {
    m.settingId = value
}
func (m *SecurityBaselineSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
func (m *SecurityBaselineSettingState) SetSourcePolicies(value []SettingSource)() {
    m.sourcePolicies = value
}
func (m *SecurityBaselineSettingState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
