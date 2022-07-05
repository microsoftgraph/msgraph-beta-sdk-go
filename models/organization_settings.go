package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationSettings provides operations to manage the collection of accessReview entities.
type OrganizationSettings struct {
    Entity
    // Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
    itemInsights InsightsSettingsable
    // The microsoftApplicationDataAccess property
    microsoftApplicationDataAccess MicrosoftApplicationDataAccessSettingsable
    // Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
    peopleInsights InsightsSettingsable
    // Contains a collection of the properties an administrator has defined as visible on the Microsoft 365 profile card. Get organization settings returns the properties configured for profile cards for the organization.
    profileCardProperties []ProfileCardPropertyable
}
// NewOrganizationSettings instantiates a new organizationSettings and sets the default values.
func NewOrganizationSettings()(*OrganizationSettings) {
    m := &OrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["microsoftApplicationDataAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftApplicationDataAccess(val.(MicrosoftApplicationDataAccessSettingsable))
        }
        return nil
    }
    res["peopleInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeopleInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["profileCardProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetMicrosoftApplicationDataAccess gets the microsoftApplicationDataAccess property value. The microsoftApplicationDataAccess property
func (m *OrganizationSettings) GetMicrosoftApplicationDataAccess()(MicrosoftApplicationDataAccessSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.microsoftApplicationDataAccess
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
// Serialize serializes information the current object
func (m *OrganizationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("microsoftApplicationDataAccess", m.GetMicrosoftApplicationDataAccess())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProfileCardProperties()))
        for i, v := range m.GetProfileCardProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetMicrosoftApplicationDataAccess sets the microsoftApplicationDataAccess property value. The microsoftApplicationDataAccess property
func (m *OrganizationSettings) SetMicrosoftApplicationDataAccess(value MicrosoftApplicationDataAccessSettingsable)() {
    if m != nil {
        m.microsoftApplicationDataAccess = value
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
