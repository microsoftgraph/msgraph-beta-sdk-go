package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OrganizationSettings struct {
    Entity
    itemInsights *ItemInsightsSettings;
    profileCardProperties []ProfileCardProperty;
}
func NewOrganizationSettings()(*OrganizationSettings) {
    m := &OrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OrganizationSettings) GetItemInsights()(*ItemInsightsSettings) {
    if m == nil {
        return nil
    } else {
        return m.itemInsights
    }
}
func (m *OrganizationSettings) GetProfileCardProperties()([]ProfileCardProperty) {
    if m == nil {
        return nil
    } else {
        return m.profileCardProperties
    }
}
func (m *OrganizationSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemInsightsSettings() })
        if err != nil {
            return err
        }
        m.SetItemInsights(val.(*ItemInsightsSettings))
        return nil
    }
    res["profileCardProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfileCardProperty() })
        if err != nil {
            return err
        }
        res := make([]ProfileCardProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProfileCardProperty))
        }
        m.SetProfileCardProperties(res)
        return nil
    }
    return res
}
func (m *OrganizationSettings) IsNil()(bool) {
    return m == nil
}
func (m *OrganizationSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProfileCardProperties()))
        for i, v := range m.GetProfileCardProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("profileCardProperties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OrganizationSettings) SetItemInsights(value *ItemInsightsSettings)() {
    m.itemInsights = value
}
func (m *OrganizationSettings) SetProfileCardProperties(value []ProfileCardProperty)() {
    m.profileCardProperties = value
}
