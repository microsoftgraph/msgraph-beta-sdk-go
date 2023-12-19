package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LegalHold 
type LegalHold struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewLegalHold instantiates a new legalHold and sets the default values.
func NewLegalHold()(*LegalHold) {
    m := &LegalHold{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLegalHoldFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLegalHoldFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegalHold(), nil
}
// GetContentQuery gets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) GetContentQuery()(*string) {
    val, err := m.GetBackingStore().Get("contentQuery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The legal hold description.
func (m *LegalHold) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the legal hold.
func (m *LegalHold) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrors gets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) GetErrors()([]string) {
    val, err := m.GetBackingStore().Get("errors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LegalHold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["siteSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SiteSourceable)
                }
            }
            m.SetSiteSources(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLegalHold_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*LegalHold_status))
        }
        return nil
    }
    res["unifiedGroupSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedGroupSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedGroupSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedGroupSourceable)
                }
            }
            m.SetUnifiedGroupSources(res)
        }
        return nil
    }
    res["userSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserSourceable)
                }
            }
            m.SetUserSources(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSiteSources gets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) GetSiteSources()([]SiteSourceable) {
    val, err := m.GetBackingStore().Get("siteSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SiteSourceable)
    }
    return nil
}
// GetStatus gets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) GetStatus()(*LegalHold_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LegalHold_status)
    }
    return nil
}
// GetUnifiedGroupSources gets the unifiedGroupSources property value. The unifiedGroupSources property
func (m *LegalHold) GetUnifiedGroupSources()([]UnifiedGroupSourceable) {
    val, err := m.GetBackingStore().Get("unifiedGroupSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedGroupSourceable)
    }
    return nil
}
// GetUserSources gets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) GetUserSources()([]UserSourceable) {
    val, err := m.GetBackingStore().Get("userSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserSourceable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnifiedGroupSources()))
        for i, v := range m.GetUnifiedGroupSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentQuery sets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) SetContentQuery(value *string)() {
    err := m.GetBackingStore().Set("contentQuery", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The legal hold description.
func (m *LegalHold) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the legal hold.
func (m *LegalHold) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetErrors sets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) SetErrors(value []string)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteSources sets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) SetSiteSources(value []SiteSourceable)() {
    err := m.GetBackingStore().Set("siteSources", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) SetStatus(value *LegalHold_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUnifiedGroupSources sets the unifiedGroupSources property value. The unifiedGroupSources property
func (m *LegalHold) SetUnifiedGroupSources(value []UnifiedGroupSourceable)() {
    err := m.GetBackingStore().Set("unifiedGroupSources", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSources sets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) SetUserSources(value []UserSourceable)() {
    err := m.GetBackingStore().Set("userSources", value)
    if err != nil {
        panic(err)
    }
}
// LegalHoldable 
type LegalHoldable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentQuery()(*string)
    GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetErrors()([]string)
    GetIsEnabled()(*bool)
    GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSiteSources()([]SiteSourceable)
    GetStatus()(*LegalHold_status)
    GetUnifiedGroupSources()([]UnifiedGroupSourceable)
    GetUserSources()([]UserSourceable)
    SetContentQuery(value *string)()
    SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetErrors(value []string)()
    SetIsEnabled(value *bool)()
    SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSiteSources(value []SiteSourceable)()
    SetStatus(value *LegalHold_status)()
    SetUnifiedGroupSources(value []UnifiedGroupSourceable)()
    SetUserSources(value []UserSourceable)()
}
