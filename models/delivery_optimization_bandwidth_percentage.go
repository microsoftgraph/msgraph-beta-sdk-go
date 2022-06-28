package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthPercentage 
type DeliveryOptimizationBandwidthPercentage struct {
    DeliveryOptimizationBandwidth
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
    maximumBackgroundBandwidthPercentage *int32
    // Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
    maximumForegroundBandwidthPercentage *int32
}
// NewDeliveryOptimizationBandwidthPercentage instantiates a new DeliveryOptimizationBandwidthPercentage and sets the default values.
func NewDeliveryOptimizationBandwidthPercentage()(*DeliveryOptimizationBandwidthPercentage) {
    m := &DeliveryOptimizationBandwidthPercentage{
        DeliveryOptimizationBandwidth: *NewDeliveryOptimizationBandwidth(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeliveryOptimizationBandwidthPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationBandwidthPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationBandwidthPercentage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationBandwidthPercentage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationBandwidthPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationBandwidth.GetFieldDeserializers()
    res["maximumBackgroundBandwidthPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumBackgroundBandwidthPercentage(val)
        }
        return nil
    }
    res["maximumForegroundBandwidthPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumForegroundBandwidthPercentage(val)
        }
        return nil
    }
    return res
}
// GetMaximumBackgroundBandwidthPercentage gets the maximumBackgroundBandwidthPercentage property value. Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthPercentage) GetMaximumBackgroundBandwidthPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumBackgroundBandwidthPercentage
    }
}
// GetMaximumForegroundBandwidthPercentage gets the maximumForegroundBandwidthPercentage property value. Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthPercentage) GetMaximumForegroundBandwidthPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumForegroundBandwidthPercentage
    }
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationBandwidthPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationBandwidth.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumBackgroundBandwidthPercentage", m.GetMaximumBackgroundBandwidthPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumForegroundBandwidthPercentage", m.GetMaximumForegroundBandwidthPercentage())
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
func (m *DeliveryOptimizationBandwidthPercentage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumBackgroundBandwidthPercentage sets the maximumBackgroundBandwidthPercentage property value. Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthPercentage) SetMaximumBackgroundBandwidthPercentage(value *int32)() {
    if m != nil {
        m.maximumBackgroundBandwidthPercentage = value
    }
}
// SetMaximumForegroundBandwidthPercentage sets the maximumForegroundBandwidthPercentage property value. Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthPercentage) SetMaximumForegroundBandwidthPercentage(value *int32)() {
    if m != nil {
        m.maximumForegroundBandwidthPercentage = value
    }
}
