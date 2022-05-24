package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecurrencePatternable 
type RecurrencePatternable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDayOfMonth()(*int32)
    GetDaysOfWeek()([]string)
    GetFirstDayOfWeek()(*DayOfWeek)
    GetIndex()(*WeekIndex)
    GetInterval()(*int32)
    GetMonth()(*int32)
    GetType()(*RecurrencePatternType)
    SetDayOfMonth(value *int32)()
    SetDaysOfWeek(value []string)()
    SetFirstDayOfWeek(value *DayOfWeek)()
    SetIndex(value *WeekIndex)()
    SetInterval(value *int32)()
    SetMonth(value *int32)()
    SetType(value *RecurrencePatternType)()
}
