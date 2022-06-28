package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EdiscoveryCaseable 
type EdiscoveryCaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Case_escapedable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]EdiscoveryCustodianable)
    GetExternalId()(*string)
    GetLegalHolds()([]EdiscoveryHoldPolicyable)
    GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]EdiscoveryReviewSetable)
    GetSearches()([]EdiscoverySearchable)
    GetSettings()(EdiscoveryCaseSettingsable)
    GetTags()([]EdiscoveryReviewTagable)
    SetClosedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []EdiscoveryCustodianable)()
    SetExternalId(value *string)()
    SetLegalHolds(value []EdiscoveryHoldPolicyable)()
    SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []EdiscoveryReviewSetable)()
    SetSearches(value []EdiscoverySearchable)()
    SetSettings(value EdiscoveryCaseSettingsable)()
    SetTags(value []EdiscoveryReviewTagable)()
}
