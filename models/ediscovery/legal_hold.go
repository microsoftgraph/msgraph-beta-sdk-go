package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LegalHold provides operations to manage the compliance singleton.
type LegalHold struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
    contentQuery *string
    // The user who created the legal hold.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date and time the legal hold was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The legal hold description.
    description *string
    // The display name of the legal hold.
    displayName *string
    // Lists any errors that happened while placing the hold.
    errors []string
    // Indicates whether the hold is enabled and actively holding content.
    isEnabled *bool
    // the user who last modified the legal hold.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date and time the legal hold was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Data source entity for SharePoint sites associated with the legal hold.
    siteSources []SiteSourceable
    // The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
    status *LegalHoldStatus
    // The unifiedGroupSources property
    unifiedGroupSources []UnifiedGroupSourceable
    // Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
    userSources []UserSourceable
}
// NewLegalHold instantiates a new legalHold and sets the default values.
func NewLegalHold()(*LegalHold) {
    m := &LegalHold{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.legalHold";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLegalHoldFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLegalHoldFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegalHold(), nil
}
// GetContentQuery gets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) GetContentQuery()(*string) {
    return m.contentQuery
}
// GetCreatedBy gets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The legal hold description.
func (m *LegalHold) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the legal hold.
func (m *LegalHold) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrors gets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) GetErrors()([]string) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LegalHold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentQuery"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentQuery)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["errors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetErrors)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["siteSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue , m.SetSiteSources)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLegalHoldStatus , m.SetStatus)
    res["unifiedGroupSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedGroupSourceFromDiscriminatorValue , m.SetUnifiedGroupSources)
    res["userSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue , m.SetUserSources)
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetLastModifiedBy gets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetSiteSources gets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) GetSiteSources()([]SiteSourceable) {
    return m.siteSources
}
// GetStatus gets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) GetStatus()(*LegalHoldStatus) {
    return m.status
}
// GetUnifiedGroupSources gets the unifiedGroupSources property value. The unifiedGroupSources property
func (m *LegalHold) GetUnifiedGroupSources()([]UnifiedGroupSourceable) {
    return m.unifiedGroupSources
}
// GetUserSources gets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) GetUserSources()([]UserSourceable) {
    return m.userSources
}
// Serialize serializes information the current object
func (m *LegalHold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
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
    if m.GetErrors() != nil {
        err = writer.WriteCollectionOfStringValues("errors", m.GetErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
    if m.GetSiteSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSiteSources())
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
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
    if m.GetUnifiedGroupSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUnifiedGroupSources())
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserSources())
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentQuery sets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// SetCreatedBy sets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The legal hold description.
func (m *LegalHold) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the legal hold.
func (m *LegalHold) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrors sets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) SetErrors(value []string)() {
    m.errors = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetSiteSources sets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) SetSiteSources(value []SiteSourceable)() {
    m.siteSources = value
}
// SetStatus sets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) SetStatus(value *LegalHoldStatus)() {
    m.status = value
}
// SetUnifiedGroupSources sets the unifiedGroupSources property value. The unifiedGroupSources property
func (m *LegalHold) SetUnifiedGroupSources(value []UnifiedGroupSourceable)() {
    m.unifiedGroupSources = value
}
// SetUserSources sets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) SetUserSources(value []UserSourceable)() {
    m.userSources = value
}
