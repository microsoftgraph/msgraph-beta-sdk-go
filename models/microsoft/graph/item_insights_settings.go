package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemInsightsSettings struct {
    Entity
    // The ID of an Azure AD group, of which the members' item insights are disabled. Default is empty. Optional.
    disabledForGroup *string;
    // true if organization item insights are enabled; false if organization item insights are disabled for all users without exceptions. Default is true. Optional.
    isEnabledInOrganization *bool;
}
// Instantiates a new itemInsightsSettings and sets the default values.
func NewItemInsightsSettings()(*ItemInsightsSettings) {
    m := &ItemInsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the disabledForGroup property value. The ID of an Azure AD group, of which the members' item insights are disabled. Default is empty. Optional.
func (m *ItemInsightsSettings) GetDisabledForGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledForGroup
    }
}
// Gets the isEnabledInOrganization property value. true if organization item insights are enabled; false if organization item insights are disabled for all users without exceptions. Default is true. Optional.
func (m *ItemInsightsSettings) GetIsEnabledInOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledInOrganization
    }
}
// The deserialization information for the current model
func (m *ItemInsightsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["disabledForGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledForGroup(val)
        }
        return nil
    }
    res["isEnabledInOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledInOrganization(val)
        }
        return nil
    }
    return res
}
func (m *ItemInsightsSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemInsightsSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("disabledForGroup", m.GetDisabledForGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledInOrganization", m.GetIsEnabledInOrganization())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the disabledForGroup property value. The ID of an Azure AD group, of which the members' item insights are disabled. Default is empty. Optional.
// Parameters:
//  - value : Value to set for the disabledForGroup property.
func (m *ItemInsightsSettings) SetDisabledForGroup(value *string)() {
    m.disabledForGroup = value
}
// Sets the isEnabledInOrganization property value. true if organization item insights are enabled; false if organization item insights are disabled for all users without exceptions. Default is true. Optional.
// Parameters:
//  - value : Value to set for the isEnabledInOrganization property.
func (m *ItemInsightsSettings) SetIsEnabledInOrganization(value *bool)() {
    m.isEnabledInOrganization = value
}
