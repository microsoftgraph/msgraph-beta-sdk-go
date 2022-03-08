package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyServicePrincipalable 
type RiskyServicePrincipalable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccountEnabled()(*bool)
    GetAppId()(*string)
    GetDisplayName()(*string)
    GetHistory()([]RiskyServicePrincipalHistoryItemable)
    GetIsProcessing()(*bool)
    GetRiskDetail()(*RiskDetail)
    GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskLevel()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetServicePrincipalType()(*string)
    SetAccountEnabled(value *bool)()
    SetAppId(value *string)()
    SetDisplayName(value *string)()
    SetHistory(value []RiskyServicePrincipalHistoryItemable)()
    SetIsProcessing(value *bool)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskLevel(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetServicePrincipalType(value *string)()
}
