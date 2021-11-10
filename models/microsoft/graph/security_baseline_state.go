package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new securityBaselineState and sets the default values.
func NewSecurityBaselineState()(*SecurityBaselineState) {
    m := &SecurityBaselineState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the security baseline
func (m *SecurityBaselineState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the securityBaselineTemplateId property value. The security baseline template id
func (m *SecurityBaselineState) GetSecurityBaselineTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.securityBaselineTemplateId
    }
}
// Gets the settingStates property value. The security baseline state for different settings for a device
func (m *SecurityBaselineState) GetSettingStates()([]SecurityBaselineSettingState) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
// Gets the state property value. Security baseline compliance state
func (m *SecurityBaselineState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
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
            cast := val.(SecurityBaselineComplianceState)
            m.SetState(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
        cast := m.GetState().String()
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
// Sets the displayName property value. The display name of the security baseline
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SecurityBaselineState) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the securityBaselineTemplateId property value. The security baseline template id
// Parameters:
//  - value : Value to set for the securityBaselineTemplateId property.
func (m *SecurityBaselineState) SetSecurityBaselineTemplateId(value *string)() {
    m.securityBaselineTemplateId = value
}
// Sets the settingStates property value. The security baseline state for different settings for a device
// Parameters:
//  - value : Value to set for the settingStates property.
func (m *SecurityBaselineState) SetSettingStates(value []SecurityBaselineSettingState)() {
    m.settingStates = value
}
// Sets the state property value. Security baseline compliance state
// Parameters:
//  - value : Value to set for the state property.
func (m *SecurityBaselineState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
// Sets the userPrincipalName property value. User Principal Name
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *SecurityBaselineState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
