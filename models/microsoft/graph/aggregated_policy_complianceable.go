package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AggregatedPolicyComplianceable 
type AggregatedPolicyComplianceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliancePolicyId()(*string)
    GetCompliancePolicyName()(*string)
    GetCompliancePolicyPlatform()(*string)
    GetCompliancePolicyType()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumberOfCompliantDevices()(*int64)
    GetNumberOfErrorDevices()(*int64)
    GetNumberOfNonCompliantDevices()(*int64)
    GetPolicyModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetCompliancePolicyId(value *string)()
    SetCompliancePolicyName(value *string)()
    SetCompliancePolicyPlatform(value *string)()
    SetCompliancePolicyType(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumberOfCompliantDevices(value *int64)()
    SetNumberOfErrorDevices(value *int64)()
    SetNumberOfNonCompliantDevices(value *int64)()
    SetPolicyModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
