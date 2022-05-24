package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EdiscoveryCase provides operations to manage the security singleton.
type EdiscoveryCase struct {
    Case_escaped
    // The closedBy property
    closedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The closedDateTime property
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The custodians property
    custodians []EdiscoveryCustodianable
    // The externalId property
    externalId *string
    // The legalHolds property
    legalHolds []EdiscoveryHoldPolicyable
    // The noncustodialDataSources property
    noncustodialDataSources []EdiscoveryNoncustodialDataSourceable
    // The operations property
    operations []CaseOperationable
    // The reviewSets property
    reviewSets []EdiscoveryReviewSetable
    // The searches property
    searches []EdiscoverySearchable
    // The settings property
    settings EdiscoveryCaseSettingsable
    // The tags property
    tags []EdiscoveryReviewTagable
}
// NewEdiscoveryCase instantiates a new ediscoveryCase and sets the default values.
func NewEdiscoveryCase()(*EdiscoveryCase) {
    m := &EdiscoveryCase{
        Case_escaped: *NewCase_escaped(),
    }
    return m
}
// CreateEdiscoveryCaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCase(), nil
}
// GetClosedBy gets the closedBy property value. The closedBy property
func (m *EdiscoveryCase) GetClosedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.closedBy
    }
}
// GetClosedDateTime gets the closedDateTime property value. The closedDateTime property
func (m *EdiscoveryCase) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
// GetCustodians gets the custodians property value. The custodians property
func (m *EdiscoveryCase) GetCustodians()([]EdiscoveryCustodianable) {
    if m == nil {
        return nil
    } else {
        return m.custodians
    }
}
// GetExternalId gets the externalId property value. The externalId property
func (m *EdiscoveryCase) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Case_escaped.GetFieldDeserializers()
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
    res["custodians"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryCustodianable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryCustodianable)
            }
            m.SetCustodians(res)
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
    res["legalHolds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryHoldPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryHoldPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryHoldPolicyable)
            }
            m.SetLegalHolds(res)
        }
        return nil
    }
    res["noncustodialDataSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryNoncustodialDataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryNoncustodialDataSourceable)
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
                res[i] = v.(CaseOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["reviewSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewSetable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryReviewSetable)
            }
            m.SetReviewSets(res)
        }
        return nil
    }
    res["searches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoverySearchable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoverySearchable)
            }
            m.SetSearches(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryCaseSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(EdiscoveryCaseSettingsable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryReviewTagable)
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetLegalHolds gets the legalHolds property value. The legalHolds property
func (m *EdiscoveryCase) GetLegalHolds()([]EdiscoveryHoldPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.legalHolds
    }
}
// GetNoncustodialDataSources gets the noncustodialDataSources property value. The noncustodialDataSources property
func (m *EdiscoveryCase) GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialDataSources
    }
}
// GetOperations gets the operations property value. The operations property
func (m *EdiscoveryCase) GetOperations()([]CaseOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetReviewSets gets the reviewSets property value. The reviewSets property
func (m *EdiscoveryCase) GetReviewSets()([]EdiscoveryReviewSetable) {
    if m == nil {
        return nil
    } else {
        return m.reviewSets
    }
}
// GetSearches gets the searches property value. The searches property
func (m *EdiscoveryCase) GetSearches()([]EdiscoverySearchable) {
    if m == nil {
        return nil
    } else {
        return m.searches
    }
}
// GetSettings gets the settings property value. The settings property
func (m *EdiscoveryCase) GetSettings()(EdiscoveryCaseSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetTags gets the tags property value. The tags property
func (m *EdiscoveryCase) GetTags()([]EdiscoveryReviewTagable) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryCase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Case_escaped.Serialize(writer)
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
    if m.GetCustodians() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("custodians", cast)
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
    if m.GetLegalHolds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLegalHolds()))
        for i, v := range m.GetLegalHolds() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("legalHolds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialDataSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialDataSources()))
        for i, v := range m.GetNoncustodialDataSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReviewSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewSets()))
        for i, v := range m.GetReviewSets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reviewSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSearches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearches()))
        for i, v := range m.GetSearches() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("searches", cast)
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
    if m.GetTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClosedBy sets the closedBy property value. The closedBy property
func (m *EdiscoveryCase) SetClosedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    if m != nil {
        m.closedBy = value
    }
}
// SetClosedDateTime sets the closedDateTime property value. The closedDateTime property
func (m *EdiscoveryCase) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.closedDateTime = value
    }
}
// SetCustodians sets the custodians property value. The custodians property
func (m *EdiscoveryCase) SetCustodians(value []EdiscoveryCustodianable)() {
    if m != nil {
        m.custodians = value
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *EdiscoveryCase) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetLegalHolds sets the legalHolds property value. The legalHolds property
func (m *EdiscoveryCase) SetLegalHolds(value []EdiscoveryHoldPolicyable)() {
    if m != nil {
        m.legalHolds = value
    }
}
// SetNoncustodialDataSources sets the noncustodialDataSources property value. The noncustodialDataSources property
func (m *EdiscoveryCase) SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)() {
    if m != nil {
        m.noncustodialDataSources = value
    }
}
// SetOperations sets the operations property value. The operations property
func (m *EdiscoveryCase) SetOperations(value []CaseOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetReviewSets sets the reviewSets property value. The reviewSets property
func (m *EdiscoveryCase) SetReviewSets(value []EdiscoveryReviewSetable)() {
    if m != nil {
        m.reviewSets = value
    }
}
// SetSearches sets the searches property value. The searches property
func (m *EdiscoveryCase) SetSearches(value []EdiscoverySearchable)() {
    if m != nil {
        m.searches = value
    }
}
// SetSettings sets the settings property value. The settings property
func (m *EdiscoveryCase) SetSettings(value EdiscoveryCaseSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetTags sets the tags property value. The tags property
func (m *EdiscoveryCase) SetTags(value []EdiscoveryReviewTagable)() {
    if m != nil {
        m.tags = value
    }
}
