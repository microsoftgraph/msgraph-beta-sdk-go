package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationGroupIdCustom groupId Support Types
type DeliveryOptimizationGroupIdCustom struct {
    DeliveryOptimizationGroupIdSource
}
// NewDeliveryOptimizationGroupIdCustom instantiates a new deliveryOptimizationGroupIdCustom and sets the default values.
func NewDeliveryOptimizationGroupIdCustom()(*DeliveryOptimizationGroupIdCustom) {
    m := &DeliveryOptimizationGroupIdCustom{
        DeliveryOptimizationGroupIdSource: *NewDeliveryOptimizationGroupIdSource(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationGroupIdCustom"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeliveryOptimizationGroupIdCustomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationGroupIdCustomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationGroupIdCustom(), nil
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
// GetGroupIdCustom gets the groupIdCustom property value. Specifies an arbitrary group ID that the device belongs to
func (m *DeliveryOptimizationGroupIdCustom) GetGroupIdCustom()(*string) {
    val, err := m.GetBackingStore().Get("groupIdCustom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeliveryOptimizationGroupIdCustom) GetOdataType()(*string) {
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupIdCustom sets the groupIdCustom property value. Specifies an arbitrary group ID that the device belongs to
func (m *DeliveryOptimizationGroupIdCustom) SetGroupIdCustom(value *string)() {
    err := m.GetBackingStore().Set("groupIdCustom", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeliveryOptimizationGroupIdCustom) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeliveryOptimizationGroupIdCustomable 
type DeliveryOptimizationGroupIdCustomable interface {
    DeliveryOptimizationGroupIdSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupIdCustom()(*string)
    GetOdataType()(*string)
    SetGroupIdCustom(value *string)()
    SetOdataType(value *string)()
}
