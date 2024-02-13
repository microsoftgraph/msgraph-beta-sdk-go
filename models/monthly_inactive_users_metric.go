package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MonthlyInactiveUsersMetric struct {
    InactiveUsersMetricBase
}
// NewMonthlyInactiveUsersMetric instantiates a new MonthlyInactiveUsersMetric and sets the default values.
func NewMonthlyInactiveUsersMetric()(*MonthlyInactiveUsersMetric) {
    m := &MonthlyInactiveUsersMetric{
        InactiveUsersMetricBase: *NewInactiveUsersMetricBase(),
    }
    return m
}
// CreateMonthlyInactiveUsersMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMonthlyInactiveUsersMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonthlyInactiveUsersMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MonthlyInactiveUsersMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InactiveUsersMetricBase.GetFieldDeserializers()
    res["inactiveCalendarMonthCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveCalendarMonthCount(val)
        }
        return nil
    }
    return res
}
// GetInactiveCalendarMonthCount gets the inactiveCalendarMonthCount property value. The inactiveCalendarMonthCount property
// returns a *int64 when successful
func (m *MonthlyInactiveUsersMetric) GetInactiveCalendarMonthCount()(*int64) {
    val, err := m.GetBackingStore().Get("inactiveCalendarMonthCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MonthlyInactiveUsersMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InactiveUsersMetricBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("inactiveCalendarMonthCount", m.GetInactiveCalendarMonthCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInactiveCalendarMonthCount sets the inactiveCalendarMonthCount property value. The inactiveCalendarMonthCount property
func (m *MonthlyInactiveUsersMetric) SetInactiveCalendarMonthCount(value *int64)() {
    err := m.GetBackingStore().Set("inactiveCalendarMonthCount", value)
    if err != nil {
        panic(err)
    }
}
type MonthlyInactiveUsersMetricable interface {
    InactiveUsersMetricBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInactiveCalendarMonthCount()(*int64)
    SetInactiveCalendarMonthCount(value *int64)()
}
