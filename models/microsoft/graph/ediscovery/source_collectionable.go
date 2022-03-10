package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// SourceCollectionable 
type SourceCollectionable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalSources()([]DataSourceable)
    GetAddToReviewSetOperation()(AddToReviewSetOperationable)
    GetContentQuery()(*string)
    GetCreatedBy()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodianSources()([]DataSourceable)
    GetDataSourceScopes()(*DataSourceScopes)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable)
    GetLastModifiedBy()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoncustodialSources()([]NoncustodialDataSourceable)
    SetAdditionalSources(value []DataSourceable)()
    SetAddToReviewSetOperation(value AddToReviewSetOperationable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodianSources(value []DataSourceable)()
    SetDataSourceScopes(value *DataSourceScopes)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)()
    SetLastModifiedBy(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoncustodialSources(value []NoncustodialDataSourceable)()
}
