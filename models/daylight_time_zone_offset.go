package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DaylightTimeZoneOffset 
type DaylightTimeZoneOffset struct {
    StandardTimeZoneOffset
}
// NewDaylightTimeZoneOffset instantiates a new daylightTimeZoneOffset and sets the default values.
func NewDaylightTimeZoneOffset()(*DaylightTimeZoneOffset) {
    m := &DaylightTimeZoneOffset{
        StandardTimeZoneOffset: *NewStandardTimeZoneOffset(),
    }
    odataTypeValue := "#microsoft.graph.daylightTimeZoneOffset"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDaylightTimeZoneOffsetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDaylightTimeZoneOffsetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDaylightTimeZoneOffset(), nil
}
// GetDaylightBias gets the daylightBias property value. The time offset from Coordinated Universal Time (UTC) for daylight saving time. This value is in minutes.
func (m *DaylightTimeZoneOffset) GetDaylightBias()(*int32) {
    val, err := m.GetBackingStore().Get("daylightBias")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DaylightTimeZoneOffset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StandardTimeZoneOffset.GetFieldDeserializers()
    res["daylightBias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaylightBias(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DaylightTimeZoneOffset) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DaylightTimeZoneOffset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StandardTimeZoneOffset.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("daylightBias", m.GetDaylightBias())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDaylightBias sets the daylightBias property value. The time offset from Coordinated Universal Time (UTC) for daylight saving time. This value is in minutes.
func (m *DaylightTimeZoneOffset) SetDaylightBias(value *int32)() {
    err := m.GetBackingStore().Set("daylightBias", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DaylightTimeZoneOffset) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DaylightTimeZoneOffsetable 
type DaylightTimeZoneOffsetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StandardTimeZoneOffsetable
    GetDaylightBias()(*int32)
    GetOdataType()(*string)
    SetDaylightBias(value *int32)()
    SetOdataType(value *string)()
}
