package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppRelationshipStateable 
type MobileAppRelationshipStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceId()(*string)
    GetErrorCode()(*int32)
    GetInstallState()(*ResultantAppState)
    GetInstallStateDetail()(*ResultantAppStateDetail)
    GetSourceIds()([]string)
    GetTargetDisplayName()(*string)
    GetTargetId()(*string)
    GetTargetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDeviceId(value *string)()
    SetErrorCode(value *int32)()
    SetInstallState(value *ResultantAppState)()
    SetInstallStateDetail(value *ResultantAppStateDetail)()
    SetSourceIds(value []string)()
    SetTargetDisplayName(value *string)()
    SetTargetId(value *string)()
    SetTargetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
