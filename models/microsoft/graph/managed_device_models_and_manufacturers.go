package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceModelsAndManufacturers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of Manufactures for managed devices in the account
    deviceManufacturers []string;
    // List of Models for managed devices in the account
    deviceModels []string;
}
// Instantiates a new managedDeviceModelsAndManufacturers and sets the default values.
func NewManagedDeviceModelsAndManufacturers()(*ManagedDeviceModelsAndManufacturers) {
    m := &ManagedDeviceModelsAndManufacturers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceModelsAndManufacturers) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceManufacturers property value. List of Manufactures for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceManufacturers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceManufacturers
    }
}
// Gets the deviceModels property value. List of Models for managed devices in the account
func (m *ManagedDeviceModelsAndManufacturers) GetDeviceModels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModels
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceModelsAndManufacturers) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceManufacturers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetDeviceManufacturers(res)
        return nil
    }
    res["deviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetDeviceModels(res)
        return nil
    }
    return res
}
func (m *ManagedDeviceModelsAndManufacturers) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceModelsAndManufacturers) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("deviceManufacturers", m.GetDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("deviceModels", m.GetDeviceModels())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagedDeviceModelsAndManufacturers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceManufacturers property value. List of Manufactures for managed devices in the account
// Parameters:
//  - value : Value to set for the deviceManufacturers property.
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceManufacturers(value []string)() {
    m.deviceManufacturers = value
}
// Sets the deviceModels property value. List of Models for managed devices in the account
// Parameters:
//  - value : Value to set for the deviceModels property.
func (m *ManagedDeviceModelsAndManufacturers) SetDeviceModels(value []string)() {
    m.deviceModels = value
}
