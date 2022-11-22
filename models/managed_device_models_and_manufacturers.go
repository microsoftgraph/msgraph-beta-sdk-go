package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceModelsAndManufacturers models and Manufactures meatadata for managed devices in the account
type ManagedDeviceModelsAndManufacturers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of Manufactures for managed devices in the account
    deviceManufacturers []string
    // List of Models for managed devices in the account
    deviceModels []string
    // The OdataType property
    odataType *string
}
// NewManagedDeviceModelsAndManufacturers instantiates a new managedDeviceModelsAndManufacturers and sets the default values.
func NewManagedDeviceModelsAndManufacturers()(*ManagedDeviceModelsAndManufacturers) {
    m := &ManagedDeviceModelsAndManufacturers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedDeviceModelsAndManufacturersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceModelsAndManufacturersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceModelsAndManufacturers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceModelsAndManufacturers) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceManufacturers gets the deviceManufacturers property value. List of Manufactures for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceManufacturers()([]string) {
    return m.deviceManufacturers
}
// GetDeviceModels gets the deviceModels property value. List of Models for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceModels()([]string) {
    return m.deviceModels
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceModelsAndManufacturers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceManufacturers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDeviceManufacturers)
    res["deviceModels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDeviceModels)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedDeviceModelsAndManufacturers) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ManagedDeviceModelsAndManufacturers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceManufacturers() != nil {
        err := writer.WriteCollectionOfStringValues("deviceManufacturers", m.GetDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceModels() != nil {
        err := writer.WriteCollectionOfStringValues("deviceModels", m.GetDeviceModels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *ManagedDeviceModelsAndManufacturers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceManufacturers sets the deviceManufacturers property value. List of Manufactures for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceManufacturers(value []string)() {
    m.deviceManufacturers = value
}
// SetDeviceModels sets the deviceModels property value. List of Models for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceModels(value []string)() {
    m.deviceModels = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDeviceModelsAndManufacturers) SetOdataType(value *string)() {
    m.odataType = value
}
