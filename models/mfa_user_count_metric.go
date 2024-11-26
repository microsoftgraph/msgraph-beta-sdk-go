package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MfaUserCountMetric struct {
    Entity
}
// NewMfaUserCountMetric instantiates a new MfaUserCountMetric and sets the default values.
func NewMfaUserCountMetric()(*MfaUserCountMetric) {
    m := &MfaUserCountMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMfaUserCountMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMfaUserCountMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMfaUserCountMetric(), nil
}
// GetCount gets the count property value. The count property
// returns a *int64 when successful
func (m *MfaUserCountMetric) GetCount()(*int64) {
    val, err := m.GetBackingStore().Get("count")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFactDate gets the factDate property value. The factDate property
// returns a *DateOnly when successful
func (m *MfaUserCountMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MfaUserCountMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
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
    res["mfaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMfaType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaType(val.(*MfaType))
        }
        return nil
    }
    return res
}
// GetMfaType gets the mfaType property value. The mfaType property
// returns a *MfaType when successful
func (m *MfaUserCountMetric) GetMfaType()(*MfaType) {
    val, err := m.GetBackingStore().Get("mfaType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MfaType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MfaUserCountMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("factDate", m.GetFactDate())
        if err != nil {
            return err
        }
    }
    if m.GetMfaType() != nil {
        cast := (*m.GetMfaType()).String()
        err = writer.WriteStringValue("mfaType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCount sets the count property value. The count property
func (m *MfaUserCountMetric) SetCount(value *int64)() {
    err := m.GetBackingStore().Set("count", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The factDate property
func (m *MfaUserCountMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaType sets the mfaType property value. The mfaType property
func (m *MfaUserCountMetric) SetMfaType(value *MfaType)() {
    err := m.GetBackingStore().Set("mfaType", value)
    if err != nil {
        panic(err)
    }
}
type MfaUserCountMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetMfaType()(*MfaType)
    SetCount(value *int64)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetMfaType(value *MfaType)()
}
