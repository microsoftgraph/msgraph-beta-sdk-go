package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentSettingsable 
type DeploymentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentApplicability()(ContentApplicabilitySettingsable)
    GetExpedite()(ExpediteSettingsable)
    GetMonitoring()(MonitoringSettingsable)
    GetOdataType()(*string)
    GetSchedule()(ScheduleSettingsable)
    GetUserExperience()(UserExperienceSettingsable)
    SetContentApplicability(value ContentApplicabilitySettingsable)()
    SetExpedite(value ExpediteSettingsable)()
    SetMonitoring(value MonitoringSettingsable)()
    SetOdataType(value *string)()
    SetSchedule(value ScheduleSettingsable)()
    SetUserExperience(value UserExperienceSettingsable)()
}
