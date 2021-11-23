package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrganizationSettings 
type OrganizationSettings struct {
    Entity
    // Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
    itemInsights *InsightsSettings;
    // 
    peopleInsights *InsightsSettings;
    // Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
    profileCardProperties []ProfileCardProperty;
}
// NewOrganizationSettings instantiates a new organizationSettings and sets the default values.
func NewOrganizationSettings()(*OrganizationSettings) {
    m := &OrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetItemInsights gets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
func (m *OrganizationSettings) GetItemInsights()(*InsightsSettings) {
    if m == nil {
        return nil
    } else {
        return m.itemInsights
    }
}
// GetPeopleInsights gets the peopleInsights property value. 
func (m *OrganizationSettings) GetPeopleInsights()(*InsightsSettings) {
    if m == nil {
        return nil
    } else {
        return m.peopleInsights
    }
}
// GetProfileCardProperties gets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
func (m *OrganizationSettings) GetProfileCardProperties()([]ProfileCardProperty) {
    if m == nil {
        return nil
    } else {
        return m.profileCardProperties
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInsightsSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(*InsightsSettings))
        }
        return nil
    }
    res["peopleInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInsightsSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeopleInsights(val.(*InsightsSettings))
        }
        return nil
    }
    res["profileCardProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfileCardProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileCardProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProfileCardProperty))
            }
            m.SetProfileCardProperties(res)
        }
        return nil
    }
    return res
}
func (m *OrganizationSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteObjectValue("peopleInsights", m.GetPeopleInsights())
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
// SetItemInsights sets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. Get itemInsightsSettings through this navigation property.
func (m *OrganizationSettings) SetItemInsights(value *InsightsSettings)() {
    m.itemInsights = value
}
// SetPeopleInsights sets the peopleInsights property value. 
func (m *OrganizationSettings) SetPeopleInsights(value *InsightsSettings)() {
    m.peopleInsights = value
}
// SetProfileCardProperties sets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
func (m *OrganizationSettings) SetProfileCardProperties(value []ProfileCardProperty)() {
    m.profileCardProperties = value
}
