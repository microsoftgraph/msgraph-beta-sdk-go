package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DeviceCompliancePolicySettingStateSummaryable 
type DeviceCompliancePolicySettingStateSummaryable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetIntuneAccountId()(*string)
    GetIntuneSettingId()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotApplicableDeviceCount()(*int32)
    GetPendingDeviceCount()(*int32)
    GetPolicyType()(*string)
    GetSettingName()(*string)
    GetSucceededDeviceCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetIntuneAccountId(value *string)()
    SetIntuneSettingId(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotApplicableDeviceCount(value *int32)()
    SetPendingDeviceCount(value *int32)()
    SetPolicyType(value *string)()
    SetSettingName(value *string)()
    SetSucceededDeviceCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
