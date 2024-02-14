package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type CaseEscaped struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCaseEscaped instantiates a new CaseEscaped and sets the default values.
func NewCaseEscaped()(*CaseEscaped) {
    m := &CaseEscaped{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCaseEscapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCaseEscapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCaseEscaped(), nil
}
// GetClosedBy gets the closedBy property value. The user who closed the case.
// returns a IdentitySetable when successful
func (m *CaseEscaped) GetClosedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("closedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *CaseEscaped) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("closedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *CaseEscaped) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustodians gets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
// returns a []Custodianable when successful
func (m *CaseEscaped) GetCustodians()([]Custodianable) {
    val, err := m.GetBackingStore().Get("custodians")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Custodianable)
    }
    return nil
}
// GetDescription gets the description property value. The case description.
// returns a *string when successful
func (m *CaseEscaped) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The case name.
// returns a *string when successful
func (m *CaseEscaped) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The external case number for customer reference.
// returns a *string when successful
func (m *CaseEscaped) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CaseEscaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["closedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["closedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
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
    res["custodians"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Custodianable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Custodianable)
                }
            }
            m.SetCustodians(res)
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
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
    res["legalHolds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLegalHoldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LegalHoldable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LegalHoldable)
                }
            }
            m.SetLegalHolds(res)
        }
        return nil
    }
    res["noncustodialDataSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NoncustodialDataSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NoncustodialDataSourceable)
                }
            }
            m.SetNoncustodialDataSources(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCaseOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CaseOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CaseOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["reviewSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReviewSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ReviewSetable)
                }
            }
            m.SetReviewSets(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCaseSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(CaseSettingsable))
        }
        return nil
    }
    res["sourceCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SourceCollectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SourceCollectionable)
                }
            }
            m.SetSourceCollections(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCaseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CaseStatus))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Tagable)
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the entity.
// returns a IdentitySetable when successful
func (m *CaseEscaped) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *CaseEscaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLegalHolds gets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
// returns a []LegalHoldable when successful
func (m *CaseEscaped) GetLegalHolds()([]LegalHoldable) {
    val, err := m.GetBackingStore().Get("legalHolds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LegalHoldable)
    }
    return nil
}
// GetNoncustodialDataSources gets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// returns a []NoncustodialDataSourceable when successful
func (m *CaseEscaped) GetNoncustodialDataSources()([]NoncustodialDataSourceable) {
    val, err := m.GetBackingStore().Get("noncustodialDataSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NoncustodialDataSourceable)
    }
    return nil
}
// GetOperations gets the operations property value. Returns a list of case operation objects for this case. Nullable.
// returns a []CaseOperationable when successful
func (m *CaseEscaped) GetOperations()([]CaseOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CaseOperationable)
    }
    return nil
}
// GetReviewSets gets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
// returns a []ReviewSetable when successful
func (m *CaseEscaped) GetReviewSets()([]ReviewSetable) {
    val, err := m.GetBackingStore().Get("reviewSets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ReviewSetable)
    }
    return nil
}
// GetSettings gets the settings property value. The settings property
// returns a CaseSettingsable when successful
func (m *CaseEscaped) GetSettings()(CaseSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CaseSettingsable)
    }
    return nil
}
// GetSourceCollections gets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
// returns a []SourceCollectionable when successful
func (m *CaseEscaped) GetSourceCollections()([]SourceCollectionable) {
    val, err := m.GetBackingStore().Get("sourceCollections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SourceCollectionable)
    }
    return nil
}
// GetStatus gets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
// returns a *CaseStatus when successful
func (m *CaseEscaped) GetStatus()(*CaseStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CaseStatus)
    }
    return nil
}
// GetTags gets the tags property value. Returns a list of tag objects associated to this case.
// returns a []Tagable when successful
func (m *CaseEscaped) GetTags()([]Tagable) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Tagable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CaseEscaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("closedBy", m.GetClosedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
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
    if m.GetCustodians() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("custodians", cast)
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
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
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
    if m.GetLegalHolds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLegalHolds()))
        for i, v := range m.GetLegalHolds() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("legalHolds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialDataSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialDataSources()))
        for i, v := range m.GetNoncustodialDataSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReviewSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewSets()))
        for i, v := range m.GetReviewSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("reviewSets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSourceCollections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSourceCollections()))
        for i, v := range m.GetSourceCollections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sourceCollections", cast)
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
    if m.GetTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClosedBy sets the closedBy property value. The user who closed the case.
func (m *CaseEscaped) SetClosedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("closedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CaseEscaped) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("closedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CaseEscaped) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustodians sets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *CaseEscaped) SetCustodians(value []Custodianable)() {
    err := m.GetBackingStore().Set("custodians", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The case description.
func (m *CaseEscaped) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The case name.
func (m *CaseEscaped) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The external case number for customer reference.
func (m *CaseEscaped) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the entity.
func (m *CaseEscaped) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CaseEscaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLegalHolds sets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *CaseEscaped) SetLegalHolds(value []LegalHoldable)() {
    err := m.GetBackingStore().Set("legalHolds", value)
    if err != nil {
        panic(err)
    }
}
// SetNoncustodialDataSources sets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *CaseEscaped) SetNoncustodialDataSources(value []NoncustodialDataSourceable)() {
    err := m.GetBackingStore().Set("noncustodialDataSources", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *CaseEscaped) SetOperations(value []CaseOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewSets sets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *CaseEscaped) SetReviewSets(value []ReviewSetable)() {
    err := m.GetBackingStore().Set("reviewSets", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings property
func (m *CaseEscaped) SetSettings(value CaseSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceCollections sets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *CaseEscaped) SetSourceCollections(value []SourceCollectionable)() {
    err := m.GetBackingStore().Set("sourceCollections", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *CaseEscaped) SetStatus(value *CaseStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. Returns a list of tag objects associated to this case.
func (m *CaseEscaped) SetTags(value []Tagable)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
type CaseEscapedable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]Custodianable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalHolds()([]LegalHoldable)
    GetNoncustodialDataSources()([]NoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]ReviewSetable)
    GetSettings()(CaseSettingsable)
    GetSourceCollections()([]SourceCollectionable)
    GetStatus()(*CaseStatus)
    GetTags()([]Tagable)
    SetClosedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []Custodianable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalHolds(value []LegalHoldable)()
    SetNoncustodialDataSources(value []NoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []ReviewSetable)()
    SetSettings(value CaseSettingsable)()
    SetSourceCollections(value []SourceCollectionable)()
    SetStatus(value *CaseStatus)()
    SetTags(value []Tagable)()
}
