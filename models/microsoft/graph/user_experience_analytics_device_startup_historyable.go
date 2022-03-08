package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceStartupHistoryable 
type UserExperienceAnalyticsDeviceStartupHistoryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCoreBootTimeInMs()(*int32)
    GetCoreLoginTimeInMs()(*int32)
    GetDeviceId()(*string)
    GetFeatureUpdateBootTimeInMs()(*int32)
    GetGroupPolicyBootTimeInMs()(*int32)
    GetGroupPolicyLoginTimeInMs()(*int32)
    GetIsFeatureUpdate()(*bool)
    GetIsFirstLogin()(*bool)
    GetOperatingSystemVersion()(*string)
    GetResponsiveDesktopTimeInMs()(*int32)
    GetRestartCategory()(*UserExperienceAnalyticsOperatingSystemRestartCategory)
    GetRestartFaultBucket()(*string)
    GetRestartStopCode()(*string)
    GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTotalBootTimeInMs()(*int32)
    GetTotalLoginTimeInMs()(*int32)
    SetCoreBootTimeInMs(value *int32)()
    SetCoreLoginTimeInMs(value *int32)()
    SetDeviceId(value *string)()
    SetFeatureUpdateBootTimeInMs(value *int32)()
    SetGroupPolicyBootTimeInMs(value *int32)()
    SetGroupPolicyLoginTimeInMs(value *int32)()
    SetIsFeatureUpdate(value *bool)()
    SetIsFirstLogin(value *bool)()
    SetOperatingSystemVersion(value *string)()
    SetResponsiveDesktopTimeInMs(value *int32)()
    SetRestartCategory(value *UserExperienceAnalyticsOperatingSystemRestartCategory)()
    SetRestartFaultBucket(value *string)()
    SetRestartStopCode(value *string)()
    SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTotalBootTimeInMs(value *int32)()
    SetTotalLoginTimeInMs(value *int32)()
}
