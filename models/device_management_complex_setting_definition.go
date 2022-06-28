package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplexSettingDefinition 
type DeviceManagementComplexSettingDefinition struct {
    DeviceManagementSettingDefinition
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The definitions of each property of the complex setting
    propertyDefinitionIds []string
}
// NewDeviceManagementComplexSettingDefinition instantiates a new DeviceManagementComplexSettingDefinition and sets the default values.
func NewDeviceManagementComplexSettingDefinition()(*DeviceManagementComplexSettingDefinition) {
    m := &DeviceManagementComplexSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplexSettingDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComplexSettingDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplexSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["propertyDefinitionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPropertyDefinitionIds(res)
        }
        return nil
    }
    return res
}
// GetPropertyDefinitionIds gets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) GetPropertyDefinitionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.propertyDefinitionIds
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementComplexSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPropertyDefinitionIds() != nil {
        err = writer.WriteCollectionOfStringValues("propertyDefinitionIds", m.GetPropertyDefinitionIds())
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
func (m *DeviceManagementComplexSettingDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPropertyDefinitionIds sets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) SetPropertyDefinitionIds(value []string)() {
    if m != nil {
        m.propertyDefinitionIds = value
    }
}
