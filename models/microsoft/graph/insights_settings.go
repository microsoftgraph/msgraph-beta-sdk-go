package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InsightsSettings 
type InsightsSettings struct {
    Entity
    // The ID of an Azure AD group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
    disabledForGroup *string;
    // true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
    isEnabledInOrganization *bool;
}
// NewInsightsSettings instantiates a new insightsSettings and sets the default values.
func NewInsightsSettings()(*InsightsSettings) {
    m := &InsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisabledForGroup gets the disabledForGroup property value. The ID of an Azure AD group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
func (m *InsightsSettings) GetDisabledForGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledForGroup
    }
}
// GetIsEnabledInOrganization gets the isEnabledInOrganization property value. true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
func (m *InsightsSettings) GetIsEnabledInOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledInOrganization
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *InsightsSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InsightsSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDisabledForGroup sets the disabledForGroup property value. The ID of an Azure AD group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
func (m *InsightsSettings) SetDisabledForGroup(value *string)() {
    if m != nil {
        m.disabledForGroup = value
    }
}
// SetIsEnabledInOrganization sets the isEnabledInOrganization property value. true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
func (m *InsightsSettings) SetIsEnabledInOrganization(value *bool)() {
    if m != nil {
        m.isEnabledInOrganization = value
    }
}
