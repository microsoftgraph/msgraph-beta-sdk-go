package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// Case_escaped 
type Case_escaped struct {
    Entity
    // The user who closed the case.
    closedBy *IdentitySet;
    // The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Returns a list of case custodian objects for this case.  Nullable.
    custodians []Custodian;
    // The case description.
    description *string;
    // The case name.
    displayName *string;
    // The external case number for customer reference.
    externalId *string;
    // The last user who modified the entity.
    lastModifiedBy *IdentitySet;
    // The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Returns a list of case legalHold objects for this case.  Nullable.
    legalHolds []LegalHold;
    // Returns a list of case noncustodialDataSource objects for this case.  Nullable.
    noncustodialDataSources []NoncustodialDataSource;
    // Returns a list of case operation objects for this case. Nullable.
    operations []CaseOperation;
    // Returns a list of reviewSet objects in the case. Read-only. Nullable.
    reviewSets []ReviewSet;
    // 
    settings *CaseSettings;
    // Returns a list of sourceCollection objects associated with this case.
    sourceCollections []SourceCollection;
    // The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
    status *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus;
    // Returns a list of tag objects associated to this case.
    tags []Tag;
}
// NewCase_escaped instantiates a new case_escaped and sets the default values.
func NewCase_escaped()(*Case_escaped) {
    m := &Case_escaped{
        Entity: *NewEntity(),
    }
    return m
}
// GetClosedBy gets the closedBy property value. The user who closed the case.
func (m *Case_escaped) GetClosedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.closedBy
    }
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustodians gets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *Case_escaped) GetCustodians()([]Custodian) {
    if m == nil {
        return nil
    } else {
        return m.custodians
    }
}
// GetDescription gets the description property value. The case description.
func (m *Case_escaped) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The case name.
func (m *Case_escaped) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. The external case number for customer reference.
func (m *Case_escaped) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the entity.
func (m *Case_escaped) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLegalHolds gets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *Case_escaped) GetLegalHolds()([]LegalHold) {
    if m == nil {
        return nil
    } else {
        return m.legalHolds
    }
}
// GetNoncustodialDataSources gets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *Case_escaped) GetNoncustodialDataSources()([]NoncustodialDataSource) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialDataSources
    }
}
// GetOperations gets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *Case_escaped) GetOperations()([]CaseOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetReviewSets gets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *Case_escaped) GetReviewSets()([]ReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.reviewSets
    }
}
// GetSettings gets the settings property value. 
func (m *Case_escaped) GetSettings()(*CaseSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetSourceCollections gets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *Case_escaped) GetSourceCollections()([]SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollections
    }
}
// GetStatus gets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *Case_escaped) GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTags gets the tags property value. Returns a list of tag objects associated to this case.
func (m *Case_escaped) GetTags()([]Tag) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Case_escaped) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["closedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
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
    res["custodians"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustodian() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Custodian, len(val))
            for i, v := range val {
                res[i] = *(v.(*Custodian))
            }
            m.SetCustodians(res)
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
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
    res["legalHolds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLegalHold() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LegalHold, len(val))
            for i, v := range val {
                res[i] = *(v.(*LegalHold))
            }
            m.SetLegalHolds(res)
        }
        return nil
    }
    res["noncustodialDataSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialDataSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NoncustodialDataSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*NoncustodialDataSource))
            }
            m.SetNoncustodialDataSources(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CaseOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*CaseOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["reviewSets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReviewSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReviewSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*ReviewSet))
            }
            m.SetReviewSets(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*CaseSettings))
        }
        return nil
    }
    res["sourceCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SourceCollection, len(val))
            for i, v := range val {
                res[i] = *(v.(*SourceCollection))
            }
            m.SetSourceCollections(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseCaseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus))
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tag, len(val))
            for i, v := range val {
                res[i] = *(v.(*Tag))
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
func (m *Case_escaped) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Case_escaped) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLegalHolds()))
        for i, v := range m.GetLegalHolds() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("legalHolds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialDataSources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNoncustodialDataSources()))
        for i, v := range m.GetNoncustodialDataSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReviewSets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewSets()))
        for i, v := range m.GetReviewSets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSourceCollections()))
        for i, v := range m.GetSourceCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClosedBy sets the closedBy property value. The user who closed the case.
func (m *Case_escaped) SetClosedBy(value *IdentitySet)() {
    if m != nil {
        m.closedBy = value
    }
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.closedDateTime = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustodians sets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *Case_escaped) SetCustodians(value []Custodian)() {
    if m != nil {
        m.custodians = value
    }
}
// SetDescription sets the description property value. The case description.
func (m *Case_escaped) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The case name.
func (m *Case_escaped) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. The external case number for customer reference.
func (m *Case_escaped) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the entity.
func (m *Case_escaped) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLegalHolds sets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *Case_escaped) SetLegalHolds(value []LegalHold)() {
    if m != nil {
        m.legalHolds = value
    }
}
// SetNoncustodialDataSources sets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *Case_escaped) SetNoncustodialDataSources(value []NoncustodialDataSource)() {
    if m != nil {
        m.noncustodialDataSources = value
    }
}
// SetOperations sets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *Case_escaped) SetOperations(value []CaseOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetReviewSets sets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *Case_escaped) SetReviewSets(value []ReviewSet)() {
    if m != nil {
        m.reviewSets = value
    }
}
// SetSettings sets the settings property value. 
func (m *Case_escaped) SetSettings(value *CaseSettings)() {
    if m != nil {
        m.settings = value
    }
}
// SetSourceCollections sets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *Case_escaped) SetSourceCollections(value []SourceCollection)() {
    if m != nil {
        m.sourceCollections = value
    }
}
// SetStatus sets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *Case_escaped) SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTags sets the tags property value. Returns a list of tag objects associated to this case.
func (m *Case_escaped) SetTags(value []Tag)() {
    if m != nil {
        m.tags = value
    }
}
