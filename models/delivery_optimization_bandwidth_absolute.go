package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthAbsolute bandwidth limits in kilobytes per second.
type DeliveryOptimizationBandwidthAbsolute struct {
    DeliveryOptimizationBandwidth
}
// NewDeliveryOptimizationBandwidthAbsolute instantiates a new deliveryOptimizationBandwidthAbsolute and sets the default values.
func NewDeliveryOptimizationBandwidthAbsolute()(*DeliveryOptimizationBandwidthAbsolute) {
    m := &DeliveryOptimizationBandwidthAbsolute{
        DeliveryOptimizationBandwidth: *NewDeliveryOptimizationBandwidth(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationBandwidthAbsolute"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeliveryOptimizationBandwidthAbsoluteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationBandwidthAbsoluteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationBandwidthAbsolute(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationBandwidthAbsolute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationBandwidth.GetFieldDeserializers()
    res["maximumDownloadBandwidthInKilobytesPerSecond"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumDownloadBandwidthInKilobytesPerSecond(val)
        }
        return nil
    }
    res["maximumUploadBandwidthInKilobytesPerSecond"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumUploadBandwidthInKilobytesPerSecond(val)
        }
        return nil
    }
    return res
}
// GetMaximumDownloadBandwidthInKilobytesPerSecond gets the maximumDownloadBandwidthInKilobytesPerSecond property value. Specifies the maximum download bandwidth in KiloBytes/second that the device can use across all concurrent download activities using Delivery Optimization. Valid values 0 to 4294967295
func (m *DeliveryOptimizationBandwidthAbsolute) GetMaximumDownloadBandwidthInKilobytesPerSecond()(*int64) {
    val, err := m.GetBackingStore().Get("maximumDownloadBandwidthInKilobytesPerSecond")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMaximumUploadBandwidthInKilobytesPerSecond gets the maximumUploadBandwidthInKilobytesPerSecond property value. Specifies the maximum upload bandwidth in KiloBytes/second that a device will use across all concurrent upload activity using Delivery Optimization (0-4000000). Valid values 0 to 4000000
func (m *DeliveryOptimizationBandwidthAbsolute) GetMaximumUploadBandwidthInKilobytesPerSecond()(*int64) {
    val, err := m.GetBackingStore().Get("maximumUploadBandwidthInKilobytesPerSecond")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationBandwidthAbsolute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationBandwidth.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("maximumDownloadBandwidthInKilobytesPerSecond", m.GetMaximumDownloadBandwidthInKilobytesPerSecond())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maximumUploadBandwidthInKilobytesPerSecond", m.GetMaximumUploadBandwidthInKilobytesPerSecond())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumDownloadBandwidthInKilobytesPerSecond sets the maximumDownloadBandwidthInKilobytesPerSecond property value. Specifies the maximum download bandwidth in KiloBytes/second that the device can use across all concurrent download activities using Delivery Optimization. Valid values 0 to 4294967295
func (m *DeliveryOptimizationBandwidthAbsolute) SetMaximumDownloadBandwidthInKilobytesPerSecond(value *int64)() {
    err := m.GetBackingStore().Set("maximumDownloadBandwidthInKilobytesPerSecond", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumUploadBandwidthInKilobytesPerSecond sets the maximumUploadBandwidthInKilobytesPerSecond property value. Specifies the maximum upload bandwidth in KiloBytes/second that a device will use across all concurrent upload activity using Delivery Optimization (0-4000000). Valid values 0 to 4000000
func (m *DeliveryOptimizationBandwidthAbsolute) SetMaximumUploadBandwidthInKilobytesPerSecond(value *int64)() {
    err := m.GetBackingStore().Set("maximumUploadBandwidthInKilobytesPerSecond", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationBandwidthAbsoluteable 
type DeliveryOptimizationBandwidthAbsoluteable interface {
    DeliveryOptimizationBandwidthable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumDownloadBandwidthInKilobytesPerSecond()(*int64)
    GetMaximumUploadBandwidthInKilobytesPerSecond()(*int64)
    SetMaximumDownloadBandwidthInKilobytesPerSecond(value *int64)()
    SetMaximumUploadBandwidthInKilobytesPerSecond(value *int64)()
}
