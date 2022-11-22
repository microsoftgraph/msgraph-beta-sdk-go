package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingValueTemplateReference setting value template reference information
type DeviceManagementConfigurationSettingValueTemplateReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Setting value template id
    settingValueTemplateId *string
    // Indicates whether to update policy setting value to match template setting default value
    useTemplateDefault *bool
}
// NewDeviceManagementConfigurationSettingValueTemplateReference instantiates a new deviceManagementConfigurationSettingValueTemplateReference and sets the default values.
func NewDeviceManagementConfigurationSettingValueTemplateReference()(*DeviceManagementConfigurationSettingValueTemplateReference) {
    m := &DeviceManagementConfigurationSettingValueTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingValueTemplateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingValueTemplateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingValueTemplateReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingValueTemplateReference) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingValueTemplateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["settingValueTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingValueTemplateId)
    res["useTemplateDefault"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUseTemplateDefault)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingValueTemplateReference) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingValueTemplateId gets the settingValueTemplateId property value. Setting value template id
func (m *DeviceManagementConfigurationSettingValueTemplateReference) GetSettingValueTemplateId()(*string) {
    return m.settingValueTemplateId
}
// GetUseTemplateDefault gets the useTemplateDefault property value. Indicates whether to update policy setting value to match template setting default value
func (m *DeviceManagementConfigurationSettingValueTemplateReference) GetUseTemplateDefault()(*bool) {
    return m.useTemplateDefault
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingValueTemplateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingValueTemplateId", m.GetSettingValueTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useTemplateDefault", m.GetUseTemplateDefault())
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
func (m *DeviceManagementConfigurationSettingValueTemplateReference) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingValueTemplateReference) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingValueTemplateId sets the settingValueTemplateId property value. Setting value template id
func (m *DeviceManagementConfigurationSettingValueTemplateReference) SetSettingValueTemplateId(value *string)() {
    m.settingValueTemplateId = value
}
// SetUseTemplateDefault sets the useTemplateDefault property value. Indicates whether to update policy setting value to match template setting default value
func (m *DeviceManagementConfigurationSettingValueTemplateReference) SetUseTemplateDefault(value *bool)() {
    m.useTemplateDefault = value
}
