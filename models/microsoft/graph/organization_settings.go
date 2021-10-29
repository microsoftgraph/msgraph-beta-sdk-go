package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OrganizationSettings struct {
    Entity
    // Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
    itemInsights *ItemInsightsSettings;
    // Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
    profileCardProperties []ProfileCardProperty;
}
// Instantiates a new organizationSettings and sets the default values.
func NewOrganizationSettings()(*OrganizationSettings) {
    m := &OrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
func (m *OrganizationSettings) GetItemInsights()(*ItemInsightsSettings) {
    if m == nil {
        return nil
    } else {
        return m.itemInsights
    }
}
// Gets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
func (m *OrganizationSettings) GetProfileCardProperties()([]ProfileCardProperty) {
    if m == nil {
        return nil
    } else {
        return m.profileCardProperties
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
// Parameters:
//  - value : Value to set for the itemInsights property.
func (m *OrganizationSettings) SetItemInsights(value *ItemInsightsSettings)() {
    m.itemInsights = value
}
// Sets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
// Parameters:
//  - value : Value to set for the profileCardProperties property.
func (m *OrganizationSettings) SetProfileCardProperties(value []ProfileCardProperty)() {
    m.profileCardProperties = value
}
