package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// AuditEventable 
type AuditEventable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityId()(*string)
    GetCategory()(*string)
    GetHttpVerb()(*string)
    GetInitiatedByAppId()(*string)
    GetInitiatedByUpn()(*string)
    GetInitiatedByUserId()(*string)
    GetIpAddress()(*string)
    GetRequestBody()(*string)
    GetRequestUrl()(*string)
    GetTenantIds()(*string)
    GetTenantNames()(*string)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityId(value *string)()
    SetCategory(value *string)()
    SetHttpVerb(value *string)()
    SetInitiatedByAppId(value *string)()
    SetInitiatedByUpn(value *string)()
    SetInitiatedByUserId(value *string)()
    SetIpAddress(value *string)()
    SetRequestBody(value *string)()
    SetRequestUrl(value *string)()
    SetTenantIds(value *string)()
    SetTenantNames(value *string)()
}
