package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalRiskDetectionable 
type ServicePrincipalRiskDetectionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivity()(*ActivityType)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdditionalInfo()(*string)
    GetAppId()(*string)
    GetCorrelationId()(*string)
    GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDetectionTimingType()(*RiskDetectionTimingType)
    GetIpAddress()(*string)
    GetKeyIds()([]string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocation()(SignInLocationable)
    GetRequestId()(*string)
    GetRiskDetail()(*RiskDetail)
    GetRiskEventType()(*string)
    GetRiskLevel()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetServicePrincipalDisplayName()(*string)
    GetServicePrincipalId()(*string)
    GetSource()(*string)
    GetTokenIssuerType()(*TokenIssuerType)
    SetActivity(value *ActivityType)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdditionalInfo(value *string)()
    SetAppId(value *string)()
    SetCorrelationId(value *string)()
    SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDetectionTimingType(value *RiskDetectionTimingType)()
    SetIpAddress(value *string)()
    SetKeyIds(value []string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocation(value SignInLocationable)()
    SetRequestId(value *string)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskEventType(value *string)()
    SetRiskLevel(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetServicePrincipalDisplayName(value *string)()
    SetServicePrincipalId(value *string)()
    SetSource(value *string)()
    SetTokenIssuerType(value *TokenIssuerType)()
}
