package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskRecurrenceable 
type PlannerTaskRecurrenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNextInSeriesTaskId()(*string)
    GetOccurrenceId()(*int32)
    GetOdataType()(*string)
    GetPreviousInSeriesTaskId()(*string)
    GetRecurrenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSchedule()(PlannerRecurrenceScheduleable)
    GetSeriesId()(*string)
    SetNextInSeriesTaskId(value *string)()
    SetOccurrenceId(value *int32)()
    SetOdataType(value *string)()
    SetPreviousInSeriesTaskId(value *string)()
    SetRecurrenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSchedule(value PlannerRecurrenceScheduleable)()
    SetSeriesId(value *string)()
}
