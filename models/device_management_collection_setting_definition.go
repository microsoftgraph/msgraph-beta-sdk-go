package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCollectionSettingDefinition 
type DeviceManagementCollectionSettingDefinition struct {
    DeviceManagementSettingDefinition
    // The Setting Definition ID that describes what each element of the collection looks like
    elementDefinitionId *string
}
// NewDeviceManagementCollectionSettingDefinition instantiates a new DeviceManagementCollectionSettingDefinition and sets the default values.
func NewDeviceManagementCollectionSettingDefinition()(*DeviceManagementCollectionSettingDefinition) {
    m := &DeviceManagementCollectionSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementCollectionSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCollectionSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCollectionSettingDefinition(), nil
}
// GetElementDefinitionId gets the elementDefinitionId property value. The Setting Definition ID that describes what each element of the collection looks like
func (m *DeviceManagementCollectionSettingDefinition) GetElementDefinitionId()(*string) {
    return m.elementDefinitionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCollectionSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["elementDefinitionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetElementDefinitionId)
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementCollectionSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("elementDefinitionId", m.GetElementDefinitionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetElementDefinitionId sets the elementDefinitionId property value. The Setting Definition ID that describes what each element of the collection looks like
func (m *DeviceManagementCollectionSettingDefinition) SetElementDefinitionId(value *string)() {
    m.elementDefinitionId = value
}
