package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteList a singleton entity which is used to specify IE mode site list metadata
type BrowserSiteList struct {
    Entity
    // The description of the site list.
    description *string
    // The name of the site list.
    displayName *string
    // The user who last modified the site list.
    lastModifiedBy IdentitySetable
    // The date and time when the site list was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user who published the site list.
    publishedBy IdentitySetable
    // The date and time when the site list was published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current revision of the site list.
    revision *string
    // A collection of shared cookies defined for the site list.
    sharedCookies []BrowserSharedCookieable
    // A collection of sites defined for the site list.
    sites []BrowserSiteable
    // The status property
    status *BrowserSiteListStatus
}
// NewBrowserSiteList instantiates a new browserSiteList and sets the default values.
func NewBrowserSiteList()(*BrowserSiteList) {
    m := &BrowserSiteList{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBrowserSiteListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSiteList(), nil
}
// GetDescription gets the description property value. The description of the site list.
func (m *BrowserSiteList) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the site list.
func (m *BrowserSiteList) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSiteList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["publishedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetPublishedBy)
    res["publishedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetPublishedDateTime)
    res["revision"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRevision)
    res["sharedCookies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBrowserSharedCookieFromDiscriminatorValue , m.SetSharedCookies)
    res["sites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBrowserSiteFromDiscriminatorValue , m.SetSites)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseBrowserSiteListStatus , m.SetStatus)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the site list.
func (m *BrowserSiteList) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the site list was last modified.
func (m *BrowserSiteList) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublishedBy gets the publishedBy property value. The user who published the site list.
func (m *BrowserSiteList) GetPublishedBy()(IdentitySetable) {
    return m.publishedBy
}
// GetPublishedDateTime gets the publishedDateTime property value. The date and time when the site list was published.
func (m *BrowserSiteList) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// GetRevision gets the revision property value. The current revision of the site list.
func (m *BrowserSiteList) GetRevision()(*string) {
    return m.revision
}
// GetSharedCookies gets the sharedCookies property value. A collection of shared cookies defined for the site list.
func (m *BrowserSiteList) GetSharedCookies()([]BrowserSharedCookieable) {
    return m.sharedCookies
}
// GetSites gets the sites property value. A collection of sites defined for the site list.
func (m *BrowserSiteList) GetSites()([]BrowserSiteable) {
    return m.sites
}
// GetStatus gets the status property value. The status property
func (m *BrowserSiteList) GetStatus()(*BrowserSiteListStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *BrowserSiteList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishedBy", m.GetPublishedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    if m.GetSharedCookies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSharedCookies())
        err = writer.WriteCollectionOfObjectValues("sharedCookies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSites() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSites())
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the site list.
func (m *BrowserSiteList) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the site list.
func (m *BrowserSiteList) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the site list.
func (m *BrowserSiteList) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the site list was last modified.
func (m *BrowserSiteList) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublishedBy sets the publishedBy property value. The user who published the site list.
func (m *BrowserSiteList) SetPublishedBy(value IdentitySetable)() {
    m.publishedBy = value
}
// SetPublishedDateTime sets the publishedDateTime property value. The date and time when the site list was published.
func (m *BrowserSiteList) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// SetRevision sets the revision property value. The current revision of the site list.
func (m *BrowserSiteList) SetRevision(value *string)() {
    m.revision = value
}
// SetSharedCookies sets the sharedCookies property value. A collection of shared cookies defined for the site list.
func (m *BrowserSiteList) SetSharedCookies(value []BrowserSharedCookieable)() {
    m.sharedCookies = value
}
// SetSites sets the sites property value. A collection of sites defined for the site list.
func (m *BrowserSiteList) SetSites(value []BrowserSiteable)() {
    m.sites = value
}
// SetStatus sets the status property value. The status property
func (m *BrowserSiteList) SetStatus(value *BrowserSiteListStatus)() {
    m.status = value
}
