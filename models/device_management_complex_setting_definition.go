package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplexSettingDefinition 
type DeviceManagementComplexSettingDefinition struct {
    DeviceManagementSettingDefinition
    // The definitions of each property of the complex setting
    propertyDefinitionIds []string
}
// NewDeviceManagementComplexSettingDefinition instantiates a new DeviceManagementComplexSettingDefinition and sets the default values.
func NewDeviceManagementComplexSettingDefinition()(*DeviceManagementComplexSettingDefinition) {
    m := &DeviceManagementComplexSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementComplexSettingDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplexSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplexSettingDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplexSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["propertyDefinitionIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPropertyDefinitionIds)
    return res
}
// GetPropertyDefinitionIds gets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) GetPropertyDefinitionIds()([]string) {
    return m.propertyDefinitionIds
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
    return nil
}
// SetPropertyDefinitionIds sets the propertyDefinitionIds property value. The definitions of each property of the complex setting
func (m *DeviceManagementComplexSettingDefinition) SetPropertyDefinitionIds(value []string)() {
    m.propertyDefinitionIds = value
}
