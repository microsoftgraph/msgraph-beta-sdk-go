package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationGroupIdSourceOptions group id options type
type DeliveryOptimizationGroupIdSourceOptions struct {
    DeliveryOptimizationGroupIdSource
}
// NewDeliveryOptimizationGroupIdSourceOptions instantiates a new deliveryOptimizationGroupIdSourceOptions and sets the default values.
func NewDeliveryOptimizationGroupIdSourceOptions()(*DeliveryOptimizationGroupIdSourceOptions) {
    m := &DeliveryOptimizationGroupIdSourceOptions{
        DeliveryOptimizationGroupIdSource: *NewDeliveryOptimizationGroupIdSource(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationGroupIdSourceOptions"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeliveryOptimizationGroupIdSourceOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationGroupIdSourceOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationGroupIdSourceOptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationGroupIdSourceOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationGroupIdSource.GetFieldDeserializers()
    res["groupIdSourceOption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeliveryOptimizationGroupIdOptionsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupIdSourceOption(val.(*DeliveryOptimizationGroupIdOptionsType))
        }
        return nil
    }
    return res
}
// GetGroupIdSourceOption gets the groupIdSourceOption property value. Possible values for the DeliveryOptimizationGroupIdOptionsType setting.
func (m *DeliveryOptimizationGroupIdSourceOptions) GetGroupIdSourceOption()(*DeliveryOptimizationGroupIdOptionsType) {
    val, err := m.GetBackingStore().Get("groupIdSourceOption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeliveryOptimizationGroupIdOptionsType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationGroupIdSourceOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationGroupIdSource.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroupIdSourceOption() != nil {
        cast := (*m.GetGroupIdSourceOption()).String()
        err = writer.WriteStringValue("groupIdSourceOption", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupIdSourceOption sets the groupIdSourceOption property value. Possible values for the DeliveryOptimizationGroupIdOptionsType setting.
func (m *DeliveryOptimizationGroupIdSourceOptions) SetGroupIdSourceOption(value *DeliveryOptimizationGroupIdOptionsType)() {
    err := m.GetBackingStore().Set("groupIdSourceOption", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationGroupIdSourceOptionsable 
type DeliveryOptimizationGroupIdSourceOptionsable interface {
    DeliveryOptimizationGroupIdSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupIdSourceOption()(*DeliveryOptimizationGroupIdOptionsType)
    SetGroupIdSourceOption(value *DeliveryOptimizationGroupIdOptionsType)()
}
