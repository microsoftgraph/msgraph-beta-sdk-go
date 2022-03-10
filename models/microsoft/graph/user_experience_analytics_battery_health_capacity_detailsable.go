package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthCapacityDetailsable 
type UserExperienceAnalyticsBatteryHealthCapacityDetailsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDevices()(*int32)
    GetBatteryCapacityFair()(*int32)
    GetBatteryCapacityGood()(*int32)
    GetBatteryCapacityPoor()(*int32)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActiveDevices(value *int32)()
    SetBatteryCapacityFair(value *int32)()
    SetBatteryCapacityGood(value *int32)()
    SetBatteryCapacityPoor(value *int32)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
