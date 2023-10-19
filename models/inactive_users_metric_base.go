package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveUsersMetricBase 
type InactiveUsersMetricBase struct {
    Entity
}
// NewInactiveUsersMetricBase instantiates a new inactiveUsersMetricBase and sets the default values.
func NewInactiveUsersMetricBase()(*InactiveUsersMetricBase) {
    m := &InactiveUsersMetricBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInactiveUsersMetricBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveUsersMetricBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.dailyInactiveUsersMetric":
                        return NewDailyInactiveUsersMetric(), nil
                    case "#microsoft.graph.monthlyInactiveUsersMetric":
                        return NewMonthlyInactiveUsersMetric(), nil
                }
            }
        }
    }
    return NewInactiveUsersMetricBase(), nil
}
// GetFactDate gets the factDate property value. The factDate property
func (m *InactiveUsersMetricBase) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("factDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveUsersMetricBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["factDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactDate(val)
        }
        return nil
    }
    res["inactive30DayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactive30DayCount(val)
        }
        return nil
    }
    res["inactive60DayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactive60DayCount(val)
        }
        return nil
    }
    res["inactive90DayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactive90DayCount(val)
        }
        return nil
    }
    return res
}
// GetInactive30DayCount gets the inactive30DayCount property value. The inactive30DayCount property
func (m *InactiveUsersMetricBase) GetInactive30DayCount()(*int64) {
    val, err := m.GetBackingStore().Get("inactive30DayCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetInactive60DayCount gets the inactive60DayCount property value. The inactive60DayCount property
func (m *InactiveUsersMetricBase) GetInactive60DayCount()(*int64) {
    val, err := m.GetBackingStore().Get("inactive60DayCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetInactive90DayCount gets the inactive90DayCount property value. The inactive90DayCount property
func (m *InactiveUsersMetricBase) GetInactive90DayCount()(*int64) {
    val, err := m.GetBackingStore().Get("inactive90DayCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InactiveUsersMetricBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("factDate", m.GetFactDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("inactive30DayCount", m.GetInactive30DayCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("inactive60DayCount", m.GetInactive60DayCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("inactive90DayCount", m.GetInactive90DayCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFactDate sets the factDate property value. The factDate property
func (m *InactiveUsersMetricBase) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetInactive30DayCount sets the inactive30DayCount property value. The inactive30DayCount property
func (m *InactiveUsersMetricBase) SetInactive30DayCount(value *int64)() {
    err := m.GetBackingStore().Set("inactive30DayCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInactive60DayCount sets the inactive60DayCount property value. The inactive60DayCount property
func (m *InactiveUsersMetricBase) SetInactive60DayCount(value *int64)() {
    err := m.GetBackingStore().Set("inactive60DayCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInactive90DayCount sets the inactive90DayCount property value. The inactive90DayCount property
func (m *InactiveUsersMetricBase) SetInactive90DayCount(value *int64)() {
    err := m.GetBackingStore().Set("inactive90DayCount", value)
    if err != nil {
        panic(err)
    }
}
// InactiveUsersMetricBaseable 
type InactiveUsersMetricBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetInactive30DayCount()(*int64)
    GetInactive60DayCount()(*int64)
    GetInactive90DayCount()(*int64)
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetInactive30DayCount(value *int64)()
    SetInactive60DayCount(value *int64)()
    SetInactive90DayCount(value *int64)()
}
