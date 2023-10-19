package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DailyInactiveUsersByApplicationMetric 
type DailyInactiveUsersByApplicationMetric struct {
    InactiveUsersByApplicationMetricBase
}
// NewDailyInactiveUsersByApplicationMetric instantiates a new dailyInactiveUsersByApplicationMetric and sets the default values.
func NewDailyInactiveUsersByApplicationMetric()(*DailyInactiveUsersByApplicationMetric) {
    m := &DailyInactiveUsersByApplicationMetric{
        InactiveUsersByApplicationMetricBase: *NewInactiveUsersByApplicationMetricBase(),
    }
    return m
}
// CreateDailyInactiveUsersByApplicationMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDailyInactiveUsersByApplicationMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDailyInactiveUsersByApplicationMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DailyInactiveUsersByApplicationMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InactiveUsersByApplicationMetricBase.GetFieldDeserializers()
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
func (m *DailyInactiveUsersByApplicationMetric) GetInactive1DayCount()(*int64) {
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
func (m *DailyInactiveUsersByApplicationMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InactiveUsersByApplicationMetricBase.Serialize(writer)
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
func (m *DailyInactiveUsersByApplicationMetric) SetInactive1DayCount(value *int64)() {
    err := m.GetBackingStore().Set("inactive1DayCount", value)
    if err != nil {
        panic(err)
    }
}
// DailyInactiveUsersByApplicationMetricable 
type DailyInactiveUsersByApplicationMetricable interface {
    InactiveUsersByApplicationMetricBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInactive1DayCount()(*int64)
    SetInactive1DayCount(value *int64)()
}
