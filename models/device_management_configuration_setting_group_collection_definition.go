package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingGroupCollectionDefinition 
type DeviceManagementConfigurationSettingGroupCollectionDefinition struct {
    DeviceManagementConfigurationSettingGroupDefinition
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Maximum number of setting group count in the collection. Valid values 1 to 100
    maximumCount *int32
    // Minimum number of setting group count in the collection. Valid values 1 to 100
    minimumCount *int32
}
// NewDeviceManagementConfigurationSettingGroupCollectionDefinition instantiates a new DeviceManagementConfigurationSettingGroupCollectionDefinition and sets the default values.
func NewDeviceManagementConfigurationSettingGroupCollectionDefinition()(*DeviceManagementConfigurationSettingGroupCollectionDefinition) {
    m := &DeviceManagementConfigurationSettingGroupCollectionDefinition{
        DeviceManagementConfigurationSettingGroupDefinition: *NewDeviceManagementConfigurationSettingGroupDefinition(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingGroupCollectionDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingGroupCollectionDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingGroupCollectionDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingGroupDefinition.GetFieldDeserializers()
    res["maximumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCount(val)
        }
        return nil
    }
    res["minimumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCount(val)
        }
        return nil
    }
    return res
}
// GetMaximumCount gets the maximumCount property value. Maximum number of setting group count in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) GetMaximumCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumCount
    }
}
// GetMinimumCount gets the minimumCount property value. Minimum number of setting group count in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) GetMinimumCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumCount
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingGroupDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumCount", m.GetMaximumCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCount", m.GetMinimumCount())
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
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumCount sets the maximumCount property value. Maximum number of setting group count in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) SetMaximumCount(value *int32)() {
    if m != nil {
        m.maximumCount = value
    }
}
// SetMinimumCount sets the minimumCount property value. Minimum number of setting group count in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSettingGroupCollectionDefinition) SetMinimumCount(value *int32)() {
    if m != nil {
        m.minimumCount = value
    }
}
