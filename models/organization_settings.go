package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationSettings 
type OrganizationSettings struct {
    Entity
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
// GetContactInsights gets the contactInsights property value. Contains the properties that are configured by an administrator as a tenant-level privacy control whether to identify duplicate contacts among a user's contacts list and suggest the user to merge those contacts to have a cleaner contacts list. List contactInsights returns the settings to display or return contact insights in an organization.
func (m *OrganizationSettings) GetContactInsights()(InsightsSettingsable) {
    val, err := m.GetBackingStore().Get("contactInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InsightsSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contactInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactInsights(val.(InsightsSettingsable))
        }
        return nil
    }
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
    return res
}
// GetItemInsights gets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
func (m *OrganizationSettings) GetItemInsights()(InsightsSettingsable) {
    val, err := m.GetBackingStore().Get("itemInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InsightsSettingsable)
    }
    return nil
}
// GetMicrosoftApplicationDataAccess gets the microsoftApplicationDataAccess property value. The microsoftApplicationDataAccess property
func (m *OrganizationSettings) GetMicrosoftApplicationDataAccess()(MicrosoftApplicationDataAccessSettingsable) {
    val, err := m.GetBackingStore().Get("microsoftApplicationDataAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MicrosoftApplicationDataAccessSettingsable)
    }
    return nil
}
// GetPeopleInsights gets the peopleInsights property value. Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
func (m *OrganizationSettings) GetPeopleInsights()(InsightsSettingsable) {
    val, err := m.GetBackingStore().Get("peopleInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InsightsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OrganizationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contactInsights", m.GetContactInsights())
        if err != nil {
            return err
        }
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
    return nil
}
// SetContactInsights sets the contactInsights property value. Contains the properties that are configured by an administrator as a tenant-level privacy control whether to identify duplicate contacts among a user's contacts list and suggest the user to merge those contacts to have a cleaner contacts list. List contactInsights returns the settings to display or return contact insights in an organization.
func (m *OrganizationSettings) SetContactInsights(value InsightsSettingsable)() {
    err := m.GetBackingStore().Set("contactInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetItemInsights sets the itemInsights property value. Contains the properties that are configured by an administrator for the visibility of Microsoft Graph-derived insights, between a user and other items in Microsoft 365, such as documents or sites. List itemInsights returns the settings to display or return item insights in an organization.
func (m *OrganizationSettings) SetItemInsights(value InsightsSettingsable)() {
    err := m.GetBackingStore().Set("itemInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftApplicationDataAccess sets the microsoftApplicationDataAccess property value. The microsoftApplicationDataAccess property
func (m *OrganizationSettings) SetMicrosoftApplicationDataAccess(value MicrosoftApplicationDataAccessSettingsable)() {
    err := m.GetBackingStore().Set("microsoftApplicationDataAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetPeopleInsights sets the peopleInsights property value. Contains the properties that are configured by an administrator for the visibility of a list of people relevant and working with a user in Microsoft 365. List peopleInsights returns the settings to display or return people insights in an organization.
func (m *OrganizationSettings) SetPeopleInsights(value InsightsSettingsable)() {
    err := m.GetBackingStore().Set("peopleInsights", value)
    if err != nil {
        panic(err)
    }
}
// OrganizationSettingsable 
type OrganizationSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactInsights()(InsightsSettingsable)
    GetItemInsights()(InsightsSettingsable)
    GetMicrosoftApplicationDataAccess()(MicrosoftApplicationDataAccessSettingsable)
    GetPeopleInsights()(InsightsSettingsable)
    SetContactInsights(value InsightsSettingsable)()
    SetItemInsights(value InsightsSettingsable)()
    SetMicrosoftApplicationDataAccess(value MicrosoftApplicationDataAccessSettingsable)()
    SetPeopleInsights(value InsightsSettingsable)()
}
