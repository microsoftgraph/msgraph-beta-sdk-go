package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// LegalHold 
type LegalHold struct {
    Entity
    // KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
    contentQuery *string;
    // The user who created the legal hold.
    createdBy *IdentitySet;
    // The date and time the legal hold was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The legal hold description.
    description *string;
    // The display name of the legal hold.
    displayName *string;
    // Lists any errors that happened while placing the hold.
    errors []string;
    // Indicates whether the hold is enabled and actively holding content.
    isEnabled *bool;
    // the user who last modified the legal hold.
    lastModifiedBy *IdentitySet;
    // The date and time the legal hold was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Data source entity for SharePoint sites associated with the legal hold.
    siteSources []SiteSource;
    // The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
    status *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus;
    // 
    unifiedGroupSources []UnifiedGroupSource;
    // Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
    userSources []UserSource;
}
// NewLegalHold instantiates a new legalHold and sets the default values.
func NewLegalHold()(*LegalHold) {
    m := &LegalHold{
        Entity: *NewEntity(),
    }
    return m
}
// GetContentQuery gets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) GetContentQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentQuery
    }
}
// GetCreatedBy gets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The legal hold description.
func (m *LegalHold) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the legal hold.
func (m *LegalHold) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetErrors gets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) GetErrors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetSiteSources gets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) GetSiteSources()([]SiteSource) {
    if m == nil {
        return nil
    } else {
        return m.siteSources
    }
}
// GetStatus gets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetUnifiedGroupSources gets the unifiedGroupSources property value. 
func (m *LegalHold) GetUnifiedGroupSources()([]UnifiedGroupSource) {
    if m == nil {
        return nil
    } else {
        return m.unifiedGroupSources
    }
}
// GetUserSources gets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) GetUserSources()([]UserSource) {
    if m == nil {
        return nil
    } else {
        return m.userSources
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LegalHold) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentQuery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["siteSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSiteSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*SiteSource))
            }
            m.SetSiteSources(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseLegalHoldStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus))
        }
        return nil
    }
    res["unifiedGroupSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedGroupSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedGroupSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedGroupSource))
            }
            m.SetUnifiedGroupSources(res)
        }
        return nil
    }
    res["userSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSource))
            }
            m.SetUserSources(res)
        }
        return nil
    }
    return res
}
func (m *LegalHold) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LegalHold) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUnifiedGroupSources()))
        for i, v := range m.GetUnifiedGroupSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m != nil {
        m.contentQuery = value
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The legal hold description.
func (m *LegalHold) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the legal hold.
func (m *LegalHold) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetErrors sets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) SetErrors(value []string)() {
    if m != nil {
        m.errors = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetSiteSources sets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) SetSiteSources(value []SiteSource)() {
    if m != nil {
        m.siteSources = value
    }
}
// SetStatus sets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.LegalHoldStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetUnifiedGroupSources sets the unifiedGroupSources property value. 
func (m *LegalHold) SetUnifiedGroupSources(value []UnifiedGroupSource)() {
    if m != nil {
        m.unifiedGroupSources = value
    }
}
// SetUserSources sets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) SetUserSources(value []UserSource)() {
    if m != nil {
        m.userSources = value
    }
}
