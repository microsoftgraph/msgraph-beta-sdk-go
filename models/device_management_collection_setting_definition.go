package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCollectionSettingDefinition 
type DeviceManagementCollectionSettingDefinition struct {
    DeviceManagementSettingDefinition
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Setting Definition ID that describes what each element of the collection looks like
    elementDefinitionId *string
}
// NewDeviceManagementCollectionSettingDefinition instantiates a new DeviceManagementCollectionSettingDefinition and sets the default values.
func NewDeviceManagementCollectionSettingDefinition()(*DeviceManagementCollectionSettingDefinition) {
    m := &DeviceManagementCollectionSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementCollectionSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCollectionSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCollectionSettingDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementCollectionSettingDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetElementDefinitionId gets the elementDefinitionId property value. The Setting Definition ID that describes what each element of the collection looks like
func (m *DeviceManagementCollectionSettingDefinition) GetElementDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.elementDefinitionId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCollectionSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["elementDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElementDefinitionId(val)
        }
        return nil
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementCollectionSettingDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetElementDefinitionId sets the elementDefinitionId property value. The Setting Definition ID that describes what each element of the collection looks like
func (m *DeviceManagementCollectionSettingDefinition) SetElementDefinitionId(value *string)() {
    if m != nil {
        m.elementDefinitionId = value
    }
}
