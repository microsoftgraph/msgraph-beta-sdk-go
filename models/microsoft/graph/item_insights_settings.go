package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemInsightsSettings struct {
    Entity
    disabledForGroup *string;
    isEnabledInOrganization *bool;
}
func NewItemInsightsSettings()(*ItemInsightsSettings) {
    m := &ItemInsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ItemInsightsSettings) GetDisabledForGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledForGroup
    }
}
func (m *ItemInsightsSettings) GetIsEnabledInOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledInOrganization
    }
}
func (m *ItemInsightsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["disabledForGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisabledForGroup(val)
        return nil
    }
    res["isEnabledInOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabledInOrganization(val)
        return nil
    }
    return res
}
func (m *ItemInsightsSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *ItemInsightsSettings) SetDisabledForGroup(value *string)() {
    m.disabledForGroup = value
}
func (m *ItemInsightsSettings) SetIsEnabledInOrganization(value *bool)() {
    m.isEnabledInOrganization = value
}
