package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthBusinessHoursLimit bandwidth business hours and percentages type
type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
    bandwidthBeginBusinessHours *int32
    // Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
    bandwidthEndBusinessHours *int32
    // Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
    bandwidthPercentageDuringBusinessHours *int32
    // Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
    bandwidthPercentageOutsideBusinessHours *int32
}
// NewDeliveryOptimizationBandwidthBusinessHoursLimit instantiates a new deliveryOptimizationBandwidthBusinessHoursLimit and sets the default values.
func NewDeliveryOptimizationBandwidthBusinessHoursLimit()(*DeliveryOptimizationBandwidthBusinessHoursLimit) {
    m := &DeliveryOptimizationBandwidthBusinessHoursLimit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationBandwidthBusinessHoursLimitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationBandwidthBusinessHoursLimit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBandwidthBeginBusinessHours gets the bandwidthBeginBusinessHours property value. Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthBeginBusinessHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthBeginBusinessHours
    }
}
// GetBandwidthEndBusinessHours gets the bandwidthEndBusinessHours property value. Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthEndBusinessHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthEndBusinessHours
    }
}
// GetBandwidthPercentageDuringBusinessHours gets the bandwidthPercentageDuringBusinessHours property value. Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthPercentageDuringBusinessHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthPercentageDuringBusinessHours
    }
}
// GetBandwidthPercentageOutsideBusinessHours gets the bandwidthPercentageOutsideBusinessHours property value. Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) GetBandwidthPercentageOutsideBusinessHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthPercentageOutsideBusinessHours
    }
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
    return res
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBandwidthBeginBusinessHours sets the bandwidthBeginBusinessHours property value. Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthBeginBusinessHours(value *int32)() {
    if m != nil {
        m.bandwidthBeginBusinessHours = value
    }
}
// SetBandwidthEndBusinessHours sets the bandwidthEndBusinessHours property value. Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthEndBusinessHours(value *int32)() {
    if m != nil {
        m.bandwidthEndBusinessHours = value
    }
}
// SetBandwidthPercentageDuringBusinessHours sets the bandwidthPercentageDuringBusinessHours property value. Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthPercentageDuringBusinessHours(value *int32)() {
    if m != nil {
        m.bandwidthPercentageDuringBusinessHours = value
    }
}
// SetBandwidthPercentageOutsideBusinessHours sets the bandwidthPercentageOutsideBusinessHours property value. Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
func (m *DeliveryOptimizationBandwidthBusinessHoursLimit) SetBandwidthPercentageOutsideBusinessHours(value *int32)() {
    if m != nil {
        m.bandwidthPercentageOutsideBusinessHours = value
    }
}
