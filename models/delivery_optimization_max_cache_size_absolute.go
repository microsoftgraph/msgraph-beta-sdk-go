package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationMaxCacheSizeAbsolute delivery Optimization max cache size absolute type.
type DeliveryOptimizationMaxCacheSizeAbsolute struct {
    DeliveryOptimizationMaxCacheSize
    // The OdataType property
    OdataType *string
}
// NewDeliveryOptimizationMaxCacheSizeAbsolute instantiates a new deliveryOptimizationMaxCacheSizeAbsolute and sets the default values.
func NewDeliveryOptimizationMaxCacheSizeAbsolute()(*DeliveryOptimizationMaxCacheSizeAbsolute) {
    m := &DeliveryOptimizationMaxCacheSizeAbsolute{
        DeliveryOptimizationMaxCacheSize: *NewDeliveryOptimizationMaxCacheSize(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationMaxCacheSizeAbsolute"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeliveryOptimizationMaxCacheSizeAbsoluteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationMaxCacheSizeAbsoluteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationMaxCacheSizeAbsolute(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationMaxCacheSize.GetFieldDeserializers()
    res["maximumCacheSizeInGigabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheSizeInGigabytes(val)
        }
        return nil
    }
    return res
}
// GetMaximumCacheSizeInGigabytes gets the maximumCacheSizeInGigabytes property value. Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) GetMaximumCacheSizeInGigabytes()(*int64) {
    val, err := m.GetBackingStore().Get("maximumCacheSizeInGigabytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationMaxCacheSize.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("maximumCacheSizeInGigabytes", m.GetMaximumCacheSizeInGigabytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumCacheSizeInGigabytes sets the maximumCacheSizeInGigabytes property value. Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) SetMaximumCacheSizeInGigabytes(value *int64)() {
    err := m.GetBackingStore().Set("maximumCacheSizeInGigabytes", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationMaxCacheSizeAbsoluteable 
type DeliveryOptimizationMaxCacheSizeAbsoluteable interface {
    DeliveryOptimizationMaxCacheSizeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumCacheSizeInGigabytes()(*int64)
    SetMaximumCacheSizeInGigabytes(value *int64)()
}
