package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// SourceCollectionable 
type SourceCollectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalSources()([]DataSourceable)
    GetAddToReviewSetOperation()(AddToReviewSetOperationable)
    GetContentQuery()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodianSources()([]DataSourceable)
    GetDataSourceScopes()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoncustodialSources()([]NoncustodialDataSourceable)
    SetAdditionalSources(value []DataSourceable)()
    SetAddToReviewSetOperation(value AddToReviewSetOperationable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodianSources(value []DataSourceable)()
    SetDataSourceScopes(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoncustodialSources(value []NoncustodialDataSourceable)()
}
