package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DailyInactiveUsersMetric 
type DailyInactiveUsersMetric struct {
    InactiveUsersMetricBase
}
// NewDailyInactiveUsersMetric instantiates a new dailyInactiveUsersMetric and sets the default values.
func NewDailyInactiveUsersMetric()(*DailyInactiveUsersMetric) {
    m := &DailyInactiveUsersMetric{
        InactiveUsersMetricBase: *NewInactiveUsersMetricBase(),
    }
    return m
}
// CreateDailyInactiveUsersMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDailyInactiveUsersMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDailyInactiveUsersMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DailyInactiveUsersMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InactiveUsersMetricBase.GetFieldDeserializers()
    res["inactive1DayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactive1DayCount(val)
        }
        return nil
    }
    return res
}
// GetInactive1DayCount gets the inactive1DayCount property value. The inactive1DayCount property
func (m *DailyInactiveUsersMetric) GetInactive1DayCount()(*int64) {
    val, err := m.GetBackingStore().Get("inactive1DayCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DailyInactiveUsersMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InactiveUsersMetricBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("inactive1DayCount", m.GetInactive1DayCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInactive1DayCount sets the inactive1DayCount property value. The inactive1DayCount property
func (m *DailyInactiveUsersMetric) SetInactive1DayCount(value *int64)() {
    err := m.GetBackingStore().Set("inactive1DayCount", value)
    if err != nil {
        panic(err)
    }
}
// DailyInactiveUsersMetricable 
type DailyInactiveUsersMetricable interface {
    InactiveUsersMetricBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInactive1DayCount()(*int64)
    SetInactive1DayCount(value *int64)()
}
