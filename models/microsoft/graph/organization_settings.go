package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrganizationSettings provides operations to manage the collection of organization entities.
type OrganizationSettings struct {
    Entity
    // Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
    itemInsights InsightsSettingsable;
    // Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
    peopleInsights InsightsSettingsable;
    // Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
    profileCardProperties []ProfileCardPropertyable;
}
// NewOrganizationSettings instantiates a new organizationSettings and sets the default values.
func NewOrganizationSettings()(*OrganizationSettings) {
    m := &OrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOrganizationSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["peopleInsights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeopleInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["profileCardProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileCardPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileCardPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ProfileCardPropertyable)
            }
            m.SetProfileCardProperties(res)
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
func (m *OrganizationSettings) GetItemInsights()(InsightsSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.itemInsights
    }
}
// GetPeopleInsights gets the peopleInsights property value. Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
func (m *OrganizationSettings) GetPeopleInsights()(InsightsSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.peopleInsights
    }
}
// GetProfileCardProperties gets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
func (m *OrganizationSettings) GetProfileCardProperties()([]ProfileCardPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.profileCardProperties
    }
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
    if m.GetProfileCardProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProfileCardProperties()))
        for i, v := range m.GetProfileCardProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("profileCardProperties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemInsights sets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
func (m *OrganizationSettings) SetItemInsights(value InsightsSettingsable)() {
    if m != nil {
        m.itemInsights = value
    }
}
// SetPeopleInsights sets the peopleInsights property value. Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
func (m *OrganizationSettings) SetPeopleInsights(value InsightsSettingsable)() {
    if m != nil {
        m.peopleInsights = value
    }
}
// SetProfileCardProperties sets the profileCardProperties property value. Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
func (m *OrganizationSettings) SetProfileCardProperties(value []ProfileCardPropertyable)() {
    if m != nil {
        m.profileCardProperties = value
    }
}
