package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsDeviceBatteryDetail collection of properties describing the current status of the battery.
type UserExperienceAnalyticsDeviceBatteryDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceBatteryDetail instantiates a new userExperienceAnalyticsDeviceBatteryDetail and sets the default values.
func NewUserExperienceAnalyticsDeviceBatteryDetail()(*UserExperienceAnalyticsDeviceBatteryDetail) {
    m := &UserExperienceAnalyticsDeviceBatteryDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsDeviceBatteryDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceBatteryDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceBatteryDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBatteryId gets the batteryId property value. Uniquely identifies the batteries in a single device.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetBatteryId()(*string) {
    val, err := m.GetBackingStore().Get("batteryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batteryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryId(val)
        }
        return nil
    }
    res["fullBatteryDrainCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullBatteryDrainCount(val)
        }
        return nil
    }
    res["maxCapacityPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCapacityPercentage(val)
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
// GetFullBatteryDrainCount gets the fullBatteryDrainCount property value. Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by discharging it from 100% to 0%. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetFullBatteryDrainCount()(*int32) {
    val, err := m.GetBackingStore().Get("fullBatteryDrainCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaxCapacityPercentage gets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery. Unit in percentage and values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetMaxCapacityPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("maxCapacityPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceBatteryDetail) GetOdataType()(*string) {
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
func (m *UserExperienceAnalyticsDeviceBatteryDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("batteryId", m.GetBatteryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("fullBatteryDrainCount", m.GetFullBatteryDrainCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxCapacityPercentage", m.GetMaxCapacityPercentage())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBatteryId sets the batteryId property value. Uniquely identifies the batteries in a single device.
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetBatteryId(value *string)() {
    err := m.GetBackingStore().Set("batteryId", value)
    if err != nil {
        panic(err)
    }
}
// SetFullBatteryDrainCount sets the fullBatteryDrainCount property value. Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by discharging it from 100% to 0%. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetFullBatteryDrainCount(value *int32)() {
    err := m.GetBackingStore().Set("fullBatteryDrainCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxCapacityPercentage sets the maxCapacityPercentage property value. Ratio of current capacity and design capacity of the battery. Unit in percentage and values range from 0-100. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetMaxCapacityPercentage(value *int32)() {
    err := m.GetBackingStore().Set("maxCapacityPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsDeviceBatteryDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceBatteryDetailable 
type UserExperienceAnalyticsDeviceBatteryDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBatteryId()(*string)
    GetFullBatteryDrainCount()(*int32)
    GetMaxCapacityPercentage()(*int32)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBatteryId(value *string)()
    SetFullBatteryDrainCount(value *int32)()
    SetMaxCapacityPercentage(value *int32)()
    SetOdataType(value *string)()
}
