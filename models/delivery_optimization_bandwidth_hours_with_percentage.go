package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthHoursWithPercentage 
type DeliveryOptimizationBandwidthHoursWithPercentage struct {
    DeliveryOptimizationBandwidth
}
// NewDeliveryOptimizationBandwidthHoursWithPercentage instantiates a new DeliveryOptimizationBandwidthHoursWithPercentage and sets the default values.
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
    return res
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
// DeliveryOptimizationBandwidthHoursWithPercentageable 
type DeliveryOptimizationBandwidthHoursWithPercentageable interface {
    DeliveryOptimizationBandwidthable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthBackgroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    GetBandwidthForegroundPercentageHours()(DeliveryOptimizationBandwidthBusinessHoursLimitable)
    SetBandwidthBackgroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)()
    SetBandwidthForegroundPercentageHours(value DeliveryOptimizationBandwidthBusinessHoursLimitable)()
}
