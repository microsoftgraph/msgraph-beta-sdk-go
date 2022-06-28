package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationMaxCacheSizeAbsolute 
type DeliveryOptimizationMaxCacheSizeAbsolute struct {
    DeliveryOptimizationMaxCacheSize
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
    maximumCacheSizeInGigabytes *int64
}
// NewDeliveryOptimizationMaxCacheSizeAbsolute instantiates a new DeliveryOptimizationMaxCacheSizeAbsolute and sets the default values.
func NewDeliveryOptimizationMaxCacheSizeAbsolute()(*DeliveryOptimizationMaxCacheSizeAbsolute) {
    m := &DeliveryOptimizationMaxCacheSizeAbsolute{
        DeliveryOptimizationMaxCacheSize: *NewDeliveryOptimizationMaxCacheSize(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeliveryOptimizationMaxCacheSizeAbsoluteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationMaxCacheSizeAbsoluteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationMaxCacheSizeAbsolute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
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
    if m == nil {
        return nil
    } else {
        return m.maximumCacheSizeInGigabytes
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumCacheSizeInGigabytes sets the maximumCacheSizeInGigabytes property value. Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
func (m *DeliveryOptimizationMaxCacheSizeAbsolute) SetMaximumCacheSizeInGigabytes(value *int64)() {
    if m != nil {
        m.maximumCacheSizeInGigabytes = value
    }
}
