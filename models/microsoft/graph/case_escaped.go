package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

type Case_escaped struct {
    Entity
    closedBy *IdentitySet;
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    custodians []Custodian;
    description *string;
    displayName *string;
    externalId *string;
    lastModifiedBy *IdentitySet;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    legalHolds []LegalHold;
    noncustodialDataSources []NoncustodialDataSource;
    operations []CaseOperation;
    reviewSets []ReviewSet;
    settings *Settings;
    sourceCollections []SourceCollection;
    status *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus;
    tags []Tag;
}
func NewCase_escaped()(*Case_escaped) {
    m := &Case_escaped{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Case_escaped) GetClosedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.closedBy
    }
}
func (m *Case_escaped) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
func (m *Case_escaped) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Case_escaped) GetCustodians()([]Custodian) {
    if m == nil {
        return nil
    } else {
        return m.custodians
    }
}
func (m *Case_escaped) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *Case_escaped) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Case_escaped) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *Case_escaped) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
func (m *Case_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *Case_escaped) GetLegalHolds()([]LegalHold) {
    if m == nil {
        return nil
    } else {
        return m.legalHolds
    }
}
func (m *Case_escaped) GetNoncustodialDataSources()([]NoncustodialDataSource) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialDataSources
    }
}
func (m *Case_escaped) GetOperations()([]CaseOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Case_escaped) GetReviewSets()([]ReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.reviewSets
    }
}
func (m *Case_escaped) GetSettings()(*Settings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *Case_escaped) GetSourceCollections()([]SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollections
    }
}
func (m *Case_escaped) GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *Case_escaped) GetTags()([]Tag) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
func (m *Case_escaped) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["closedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetClosedBy(val.(*IdentitySet))
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetClosedDateTime(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["custodians"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustodian() })
        if err != nil {
            return err
        }
        res := make([]Custodian, len(val))
        for i, v := range val {
            res[i] = *(v.(*Custodian))
        }
        m.SetCustodians(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*IdentitySet))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["legalHolds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLegalHold() })
        if err != nil {
            return err
        }
        res := make([]LegalHold, len(val))
        for i, v := range val {
            res[i] = *(v.(*LegalHold))
        }
        m.SetLegalHolds(res)
        return nil
    }
    res["noncustodialDataSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialDataSource() })
        if err != nil {
            return err
        }
        res := make([]NoncustodialDataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*NoncustodialDataSource))
        }
        m.SetNoncustodialDataSources(res)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseOperation() })
        if err != nil {
            return err
        }
        res := make([]CaseOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*CaseOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["reviewSets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReviewSet() })
        if err != nil {
            return err
        }
        res := make([]ReviewSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*ReviewSet))
        }
        m.SetReviewSets(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*Settings))
        return nil
    }
    res["sourceCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        res := make([]SourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*SourceCollection))
        }
        m.SetSourceCollections(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseCaseStatus)
        if err != nil {
            return err
        }
        cast := val.(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        res := make([]Tag, len(val))
        for i, v := range val {
            res[i] = *(v.(*Tag))
        }
        m.SetTags(res)
        return nil
    }
    return res
}
func (m *Case_escaped) IsNil()(bool) {
    return m == nil
}
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
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
func (m *Case_escaped) SetClosedBy(value *IdentitySet)() {
    m.closedBy = value
}
func (m *Case_escaped) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
func (m *Case_escaped) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Case_escaped) SetCustodians(value []Custodian)() {
    m.custodians = value
}
func (m *Case_escaped) SetDescription(value *string)() {
    m.description = value
}
func (m *Case_escaped) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Case_escaped) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *Case_escaped) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
func (m *Case_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Case_escaped) SetLegalHolds(value []LegalHold)() {
    m.legalHolds = value
}
func (m *Case_escaped) SetNoncustodialDataSources(value []NoncustodialDataSource)() {
    m.noncustodialDataSources = value
}
func (m *Case_escaped) SetOperations(value []CaseOperation)() {
    m.operations = value
}
func (m *Case_escaped) SetReviewSets(value []ReviewSet)() {
    m.reviewSets = value
}
func (m *Case_escaped) SetSettings(value *Settings)() {
    m.settings = value
}
func (m *Case_escaped) SetSourceCollections(value []SourceCollection)() {
    m.sourceCollections = value
}
func (m *Case_escaped) SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus)() {
    m.status = value
}
func (m *Case_escaped) SetTags(value []Tag)() {
    m.tags = value
}
