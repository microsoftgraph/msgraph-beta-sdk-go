package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthHoursWithPercentage bandwidth limit as a percentage with business hours.
type DeliveryOptimizationBandwidthHoursWithPercentage struct {
    DeliveryOptimizationBandwidth
}
// NewDeliveryOptimizationBandwidthHoursWithPercentage instantiates a new deliveryOptimizationBandwidthHoursWithPercentage and sets the default values.
func NewDeliveryOptimizationBandwidthHoursWithPercentage()(*DeliveryOptimizationBandwidthHoursWithPercentage) {
    m := &DeliveryOptimizationBandwidthHoursWithPercentage{
        DeliveryOptimizationBandwidth: *NewDeliveryOptimizationBandwidth(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationBandwidthHoursWithPercentage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeliveryOptimizationBandwidthHoursWithPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationBandwidthHoursWithPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationBandwidthHoursWithPercentage(), nil
}
// GetBandwidthBackgroundPercentageHours gets the bandwidthBackgroundPercentageHours property value. Background download percentage hours.
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) GetBandwidthBackgroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable) {
    val, err := m.GetBackingStore().Get("bandwidthBackgroundPercentageHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    }
    return nil
}
// GetBandwidthForegroundPercentageHours gets the bandwidthForegroundPercentageHours property value. Foreground download percentage hours.
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) GetBandwidthForegroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable) {
    val, err := m.GetBackingStore().Get("bandwidthForegroundPercentageHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationBandwidth.GetFieldDeserializers()
    res["bandwidthBackgroundPercentageHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthBackgroundPercentageHours(val.(DeliveryOptimizationBandwidthBusinessHoursLimitable))
        }
        return nil
    }
    res["bandwidthForegroundPercentageHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthForegroundPercentageHours(val.(DeliveryOptimizationBandwidthBusinessHoursLimitable))
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
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) GetOdataType()(*string) {
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
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationBandwidth.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bandwidthBackgroundPercentageHours", m.GetBandwidthBackgroundPercentageHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bandwidthForegroundPercentageHours", m.GetBandwidthForegroundPercentageHours())
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
// SetBandwidthBackgroundPercentageHours sets the bandwidthBackgroundPercentageHours property value. Background download percentage hours.
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) SetBandwidthBackgroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)() {
    err := m.GetBackingStore().Set("bandwidthBackgroundPercentageHours", value)
    if err != nil {
        panic(err)
    }
}
// SetBandwidthForegroundPercentageHours sets the bandwidthForegroundPercentageHours property value. Foreground download percentage hours.
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) SetBandwidthForegroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)() {
    err := m.GetBackingStore().Set("bandwidthForegroundPercentageHours", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeliveryOptimizationBandwidthHoursWithPercentage) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationBandwidthHoursWithPercentageable 
type DeliveryOptimizationBandwidthHoursWithPercentageable interface {
    DeliveryOptimizationBandwidthable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthBackgroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    GetBandwidthForegroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    GetOdataType()(*string)
    SetBandwidthBackgroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)()
    SetBandwidthForegroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)()
    SetOdataType(value *string)()
}
