package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeliveryOptimizationBandwidthBusinessHoursLimit bandwidth business hours and percentages type
type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeliveryOptimizationBandwidthBusinessHoursLimit instantiates a new deliveryOptimizationBandwidthBusinessHoursLimit and sets the default values.
func NewDeliveryOptimizationBandwidthBusinessHoursLimit()(*DeliveryOptimizationBandwidthBusinessHoursLimit) {
    m := &DeliveryOptimizationBandwidthBusinessHoursLimit{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationBandwidthBusinessHoursLimit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetAdditionalData()(map[string]any) {
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
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBandwidthBeginBusinessHours gets the bandwidthBeginBusinessHours property value. Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthBeginBusinessHours()(*int32) {
    val, err := m.GetBackingStore().Get("bandwidthBeginBusinessHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBandwidthEndBusinessHours gets the bandwidthEndBusinessHours property value. Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthEndBusinessHours()(*int32) {
    val, err := m.GetBackingStore().Get("bandwidthEndBusinessHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBandwidthPercentageDuringBusinessHours gets the bandwidthPercentageDuringBusinessHours property value. Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthPercentageDuringBusinessHours()(*int32) {
    val, err := m.GetBackingStore().Get("bandwidthPercentageDuringBusinessHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBandwidthPercentageOutsideBusinessHours gets the bandwidthPercentageOutsideBusinessHours property value. Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthPercentageOutsideBusinessHours()(*int32) {
    val, err := m.GetBackingStore().Get("bandwidthPercentageOutsideBusinessHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bandwidthBeginBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthBeginBusinessHours(val)
        }
        return nil
    }
    res["bandwidthEndBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthEndBusinessHours(val)
        }
        return nil
    }
    res["bandwidthPercentageDuringBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthPercentageDuringBusinessHours(val)
        }
        return nil
    }
    res["bandwidthPercentageOutsideBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthPercentageOutsideBusinessHours(val)
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
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetOdataType()(*string) {
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
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("bandwidthBeginBusinessHours", m.GetBandwidthBeginBusinessHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("bandwidthEndBusinessHours", m.GetBandwidthEndBusinessHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("bandwidthPercentageDuringBusinessHours", m.GetBandwidthPercentageDuringBusinessHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("bandwidthPercentageOutsideBusinessHours", m.GetBandwidthPercentageOutsideBusinessHours())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBandwidthBeginBusinessHours sets the bandwidthBeginBusinessHours property value. Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthBeginBusinessHours(value *int32)() {
    err := m.GetBackingStore().Set("bandwidthBeginBusinessHours", value)
    if err != nil {
        panic(err)
    }
}
// SetBandwidthEndBusinessHours sets the bandwidthEndBusinessHours property value. Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthEndBusinessHours(value *int32)() {
    err := m.GetBackingStore().Set("bandwidthEndBusinessHours", value)
    if err != nil {
        panic(err)
    }
}
// SetBandwidthPercentageDuringBusinessHours sets the bandwidthPercentageDuringBusinessHours property value. Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthPercentageDuringBusinessHours(value *int32)() {
    err := m.GetBackingStore().Set("bandwidthPercentageDuringBusinessHours", value)
    if err != nil {
        panic(err)
    }
}
// SetBandwidthPercentageOutsideBusinessHours sets the bandwidthPercentageOutsideBusinessHours property value. Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthPercentageOutsideBusinessHours(value *int32)() {
    err := m.GetBackingStore().Set("bandwidthPercentageOutsideBusinessHours", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationBandwidthBusinessHoursLimitable 
type DeliveryOptimizationBandwidthBusinessHoursLimitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBandwidthBeginBusinessHours()(*int32)
    GetBandwidthEndBusinessHours()(*int32)
    GetBandwidthPercentageDuringBusinessHours()(*int32)
    GetBandwidthPercentageOutsideBusinessHours()(*int32)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBandwidthBeginBusinessHours(value *int32)()
    SetBandwidthEndBusinessHours(value *int32)()
    SetBandwidthPercentageDuringBusinessHours(value *int32)()
    SetBandwidthPercentageOutsideBusinessHours(value *int32)()
    SetOdataType(value *string)()
}
