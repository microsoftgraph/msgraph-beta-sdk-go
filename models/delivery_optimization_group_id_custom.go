package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationGroupIdCustom 
type DeliveryOptimizationGroupIdCustom struct {
    DeliveryOptimizationGroupIdSource
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies an arbitrary group ID that the device belongs to
    groupIdCustom *string
}
// NewDeliveryOptimizationGroupIdCustom instantiates a new DeliveryOptimizationGroupIdCustom and sets the default values.
func NewDeliveryOptimizationGroupIdCustom()(*DeliveryOptimizationGroupIdCustom) {
    m := &DeliveryOptimizationGroupIdCustom{
        DeliveryOptimizationGroupIdSource: *NewDeliveryOptimizationGroupIdSource(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeliveryOptimizationGroupIdCustomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationGroupIdCustomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationGroupIdCustom(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationGroupIdCustom) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationGroupIdCustom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationGroupIdSource.GetFieldDeserializers()
    res["groupIdCustom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupIdCustom(val)
        }
        return nil
    }
    return res
}
// GetGroupIdCustom gets the groupIdCustom property value. Specifies an arbitrary group ID that the device belongs to
func (m *DeliveryOptimizationGroupIdCustom) GetGroupIdCustom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupIdCustom
    }
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationGroupIdCustom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationGroupIdSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupIdCustom", m.GetGroupIdCustom())
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
func (m *DeliveryOptimizationGroupIdCustom) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupIdCustom sets the groupIdCustom property value. Specifies an arbitrary group ID that the device belongs to
func (m *DeliveryOptimizationGroupIdCustom) SetGroupIdCustom(value *string)() {
    if m != nil {
        m.groupIdCustom = value
    }
}
