package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityBaselineState 
type SecurityBaselineState struct {
    Entity
    // The display name of the security baseline
    displayName *string;
    // The security baseline template id
    securityBaselineTemplateId *string;
    // The security baseline state for different settings for a device
    settingStates []SecurityBaselineSettingState;
    // Security baseline compliance state
    state *SecurityBaselineComplianceState;
    // User Principal Name
    userPrincipalName *string;
}
// NewSecurityBaselineState instantiates a new securityBaselineState and sets the default values.
func NewSecurityBaselineState()(*SecurityBaselineState) {
    m := &SecurityBaselineState{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The display name of the security baseline
func (m *SecurityBaselineState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
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
func (m *SecurityBaselineState) GetSettingStates()([]SecurityBaselineSettingState) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["securityBaselineTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBaselineTemplateId(val)
        }
        return nil
    }
    res["settingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityBaselineSettingState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineSettingState, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecurityBaselineSettingState))
            }
            m.SetSettingStates(res)
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
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *SecurityBaselineState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *SecurityBaselineState) SetSettingStates(value []SecurityBaselineSettingState)() {
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
