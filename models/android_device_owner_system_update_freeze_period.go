package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidDeviceOwnerSystemUpdateFreezePeriod represents one item in the list of freeze periods for Android Device Owner system updates
type AndroidDeviceOwnerSystemUpdateFreezePeriod struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidDeviceOwnerSystemUpdateFreezePeriod instantiates a new androidDeviceOwnerSystemUpdateFreezePeriod and sets the default values.
func NewAndroidDeviceOwnerSystemUpdateFreezePeriod()(*AndroidDeviceOwnerSystemUpdateFreezePeriod) {
    m := &AndroidDeviceOwnerSystemUpdateFreezePeriod{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerSystemUpdateFreezePeriod(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEndDay gets the endDay property value. The day of the end date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetEndDay()(*int32) {
    val, err := m.GetBackingStore().Get("endDay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEndMonth gets the endMonth property value. The month of the end date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetEndMonth()(*int32) {
    val, err := m.GetBackingStore().Get("endMonth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDay(val)
        }
        return nil
    }
    res["endMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonth(val)
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
    res["startDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDay(val)
        }
        return nil
    }
    res["startMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonth(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartDay gets the startDay property value. The day of the start date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetStartDay()(*int32) {
    val, err := m.GetBackingStore().Get("startDay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartMonth gets the startMonth property value. The month of the start date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetStartMonth()(*int32) {
    val, err := m.GetBackingStore().Get("startMonth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("endDay", m.GetEndDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("endMonth", m.GetEndMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("startDay", m.GetStartDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("startMonth", m.GetStartMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEndDay sets the endDay property value. The day of the end date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetEndDay(value *int32)() {
    err := m.GetBackingStore().Set("endDay", value)
    if err != nil {
        panic(err)
    }
}
// SetEndMonth sets the endMonth property value. The month of the end date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetEndMonth(value *int32)() {
    err := m.GetBackingStore().Set("endMonth", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDay sets the startDay property value. The day of the start date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetStartDay(value *int32)() {
    err := m.GetBackingStore().Set("startDay", value)
    if err != nil {
        panic(err)
    }
}
// SetStartMonth sets the startMonth property value. The month of the start date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetStartMonth(value *int32)() {
    err := m.GetBackingStore().Set("startMonth", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerSystemUpdateFreezePeriodable 
type AndroidDeviceOwnerSystemUpdateFreezePeriodable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEndDay()(*int32)
    GetEndMonth()(*int32)
    GetOdataType()(*string)
    GetStartDay()(*int32)
    GetStartMonth()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEndDay(value *int32)()
    SetEndMonth(value *int32)()
    SetOdataType(value *string)()
    SetStartDay(value *int32)()
    SetStartMonth(value *int32)()
}
