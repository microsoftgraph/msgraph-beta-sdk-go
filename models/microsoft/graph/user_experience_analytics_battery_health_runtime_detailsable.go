package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthRuntimeDetailsable 
type UserExperienceAnalyticsBatteryHealthRuntimeDetailsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDevices()(*int32)
    GetBatteryRuntimeFair()(*int32)
    GetBatteryRuntimeGood()(*int32)
    GetBatteryRuntimePoor()(*int32)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActiveDevices(value *int32)()
    SetBatteryRuntimeFair(value *int32)()
    SetBatteryRuntimeGood(value *int32)()
    SetBatteryRuntimePoor(value *int32)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
