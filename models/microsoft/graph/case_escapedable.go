package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// Case_escapedable 
type Case_escapedable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClosedBy()(IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]Custodianable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalHolds()([]LegalHoldable)
    GetNoncustodialDataSources()([]NoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]ReviewSetable)
    GetSettings()(CaseSettingsable)
    GetSourceCollections()([]SourceCollectionable)
    GetStatus()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus)
    GetTags()([]Tagable)
    SetClosedBy(value IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []Custodianable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalHolds(value []LegalHoldable)()
    SetNoncustodialDataSources(value []NoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []ReviewSetable)()
    SetSettings(value CaseSettingsable)()
    SetSourceCollections(value []SourceCollectionable)()
    SetStatus(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CaseStatus)()
    SetTags(value []Tagable)()
}
