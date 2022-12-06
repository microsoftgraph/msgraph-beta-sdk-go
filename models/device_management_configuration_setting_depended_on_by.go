package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingDependedOnBy 
type DeviceManagementConfigurationSettingDependedOnBy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifier of child setting that is dependent on the current setting
    dependedOnBy *string
    // The OdataType property
    odataType *string
    // Value that determines if the child setting is required based on the parent setting's selection
    required *bool
}
// NewDeviceManagementConfigurationSettingDependedOnBy instantiates a new deviceManagementConfigurationSettingDependedOnBy and sets the default values.
func NewDeviceManagementConfigurationSettingDependedOnBy()(*DeviceManagementConfigurationSettingDependedOnBy) {
    m := &DeviceManagementConfigurationSettingDependedOnBy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingDependedOnByFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingDependedOnByFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingDependedOnBy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingDependedOnBy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDependedOnBy gets the dependedOnBy property value. Identifier of child setting that is dependent on the current setting
func (m *DeviceManagementConfigurationSettingDependedOnBy) GetDependedOnBy()(*string) {
    return m.dependedOnBy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingDependedOnBy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependedOnBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDependedOnBy)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["required"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequired)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingDependedOnBy) GetOdataType()(*string) {
    return m.odataType
}
// GetRequired gets the required property value. Value that determines if the child setting is required based on the parent setting's selection
func (m *DeviceManagementConfigurationSettingDependedOnBy) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingDependedOnBy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dependedOnBy", m.GetDependedOnBy())
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
        err := writer.WriteBoolValue("required", m.GetRequired())
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
func (m *DeviceManagementConfigurationSettingDependedOnBy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDependedOnBy sets the dependedOnBy property value. Identifier of child setting that is dependent on the current setting
func (m *DeviceManagementConfigurationSettingDependedOnBy) SetDependedOnBy(value *string)() {
    m.dependedOnBy = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingDependedOnBy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequired sets the required property value. Value that determines if the child setting is required based on the parent setting's selection
func (m *DeviceManagementConfigurationSettingDependedOnBy) SetRequired(value *bool)() {
    m.required = value
}
