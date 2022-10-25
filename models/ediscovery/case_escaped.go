package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Case_escaped provides operations to manage the collection of accessReviewDecision entities.
type Case_escaped struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The user who closed the case.
    closedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns a list of case custodian objects for this case.  Nullable.
    custodians []Custodianable
    // The case description.
    description *string
    // The case name.
    displayName *string
    // The external case number for customer reference.
    externalId *string
    // The last user who modified the entity.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns a list of case legalHold objects for this case.  Nullable.
    legalHolds []LegalHoldable
    // Returns a list of case noncustodialDataSource objects for this case.  Nullable.
    noncustodialDataSources []NoncustodialDataSourceable
    // Returns a list of case operation objects for this case. Nullable.
    operations []CaseOperationable
    // Returns a list of reviewSet objects in the case. Read-only. Nullable.
    reviewSets []ReviewSetable
    // The settings property
    settings CaseSettingsable
    // Returns a list of sourceCollection objects associated with this case.
    sourceCollections []SourceCollectionable
    // The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
    status *CaseStatus
    // Returns a list of tag objects associated to this case.
    tags []Tagable
}
// NewCase_escaped instantiates a new case_escaped and sets the default values.
func NewCase_escaped()(*Case_escaped) {
    m := &Case_escaped{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.case";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCase_escapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCase_escapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCase_escaped(), nil
}
// GetClosedBy gets the closedBy property value. The user who closed the case.
func (m *Case_escaped) GetClosedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.closedBy
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustodians gets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *Case_escaped) GetCustodians()([]Custodianable) {
    return m.custodians
}
// GetDescription gets the description property value. The case description.
func (m *Case_escaped) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The case name.
func (m *Case_escaped) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. The external case number for customer reference.
func (m *Case_escaped) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Case_escaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["closedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetClosedBy)
    res["closedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetClosedDateTime)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["custodians"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustodianFromDiscriminatorValue , m.SetCustodians)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["legalHolds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLegalHoldFromDiscriminatorValue , m.SetLegalHolds)
    res["noncustodialDataSources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNoncustodialDataSourceFromDiscriminatorValue , m.SetNoncustodialDataSources)
    res["operations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCaseOperationFromDiscriminatorValue , m.SetOperations)
    res["reviewSets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateReviewSetFromDiscriminatorValue , m.SetReviewSets)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCaseSettingsFromDiscriminatorValue , m.SetSettings)
    res["sourceCollections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSourceCollectionFromDiscriminatorValue , m.SetSourceCollections)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCaseStatus , m.SetStatus)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTagFromDiscriminatorValue , m.SetTags)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the entity.
func (m *Case_escaped) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLegalHolds gets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *Case_escaped) GetLegalHolds()([]LegalHoldable) {
    return m.legalHolds
}
// GetNoncustodialDataSources gets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *Case_escaped) GetNoncustodialDataSources()([]NoncustodialDataSourceable) {
    return m.noncustodialDataSources
}
// GetOperations gets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *Case_escaped) GetOperations()([]CaseOperationable) {
    return m.operations
}
// GetReviewSets gets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *Case_escaped) GetReviewSets()([]ReviewSetable) {
    return m.reviewSets
}
// GetSettings gets the settings property value. The settings property
func (m *Case_escaped) GetSettings()(CaseSettingsable) {
    return m.settings
}
// GetSourceCollections gets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *Case_escaped) GetSourceCollections()([]SourceCollectionable) {
    return m.sourceCollections
}
// GetStatus gets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *Case_escaped) GetStatus()(*CaseStatus) {
    return m.status
}
// GetTags gets the tags property value. Returns a list of tag objects associated to this case.
func (m *Case_escaped) GetTags()([]Tagable) {
    return m.tags
}
// Serialize serializes information the current object
func (m *Case_escaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustodians())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLegalHolds())
        err = writer.WriteCollectionOfObjectValues("legalHolds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialDataSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNoncustodialDataSources())
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOperations())
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReviewSets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReviewSets())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSourceCollections())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTags())
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClosedBy sets the closedBy property value. The user who closed the case.
func (m *Case_escaped) SetClosedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.closedBy = value
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustodians sets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *Case_escaped) SetCustodians(value []Custodianable)() {
    m.custodians = value
}
// SetDescription sets the description property value. The case description.
func (m *Case_escaped) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The case name.
func (m *Case_escaped) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. The external case number for customer reference.
func (m *Case_escaped) SetExternalId(value *string)() {
    m.externalId = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the entity.
func (m *Case_escaped) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLegalHolds sets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *Case_escaped) SetLegalHolds(value []LegalHoldable)() {
    m.legalHolds = value
}
// SetNoncustodialDataSources sets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *Case_escaped) SetNoncustodialDataSources(value []NoncustodialDataSourceable)() {
    m.noncustodialDataSources = value
}
// SetOperations sets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *Case_escaped) SetOperations(value []CaseOperationable)() {
    m.operations = value
}
// SetReviewSets sets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *Case_escaped) SetReviewSets(value []ReviewSetable)() {
    m.reviewSets = value
}
// SetSettings sets the settings property value. The settings property
func (m *Case_escaped) SetSettings(value CaseSettingsable)() {
    m.settings = value
}
// SetSourceCollections sets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *Case_escaped) SetSourceCollections(value []SourceCollectionable)() {
    m.sourceCollections = value
}
// SetStatus sets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *Case_escaped) SetStatus(value *CaseStatus)() {
    m.status = value
}
// SetTags sets the tags property value. Returns a list of tag objects associated to this case.
func (m *Case_escaped) SetTags(value []Tagable)() {
    m.tags = value
}
